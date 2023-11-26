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

	sprite *ge.Sprite

	speedLevel int

	orders vesselOrders

	disposed bool
}

type vesselOrders struct {
	accelerate  bool
	decelerate  bool
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
	return 50 * float64(v.speedLevel+1)
}

func (v *vesselNode) movementSpeed() float64 {
	return 100 * float64(v.speedLevel)
}

func (v *vesselNode) rotationSpeed() gmath.Rad {
	return 2.5 + gmath.Rad(float64(v.speedLevel)*0.25)
}

func (v *vesselNode) Update(delta float64) {
	orders := v.orders
	v.orders = vesselOrders{}

	strafing := false

	switch {
	case orders.rotateLeft == orders.rotateRight:
		// Do nothing.
		v.sprite.FrameOffset.X = 0
	case orders.rotateLeft:
		if orders.strafe {
			strafing = true
			v.pos = v.pos.MoveInDirection(v.strafeSpeed()*delta, v.rotation-math.Pi/2)
		} else {
			v.rotation -= v.rotationSpeed() * gmath.Rad(delta)
		}
		v.sprite.FrameOffset.X = 2 * v.sprite.FrameWidth
	case orders.rotateRight:
		if orders.strafe {
			strafing = true
			v.pos = v.pos.MoveInDirection(v.strafeSpeed()*delta, v.rotation+math.Pi/2)
		} else {
			v.rotation += v.rotationSpeed() * gmath.Rad(delta)
		}
		v.sprite.FrameOffset.X = 1 * v.sprite.FrameWidth
	}

	speed := v.movementSpeed()
	if strafing {
		speed *= 0.2
	}
	v.pos = v.pos.MoveInDirection(speed*delta, v.rotation)

	switch {
	case orders.accelerate == orders.decelerate:
		// Do nothing.
	case orders.accelerate:
		v.speedLevel = gmath.ClampMax(v.speedLevel+1, 5)
	case orders.decelerate:
		v.speedLevel = gmath.ClampMin(v.speedLevel-1, 0)
	}

	if orders.fire {
		v.maybeFire(orders.fireCharge)
	}
	if orders.altFire {
		v.maybeAltFire(orders.altFireCharge)
	}
}

func (v *vesselNode) maybeAltFire(charge float32) {
	v.scene.AddObject(newProjectileNode(projectileConfig{
		extraSpeed: v.movementSpeed(),
		world:      v.world,
		weapon:     gamedata.TestWeapon,
		pos:        v.pos.MoveInDirection(30, v.rotation+0.15),
		targetPos:  v.pos.MoveInDirection(300, v.rotation+0.15),
		charge:     charge,
	}))

	v.scene.AddObject(newProjectileNode(projectileConfig{
		extraSpeed: v.movementSpeed(),
		world:      v.world,
		weapon:     gamedata.TestWeapon,
		pos:        v.pos.MoveInDirection(30, v.rotation-0.15),
		targetPos:  v.pos.MoveInDirection(300, v.rotation-0.15),
		charge:     charge,
	}))
}

func (v *vesselNode) maybeFire(charge float32) {
	projectile := newProjectileNode(projectileConfig{
		extraSpeed: v.movementSpeed(),
		world:      v.world,
		weapon:     gamedata.TestWeapon,
		pos:        v.pos.MoveInDirection(30, v.rotation),
		targetPos:  v.pos.MoveInDirection(300, v.rotation),
		charge:     charge,
	})
	v.scene.AddObject(projectile)
}
