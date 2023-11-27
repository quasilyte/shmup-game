package battle

import (
	"math"

	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/shmup-game/gamedata"
	"github.com/quasilyte/shmup-game/session"
	"github.com/quasilyte/shmup-game/viewport"
	"github.com/quasilyte/xm"
)

type Runner struct {
	session    *session.State
	state      *battleState
	eventQueue *queue[xm.StreamEvent]
	music      *gamedata.MusicInfo
	t          float64
}

type RunnerConfig struct {
	WorldRect gmath.Rect
	Music     *gamedata.MusicInfo
	Session   *session.State
	Stage     *viewport.Stage
}

func NewRunner(config RunnerConfig) *Runner {
	return &Runner{
		session: config.Session,
		state: &battleState{
			stage: config.Stage,
			rect:  config.WorldRect,
		},
		eventQueue: newQueue[xm.StreamEvent](320),
		music:      config.Music,
	}
}

func (r *Runner) Init(scene *ge.Scene) {
	r.state.scene = scene

	cam := viewport.NewCamera(r.state.stage, r.state.rect, 1920.0/2, 1080.0/2)
	cam.Offset = gmath.Vec{X: 900, Y: 900}
	scene.AddGraphics(cam)

	vessel := newVesselNode(vesselConfig{
		world:  r.state,
		design: gamedata.InterceptorDesign1,
		weapon: gamedata.IonCannonWeapon,
	})
	vessel.pos = gmath.Vec{X: 1024 / 2, Y: (1024 * 4) - 500}
	vessel.rotation = 3 * math.Pi / 2
	scene.AddObject(vessel)

	// overlay := scene.NewSprite(assets.ImageBattleOverlay)
	// overlay.Centered = false
	// cam.UI.AddGraphics(overlay)

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
			world:  r.state,
			design: gamedata.BossVessel1,
			weapon: gamedata.SpinCannonWeapon,
		})
		vessel.pos = gmath.Vec{X: 1024 / 2, Y: (1024 * 4) - 1000}
		vessel.rotation = math.Pi / 2
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
			note, vol := e.NoteEventData()
			if note == 97 || vol == 0 {
				return
			}
			r.eventQueue.Push(e)
		}
	})
}

func (r *Runner) Update(delta float64) {
	for r.eventQueue.Len() != 0 {
		current := r.eventQueue.Peek()
		if r.t < current.Time {
			break
		}
		r.eventQueue.Pop()
		if current.Kind == xm.EventSync {
			r.t = current.SyncEventData()
			continue
		}
		note, vol := current.NoteEventData()
		if current.Channel >= len(r.music.Channels) {
			continue
		}
		channelInfo := r.music.Channels[current.Channel]
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
		}
	}

	r.t += delta

	r.state.stage.Update()
}
