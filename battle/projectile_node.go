package battle

import (
	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/shmup-game/gamedata"
)

type projectileNode struct {
	sprite *ge.Sprite

	scene *ge.Scene

	extraSpeed float64
	weapon     *gamedata.WeaponDesign
	world      *battleState

	rotation gmath.Rad

	pos   gmath.Vec
	toPos gmath.Vec

	dist float64

	charge float32

	slow     bool
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
	p.scene = scene

	p.sprite = scene.NewSprite(p.weapon.ProjectileImage)
	p.sprite.Pos.Base = &p.pos
	p.sprite.Rotation = &p.rotation
	p.sprite.SetColorScale(calculateColorScale(p.charge))
	p.world.stage.AddGraphics(p.sprite)

	p.rotation = p.pos.AngleToPoint(p.toPos)

	p.dist = p.pos.DistanceTo(p.toPos)
}

func (p *projectileNode) IsDisposed() bool {
	return p.disposed
}

func (p *projectileNode) Dispose() {
	p.disposed = true
	p.sprite.Dispose()
}

func (p *projectileNode) Detonate() {
	effect := newEffectNode(effectConfig{
		world: p.world,
		pos:   p.toPos,
		layer: aboveEffectLayer,
		image: p.weapon.ProjectileExplosion,
	})
	effect.colorScale = p.sprite.GetColorScale()
	p.scene.AddObject(effect)

	p.Dispose()
}

func (p *projectileNode) movementSpeed() float64 {
	speed := p.weapon.ProjectileSpeed + p.extraSpeed
	if p.slow {
		speed *= 0.6
	}
	return speed
}

func (p *projectileNode) Update(delta float64) {
	travelled := p.movementSpeed() * delta
	p.dist -= travelled

	if !p.slow && p.dist < 128 {
		p.slow = true
		p.sprite.SetColorScale(multiplyColorScale(p.sprite.GetColorScale(), 2.0))
	}

	if p.dist <= 0 {
		p.Detonate()
		return
	}

	p.pos = p.pos.MoveTowards(p.toPos, travelled)
}
