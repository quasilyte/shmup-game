package battle

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/ge"
	"github.com/quasilyte/ge/xslices"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gsignal"
	"github.com/quasilyte/shmup-game/assets"
	"github.com/quasilyte/shmup-game/gamedata"
	"github.com/quasilyte/shmup-game/session"
	"github.com/quasilyte/shmup-game/viewport"
	"github.com/quasilyte/xm"
)

type Runner struct {
	scene      *ge.Scene
	session    *session.State
	state      *battleState
	eventQueue *queue[xm.StreamEvent]
	music      *gamedata.MusicInfo
	t          float64
	startTime  time.Time

	transitionQueued bool

	cam            *viewport.Camera
	sectorSize     gmath.Vec
	sectorTextures []*ebiten.Image
	sectorSprites  []*ge.Sprite
	currentSectors []int
	tmpSlice       []int

	EventBattleOver gsignal.Event[Result]
}

type RunnerConfig struct {
	SectorSize     gmath.Vec
	SectorTextures []*ebiten.Image
	Music          *gamedata.MusicInfo
	Session        *session.State
	Stage          *viewport.Stage
}

func NewRunner(config RunnerConfig) *Runner {
	return &Runner{
		session: config.Session,
		state: &battleState{
			difficulty: config.Session.Settings.Difficulty,
			stage:      config.Stage,
			result:     &Result{},
		},
		eventQueue:     newQueue[xm.StreamEvent](320),
		music:          config.Music,
		sectorTextures: config.SectorTextures,
		sectorSize:     config.SectorSize,
		tmpSlice:       make([]int, 0, 12),
		sectorSprites:  make([]*ge.Sprite, 12),
	}
}

func (r *Runner) Init(scene *ge.Scene) {
	r.state.scene = scene
	r.startTime = time.Now()
	r.scene = scene

	cam := viewport.NewCamera(r.state.stage, r.state.rect, 1920.0/2, 1080.0/2)
	r.cam = cam
	scene.AddGraphics(cam)

	playerWeapons := [...]*gamedata.WeaponDesign{
		gamedata.PulseLaserWeapon,
		gamedata.RearCannonWeapon,
		gamedata.IonCannonWeapon,
		gamedata.TwinCannonWeapon,
	}
	playerSpecialWeapons := [...]*gamedata.SpecialWeaponDesign{
		gamedata.DashSpecialWeapon,
		gamedata.MegaBombSpecialWeapon,
		gamedata.SpinningShieldSpecialWeapon,
	}
	playerVessels := [...]*gamedata.VesselDesign{
		gamedata.InterceptorDesign1,
		gamedata.InterceptorDesign2,
	}

	vessel := newVesselNode(vesselConfig{
		world:         r.state,
		design:        playerVessels[r.session.Settings.Vessel],
		weapon:        playerWeapons[r.session.Settings.Weapon],
		specialWeapon: playerSpecialWeapons[r.session.Settings.Special],
	})
	vessel.pos = gmath.Vec{X: 0, Y: 0}
	vessel.rotation = 3 * math.Pi / 2
	vessel.EventDestroyed.Connect(nil, func(gsignal.Void) {
		r.onBattleOver(false)
	})
	scene.AddObject(vessel)

	overlay := scene.NewSprite(assets.ImageBattleOverlay)
	overlay.Centered = false
	cam.UI.AddGraphics(overlay)

	human := &humanPlayer{
		world:  r.state,
		input:  r.session.Input,
		camera: cam,
		vessel: vessel,
	}
	scene.AddObject(human)
	r.state.human = human

	{
		vessel := newVesselNode(vesselConfig{
			world:         r.state,
			design:        gamedata.BossVessel1,
			weapon:        gamedata.SpinCannonWeapon,
			specialWeapon: gamedata.HomingMissileSpecialWeapon,
		})
		vessel.pos = gmath.Vec{X: -80, Y: -240}
		vessel.rotation = math.Pi / 2
		vessel.bot = true
		vessel.EventDestroyed.Connect(nil, func(gsignal.Void) {
			r.onBattleOver(true)
		})
		scene.AddObject(vessel)

		bot := newBoss1Player(r.state, vessel)
		scene.AddObject(bot)
		r.state.bot = bot

		vessel.enemy = human.vessel
		human.vessel.enemy = vessel
	}

	r.session.EventPlayerUpdate.Connect(nil, func(e xm.StreamEvent) {
		switch e.Kind {
		case xm.EventSync:
			r.eventQueue.Push(e)

		case xm.EventNote:
			note, _, vol := e.NoteEventData()
			if note == 97 || vol == 0 {
				return
			}
			r.eventQueue.Push(e)
		}
	})

	for i := range r.sectorSprites {
		s := ge.NewSprite(scene.Context())
		s.Centered = false
		s.Visible = false
		r.sectorSprites[i] = s
		r.cam.Stage.AddSpriteBelow(s)
	}

	r.updateSectors()

	switch r.session.Settings.Difficulty {
	case 0: // Easy
		r.state.botDamageMultiplier = 1.2
		r.state.playerDamageMultiplier = 0.7
	case 1: // Normal
		r.state.botDamageMultiplier = 1.0
		r.state.playerDamageMultiplier = 1.0
	case 2: // Hard
		r.state.botDamageMultiplier = 0.85
		r.state.playerDamageMultiplier = 1.2
	case 3: // Nightmare
		r.state.botDamageMultiplier = 0.6
		r.state.playerDamageMultiplier = 1.25
	}
}

