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

	dist     float64
	velocity gmath.Vec

	charge float32
	target *vesselNode

	dodge    bool
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

	p.dodge = p.target == p.world.human.vessel

	p.sprite = scene.NewSprite(p.weapon.ProjectileImage)
	p.sprite.Pos.Base = &p.pos
	p.sprite.Rotation = &p.rotation
	if !p.weapon.IgnoreChargeColor {
		p.sprite.SetColorScale(multiplyColorScale(calculateColorScale(p.charge), 0.7))
	}
	p.world.stage.AddGraphics(p.sprite)

	p.rotation = p.pos.AngleToPoint(p.toPos)

	p.dist = p.pos.DistanceTo(p.toPos)
	if p.weapon.ProjectileHoming != 0 {
		p.velocity = gmath.RadToVec(p.rotation).Mulf(p.weapon.ProjectileSpeed)
	}

	if p.weapon.ProjectileSpawnEffect != assets.ImageNone {
		effect := newEffectNode(effectConfig{
			world: p.world,
			pos:   p.pos,
			layer: aboveEffectLayer,
			image: p.weapon.ProjectileSpawnEffect,
			speed: fastEffectSpeed,
		})
		effect.noFlip = true
		effect.rotation = p.rotation
		effect.colorScale = p.sprite.GetColorScale()
		p.scene.AddObject(effect)
	}
}

func (p *projectileNode) IsDisposed() bool {
	return p.disposed
}

func (p *projectileNode) Dispose() {
	p.disposed = true
	p.sprite.Dispose()
}

func (p *projectileNode) Detonate(collided bool) {
	if p.weapon == gamedata.MineSpecialWeapon.Base {
		for i := 0; i < 3; i++ {
			p.scene.AddObject(newEffectNode(effectConfig{
				world:    p.world,
				pos:      p.pos.Add(p.scene.Rand().Offset(-10, 10)),
				layer:    slightlyAboveEffectLayer,
				speed:    fastEffectSpeed,
				image:    assets.ImageExplosionSmoke,
				rotation: p.scene.Rand().Rad(),
			}))
		}
		p.scene.AddObject(newEffectNode(effectConfig{
			world:    p.world,
			pos:      p.pos.Add(p.scene.Rand().Offset(-2, 2)),
			layer:    slightlyAboveEffectLayer,
			speed:    normalEffectSpeed,
			image:    assets.ImageFireExplosion,
			rotation: p.scene.Rand().Rad(),
		}))
	} else {
		effect := newEffectNode(effectConfig{
			world: p.world,
			pos:   p.pos,
			layer: aboveEffectLayer,
			image: p.weapon.ProjectileExplosion,
		})
		effect.colorScale = p.sprite.GetColorScale()
		p.scene.AddObject(effect)
	}

	if collided || (p.toPos.DistanceTo(p.target.pos) < p.weapon.ExplosionRange+p.target.design.Size) {
		p.target.OnDamage(p.weapon.Damage)
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

func (p *projectileNode) seek() gmath.Vec {
	dst := p.target.pos.Sub(p.pos).Normalized().Mulf(p.weapon.ProjectileSpeed)
	return dst.Sub(p.velocity).Normalized().Mulf(p.weapon.ProjectileHoming)
}

func (p *projectileNode) Update(delta float64) {
	travelled := p.movementSpeed() * delta
	p.dist -= travelled

	if p.weapon.ProjectileRotateSpeed != 0 {
		p.rotation += gmath.Rad(delta * p.weapon.ProjectileRotateSpeed)
	}

	if !p.slow && p.dist < 100 {
		p.slow = true
		if !p.weapon.IgnoreChargeColor {
			p.sprite.SetColorScale(calculateColorScale(p.charge))
		}
	}

	if p.weapon.CollisionRange != 0 {
		dist := p.pos.DistanceTo(p.target.pos)
		distDelta := dist - (p.weapon.CollisionRange + p.target.design.Size)
		if distDelta < 0 {
			p.Detonate(true)
			return
		}
		if p.dodge && distDelta < 20 {
			p.dodge = false
			p.world.result.DodgePoints++
		}
	}

	if p.dist <= 0 {
		p.Detonate(false)
		return
	}

	if p.weapon.ProjectileHoming == 0 {
		p.pos = p.pos.MoveTowards(p.toPos, travelled)
	} else {
		accel := p.seek()
		p.velocity = p.velocity.Add(accel.Mulf(delta)).ClampLen(p.weapon.ProjectileSpeed)
		p.rotation = p.velocity.Angle()
		p.pos = p.pos.Add(p.velocity.Mulf(delta))
	}
}
