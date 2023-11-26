package battle

import (
	"math"

	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/shmup-game/assets"
	"github.com/quasilyte/shmup-game/gamedata"
)

type vesselNode struct {
	rotation gmath.Rad
	pos      gmath.Vec

	world *battleState
	scene *ge.Scene

	strafing  bool
	thrusting bool

	velocity gmath.Vec

	sprite *ge.Sprite

	orders vesselOrders

	disposed bool
}

type vesselOrders struct {
	turbo       bool
	strafe      bool
	rotateLeft  bool
	rotateRight bool

	fire          bool
	altFire       bool
	fireCharge    float32
	altFireCharge float32
}

func newVesselNode(world *battleState) *vesselNode {
	return &vesselNode{world: world}
}

func (v *vesselNode) Init(scene *ge.Scene) {
	v.scene = scene
	v.sprite = scene.NewSprite(assets.ImageInterceptor1)
	v.sprite.Pos.Base = &v.pos
	v.sprite.Rotation = &v.rotation
	v.world.stage.AddSpriteAbove(v.sprite)
}

func (v *vesselNode) IsDisposed() bool {
	return v.disposed
}

func (v *vesselNode) strafeSpeed() float64 {
	return 300
}

func (v *vesselNode) acceleration() float64 {
	return 400
}

func (v *vesselNode) maxSpeed() float64 {
	return 400
}

func (v *vesselNode) rotationSpeed() gmath.Rad {
	return 3.5
}

func (v *vesselNode) Update(delta float64) {
	orders := v.orders
	v.orders = vesselOrders{}

	wasStrafing := v.strafing
	v.strafing = false

	wasThrusting := v.thrusting
	v.thrusting = false

	switch {
	case orders.rotateLeft == orders.rotateRight:
		// Do nothing.
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
			v.rotation -= v.rotationSpeed() * gmath.Rad(delta)
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
			v.rotation += v.rotationSpeed() * gmath.Rad(delta)
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
}

func (v *vesselNode) maybeAltFire(charge float32) {
	v.scene.AddObject(newProjectileNode(projectileConfig{
		// extraSpeed: v.movementSpeed(),
		world:     v.world,
		weapon:    gamedata.TestWeapon,
		pos:       v.pos.MoveInDirection(30, v.rotation+0.15),
		targetPos: v.pos.MoveInDirection(300, v.rotation+0.15),
		charge:    charge,
	}))

	v.scene.AddObject(newProjectileNode(projectileConfig{
		// extraSpeed: v.movementSpeed(),
		world:     v.world,
		weapon:    gamedata.TestWeapon,
		pos:       v.pos.MoveInDirection(30, v.rotation-0.15),
		targetPos: v.pos.MoveInDirection(300, v.rotation-0.15),
		charge:    charge,
	}))
}

func (v *vesselNode) maybeFire(charge float32) {
	projectile := newProjectileNode(projectileConfig{
		// extraSpeed: v.movementSpeed(),
		world:     v.world,
		weapon:    gamedata.TestWeapon,
		pos:       v.pos.MoveInDirection(30, v.rotation),
		targetPos: v.pos.MoveInDirection(300, v.rotation),
		charge:    charge,
	})
	v.scene.AddObject(projectile)
}