func (r *Runner) onBattleOver(victory bool) {
	if r.transitionQueued {
		return
	}
	r.transitionQueued = true

	r.state.result.Victory = victory
	r.state.result.TimePlayed = time.Since(r.startTime)
	r.state.result.Difficulty = r.session.Settings.Difficulty

	r.scene.DelayedCall(2, func() {
		r.EventBattleOver.Emit(*r.state.result)
	})
}

func (r *Runner) findChannelVariant(inst int, variants []gamedata.MusicChannelVariant) *gamedata.MusicChannelVariant {
	for i, v := range variants {
		if v.Inst == inst {
			return &variants[i]
		}
	}
	return nil
}

func (r *Runner) Update(delta float64) {
	for r.eventQueue.Len() != 0 {
		current := r.eventQueue.Peek()
		if current.Kind == xm.EventSync && r.t == 0 {
			r.t = current.SyncEventData()
			r.eventQueue.Pop()
			continue
		}
		if r.t < current.Time {
			break
		}
		r.eventQueue.Pop()
		if current.Kind == xm.EventSync {
			r.t = current.SyncEventData()
			continue
		}
		note, inst, vol := current.NoteEventData()
		if current.Channel >= len(r.music.Channels) {
			continue
		}
		channelInfo := r.findChannelVariant(inst, r.music.Channels[current.Channel])
		if channelInfo == nil {
			continue
		}
		switch channelInfo.Kind {
		case gamedata.ChannelPlayerAttack:
			if note < channelInfo.HighNote {
				r.state.human.vessel.orders.fire = true
				r.state.human.vessel.orders.fireCharge = vol
			} else {
				r.state.human.vessel.orders.altFire = true
				r.state.human.vessel.orders.altFireCharge = vol
			}
		case gamedata.ChannelEnemyAttack:
			v := r.state.bot.GetVessel()
			v.orders.fire = true
			v.orders.fireCharge = vol
		case gamedata.ChannelEnemyAltAttack:
			v := r.state.bot.GetVessel()
			v.orders.altFire = true
			v.orders.altFireCharge = vol
		case gamedata.ChannelEnemySpecialAttack:
			v := r.state.bot.GetVessel()
			v.orders.specialFire = true
			v.orders.specialFireCharge = vol
		}
	}

	r.updateSectors()

	r.t += delta

	r.state.stage.Update()
}

