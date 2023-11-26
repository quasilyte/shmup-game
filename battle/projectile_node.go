package battle

import (
	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/shmup-game/gamedata"
)

type projectileNode struct {
	sprite *ge.Sprite

	extraSpeed float64
	weapon     *gamedata.WeaponDesign
	world      *battleState

	rotation gmath.Rad

	pos   gmath.Vec
	toPos gmath.Vec

	charge float32

	disposed bool
}

type projectileConfig struct {
	world *battleState

	weapon *gamedata.WeaponDesign

	extraSpeed float64
	charge     float32
	pos        gmath.Vec
	targetPos  gmath.Vec
}

func newProjectileNode(config projectileConfig) *projectileNode {
	return &projectileNode{
		pos:        config.pos,
		toPos:      config.targetPos,
		world:      config.world,
		weapon:     config.weapon,
		extraSpeed: config.extraSpeed * 0.5,
		charge:     config.charge,
	}
}

func (p *projectileNode) Init(scene *ge.Scene) {
	p.sprite = scene.NewSprite(p.weapon.ProjectileImage)
	p.sprite.Pos.Base = &p.pos
	p.sprite.Rotation = &p.rotation
	if p.charge != 1 {
		var cs ge.ColorScale
		cs.A = float32(gmath.ClampMax(float64(p.charge)*1.5, 1))
		cs.G = 1 + p.charge*0.5
		cs.R = 1 - p.charge*0.5
		cs.B = 3 * (1 - p.charge)
		p.sprite.SetColorScale(cs)
	}
	p.world.stage.AddGraphics(p.sprite)

	p.rotation = p.pos.AngleToPoint(p.toPos)
}

func (p *projectileNode) IsDisposed() bool {
	return p.disposed
}

func (p *projectileNode) Dispose() {
	p.disposed = true
	p.sprite.Dispose()
}

func (p *projectileNode) Update(delta float64) {
	travelled := (p.weapon.ProjectileSpeed + p.extraSpeed) * delta

	if p.pos.DistanceTo(p.toPos) <= travelled {
		p.Dispose()
		return
	}

	p.pos = p.pos.MoveTowards(p.toPos, travelled)
}
