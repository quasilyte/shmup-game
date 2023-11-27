package battle

import (
	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/shmup-game/assets"
	"github.com/quasilyte/shmup-game/gamedata"
)

type projectileNode struct {
	sprite *ge.Sprite

	scene *ge.Scene

	weapon *gamedata.WeaponDesign
	world  *battleState

	rotation gmath.Rad

	pos   gmath.Vec
	toPos gmath.Vec

	dist float64

	charge float32
	target *vesselNode

	slow     bool
	disposed bool
}

type projectileConfig struct {
	world *battleState

	target *vesselNode

	weapon *gamedata.WeaponDesign

	charge    float32
	pos       gmath.Vec
	targetPos gmath.Vec
}

func newProjectileNode(config projectileConfig) *projectileNode {
	return &projectileNode{
		pos:    config.pos,
		toPos:  config.targetPos,
		world:  config.world,
		weapon: config.weapon,
		charge: config.charge,
		target: config.target,
	}
}

func (p *projectileNode) Init(scene *ge.Scene) {
	p.scene = scene

	p.sprite = scene.NewSprite(p.weapon.ProjectileImage)
	p.sprite.Pos.Base = &p.pos
	p.sprite.Rotation = &p.rotation
	p.sprite.SetColorScale(multiplyColorScale(calculateColorScale(p.charge), 0.7))
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

	if p.toPos.DistanceTo(p.target.pos) < p.weapon.ExplosionRange+p.target.design.Size {
		if p.weapon.ImpactSound != assets.AudioNone {
			playSound(p.world, p.toPos, p.weapon.ImpactSound)
		}
	}

	p.Dispose()
}

func (p *projectileNode) movementSpeed() float64 {
	speed := p.weapon.ProjectileSpeed
	if p.slow {
		speed *= 0.55
	}
	return speed
}

func (p *projectileNode) Update(delta float64) {
	travelled := p.movementSpeed() * delta
	p.dist -= travelled

	if p.weapon.ProjectileRotateSpeed != 0 {
		p.rotation += gmath.Rad(delta * p.weapon.ProjectileRotateSpeed)
	}

	if !p.slow && p.dist < 100 {
		p.slow = true
		p.sprite.SetColorScale(calculateColorScale(p.charge))
	}

	if p.dist <= 0 {
		p.Detonate()
		return
	}

	p.pos = p.pos.MoveTowards(p.toPos, travelled)
}
