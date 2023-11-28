package battle

import (
	"math"

	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gsignal"
	"github.com/quasilyte/shmup-game/assets"
	"github.com/quasilyte/shmup-game/gamedata"
)

type vesselNode struct {
	rotation gmath.Rad
	pos      gmath.Vec

	enemy *vesselNode

	world *battleState
	scene *ge.Scene

	design *gamedata.VesselDesign
	weapon *weaponSystem

	hp     float64
	energy float64

	strafing      bool
	thrusting     bool
	rotatingLeft  bool
	rotatingRight bool

	velocity gmath.Vec

	rotationInitialVelocity float64
	rotationAcceleration    float64
	rotationVelocity        float64

	sprite *ge.Sprite

	orders vesselOrders

	disposed bool

	EventOnDamage  gsignal.Event[gamedata.Damage]
	EventDestroyed gsignal.Event[gsignal.Void]
}

type vesselOrders struct {
	turbo       bool
	strafe      bool
	rotateLeft  bool
	rotateRight bool

	specialFire       bool
	fire              bool
	altFire           bool
	specialFireCharge float32
	fireCharge        float32
	altFireCharge     float32
}

type vesselConfig struct {
	world         *battleState
	design        *gamedata.VesselDesign
	weapon        *gamedata.WeaponDesign
	specialWeapon *gamedata.SpecialWeaponDesign
}

func newVesselNode(config vesselConfig) *vesselNode {
	v := &vesselNode{
		world:  config.world,
		design: config.design,
	}
	v.weapon = newWeaponSystem(config.world, v, config.weapon, config.specialWeapon)
	return v
}

func (v *vesselNode) Init(scene *ge.Scene) {
	v.scene = scene
	v.sprite = scene.NewSprite(v.design.Image)
	v.sprite.Pos.Base = &v.pos
	v.sprite.Rotation = &v.rotation
	v.world.stage.AddSpriteAbove(v.sprite)

	v.hp = v.design.HP
	v.energy = v.design.Energy

	v.rotationAcceleration = float64(v.design.RotationMaxSpeed * 1.2)
	v.rotationInitialVelocity = float64(v.design.RotationMaxSpeed * 0.1)
}

func (v *vesselNode) OnDamage(dmg gamedata.Damage) {
	return
	if v.disposed {
		return
	}

	v.hp = gmath.ClampMin(v.hp-dmg.HP, 0)
	if v.hp == 0 {
		v.destroy()
		return
	}

	v.EventOnDamage.Emit(dmg)
}

func (v *vesselNode) destroy() {
	for i := 0; i < 7; i++ {
		v.scene.AddObject(newEffectNode(effectConfig{
			world:    v.world,
			pos:      v.pos.Add(v.scene.Rand().Offset(-22, 22)),
			layer:    slightlyAboveEffectLayer,
			speed:    fastEffectSpeed,
			image:    assets.ImageExplosionSmoke,
			rotation: v.scene.Rand().Rad(),
		}))
	}
	for i := 0; i < 4; i++ {
		v.scene.AddObject(newEffectNode(effectConfig{
			world:    v.world,
			pos:      v.pos.Add(v.scene.Rand().Offset(-16, 16)),
			layer:    slightlyAboveEffectLayer,
			speed:    normalEffectSpeed,
			image:    assets.ImageFireExplosion,
			rotation: v.scene.Rand().Rad(),
		}))
	}

	v.Dispose()
	v.EventDestroyed.Emit(gsignal.Void{})
}

func (v *vesselNode) Dispose() {
	v.sprite.Dispose()
	v.disposed = true
}

func (v *vesselNode) IsDisposed() bool {
	return v.disposed
}

func (v *vesselNode) strafeSpeed() float64 {
	return v.design.StrafeSpeed
}

func (v *vesselNode) acceleration() float64 {
	return v.design.Acceleration
}

func (v *vesselNode) maxSpeed() float64 {
	return v.design.Speed
}

func (v *vesselNode) rotationSpeed() gmath.Rad {
	return gmath.Rad(v.rotationVelocity)
}

func (v *vesselNode) currentSpeed() float64 {
	return v.velocity.Len()
}