func (r *Runner) updateSectors() {
	startX, startY, endX, endY := r.getVisibleSectors(r.state.human.CameraPos(), 1080/2, 1080/2)

	r.tmpSlice = r.tmpSlice[:0]
	for y := startY; y <= endY; y++ {
		for x := startX; x <= endX; x++ {
			r.tmpSlice = append(r.tmpSlice, r.getSectorID(x, y))
		}
	}
	if xslices.Equal(r.tmpSlice, r.currentSectors) {
		// Nothing to do.
		return
	}

	// Update sector sprites.
	r.currentSectors = r.currentSectors[:0]
	r.currentSectors = append(r.currentSectors, r.tmpSlice...)
	for _, s := range r.sectorSprites {
		s.Visible = false
	}
	spriteIndex := 0
	for y := startY; y <= endY; y++ {
		for x := startX; x <= endX; x++ {
			s := r.sectorSprites[spriteIndex]
			s.Visible = true
			s.Pos.Offset = r.getSectorPos(x, y)
			s.SetImage(resource.Image{Data: r.getSectorTexture(x, y)})
			spriteIndex++
		}
	}
}

func (r *Runner) getSectorID(x, y int) int {
	return y + 10*x
}

func (r *Runner) getSectorCoords(pos gmath.Vec) (x, y int) {
	return int(pos.X / r.sectorSize.X), int(pos.Y / r.sectorSize.Y)
}

func (r *Runner) getSectorRect(x, y int) gmath.Rect {
	return gmath.Rect{
		Min: gmath.Vec{X: float64(x) * r.sectorSize.X, Y: float64(y) * r.sectorSize.Y},
		Max: gmath.Vec{X: float64(x+1) * r.sectorSize.X, Y: float64(y+1) * r.sectorSize.Y},
	}
}

func (r *Runner) getSectorPos(x, y int) gmath.Vec {
	return gmath.Vec{
		X: float64(x) * r.sectorSize.X,
		Y: float64(y) * r.sectorSize.Y,
	}
}

func (r *Runner) getSectorTexture(x, y int) *ebiten.Image {
	hash := uint64(r.getSectorID(x, y))
	return r.sectorTextures[hash%uint64(len(r.sectorTextures))]
}

func (r *Runner) getVisibleSectors(pos gmath.Vec, camWidth, camHeight float64) (startX, startY, endX, endY int) {
	sectorX, sectorY := r.getSectorCoords(pos)
	sectorRect := r.getSectorRect(sectorX, sectorY)

	// Determine how many sectors we need to consider.
	// In the simplest case, it's a single sector,
	// but sometimes we need to check the adjacent sectors too.
	startX = sectorX
	startY = sectorY
	endX = sectorX
	endY = sectorY
	leftmostPos := gmath.Vec{X: pos.X - camWidth, Y: pos.Y - camHeight}
	rightmostPos := gmath.Vec{X: pos.X + camWidth, Y: pos.Y + camHeight}
	if leftmostPos.X < sectorRect.Min.X {
		delta := sectorRect.Min.X - leftmostPos.X
		startX -= int(math.Ceil(delta / r.sectorSize.X))
	}
	if rightmostPos.X > sectorRect.Max.X {
		delta := rightmostPos.X - sectorRect.Max.X
		endX += int(math.Ceil(delta / r.sectorSize.X))
	}
	if leftmostPos.Y < sectorRect.Min.Y {
		delta := sectorRect.Min.Y - leftmostPos.Y
		startY -= int(math.Ceil(delta / r.sectorSize.Y))
	}
	if rightmostPos.Y > sectorRect.Max.Y {
		delta := rightmostPos.Y - sectorRect.Max.Y
		endY += int(math.Ceil(delta / r.sectorSize.Y))
	}

	return startX, startY, endX, endY
}