func (v *vesselNode) Update(delta float64) {
	orders := v.orders
	v.orders = vesselOrders{}

	wasStrafing := v.strafing
	v.strafing = false

	wasThrusting := v.thrusting
	v.thrusting = false

	wasRotatingLeft := v.rotatingLeft
	v.rotatingLeft = false

	wasRotatingRight := v.rotatingRight
	v.rotatingRight = false

	switch {
	case orders.rotateLeft == orders.rotateRight:
		// Do nothing.
		v.rotationVelocity = v.rotationInitialVelocity
		v.sprite.FrameOffset.X = 0
	case orders.rotateLeft:
		if orders.strafe {
			v.strafing = true
			if !wasStrafing {
				v.velocity = v.velocity.Add(gmath.RadToVec(v.rotation - math.Pi/2).Mulf(50 * v.strafeSpeed() * delta)).ClampLen(0.6 * v.strafeSpeed())
				v.scene.AddObject(newEffectNode(effectConfig{
					world:    v.world,
					pos:      v.pos.MoveInDirection(20, v.rotation+math.Pi-0.5),
					layer:    slightlyAboveEffectLayer,
					speed:    fastEffectSpeed,
					image:    assets.ImageDashEffect,
					rotation: v.rotation + math.Pi/4 + math.Pi,
				}))
			}
			v.velocity = v.velocity.Add(gmath.RadToVec(v.rotation - math.Pi/2).Mulf(0.8 * v.strafeSpeed() * delta)).ClampLen(v.strafeSpeed())
		} else {
			if !wasRotatingLeft {
				v.rotationVelocity = v.rotationInitialVelocity
			}
			v.rotationVelocity = gmath.ClampMax(v.rotationVelocity+v.rotationAcceleration*delta, float64(v.design.RotationMaxSpeed))
			v.rotation -= v.rotationSpeed() * gmath.Rad(delta)
			v.rotatingLeft = true
		}
		v.sprite.FrameOffset.X = 2 * v.sprite.FrameWidth
	case orders.rotateRight:
		if orders.strafe {
			v.strafing = true
			if !wasStrafing {
				v.velocity = v.velocity.Add(gmath.RadToVec(v.rotation + math.Pi/2).Mulf(50 * v.strafeSpeed() * delta)).ClampLen(0.6 * v.strafeSpeed())
				v.scene.AddObject(newEffectNode(effectConfig{
					world:    v.world,
					pos:      v.pos.MoveInDirection(20, v.rotation-math.Pi+0.5),
					layer:    slightlyAboveEffectLayer,
					speed:    fastEffectSpeed,
					image:    assets.ImageDashEffect,
					rotation: v.rotation - math.Pi/4,
				}))
			}
			v.velocity = v.velocity.Add(gmath.RadToVec(v.rotation + math.Pi/2).Mulf(0.8 * v.strafeSpeed() * delta)).ClampLen(v.strafeSpeed())
		} else {
			if !wasRotatingRight {
				v.rotationVelocity = v.rotationInitialVelocity
			}
			v.rotationVelocity = gmath.ClampMax(v.rotationVelocity+v.rotationAcceleration*delta, float64(v.design.RotationMaxSpeed))
			v.rotation += v.rotationSpeed() * gmath.Rad(delta)
			v.rotatingRight = true
		}
		v.sprite.FrameOffset.X = 1 * v.sprite.FrameWidth
	}

	accel := v.acceleration()

	if orders.turbo {
		v.thrusting = true
		if !wasThrusting {
			v.velocity = v.velocity.Add(gmath.RadToVec(v.rotation).Mulf(25 * accel * delta)).ClampLen(0.5 * v.maxSpeed())
			v.scene.AddObject(newEffectNode(effectConfig{
				world:    v.world,
				pos:      v.pos.MoveInDirection(20, v.rotation-math.Pi),
				layer:    slightlyAboveEffectLayer,
				speed:    veryFastEffectSpeed,
				image:    assets.ImageDashEffect,
				rotation: v.rotation - math.Pi/2,
			}))
		}
		v.velocity = v.velocity.Add(gmath.RadToVec(v.rotation).Mulf(accel * delta)).ClampLen(v.maxSpeed())
	}

	v.pos = v.pos.Add(v.velocity.Mulf(delta))

	if orders.fire {
		v.maybeFire(orders.fireCharge)
	}
	if orders.altFire {
		v.maybeAltFire(orders.altFireCharge)
	}
	if orders.specialFire {
		v.maybeSpecialFire(orders.specialFireCharge)
	}
}

func (v *vesselNode) maybeSpecialFire(charge float32) {
	v.weapon.Special(charge)
}

func (v *vesselNode) maybeAltFire(charge float32) {
	v.weapon.AltAttack(charge)
}

func (v *vesselNode) maybeFire(charge float32) {
	v.weapon.Attack(charge)
}
