package battle

import (
	"math"

	"github.com/quasilyte/gmath"
	"github.com/quasilyte/shmup-game/assets"
	"github.com/quasilyte/shmup-game/gamedata"
)

var homingMissileOffsets = []gmath.Vec{
	{Y: -10},
	{},
	{Y: 10},
}

type weaponSystem struct {
	vessel  *vesselNode
	design  *gamedata.WeaponDesign
	special *gamedata.SpecialWeaponDesign

	primaryCounter   int
	attackCounter    int
	altAttackCounter int
	specialCounter   int
}

func newWeaponSystem(world *battleState, vessel *vesselNode, design *gamedata.WeaponDesign, special *gamedata.SpecialWeaponDesign) *weaponSystem {
	return &weaponSystem{
		vessel:  vessel,
		design:  design,
		special: special,
	}
}

func (w *weaponSystem) createSimpleProjectile(charge float32, rotation gmath.Rad, design *gamedata.WeaponDesign) *projectileNode {
	v := w.vessel
	return newProjectileNode(projectileConfig{
		target:    v.enemy,
		world:     v.world,
		weapon:    design,
		pos:       v.pos.MoveInDirection(30, rotation),
		targetPos: v.pos.MoveInDirection(design.AttackRange, rotation),
		charge:    charge,
	})
}

func (w *weaponSystem) Special(charge float32) {
	v := w.vessel

	fireRotation := v.getFireRotation()

	if w.special.EnergyCost > 0 {
		if v.energy < w.special.EnergyCost {
			return
		}
		v.energy -= w.special.EnergyCost
	}

	switch w.special {
	case gamedata.HomingMissileSpecialWeapon:
		offset := homingMissileOffsets[w.specialCounter%3]
		firePos := v.pos.MoveInDirection(30, fireRotation).Add(offset.Rotated(fireRotation))
		missile := newProjectileNode(projectileConfig{
			target:    v.enemy,
			world:     v.world,
			weapon:    w.special.Base,
			pos:       firePos,
			targetPos: v.pos.MoveInDirection(w.special.Base.AttackRange, fireRotation),
			charge:    charge,
		})
		v.scene.AddObject(missile)

	case gamedata.MegaBombSpecialWeapon:
		v.scene.AddObject(w.createSimpleProjectile(1, fireRotation, w.special.Base))

	case gamedata.SpinningShieldSpecialWeapon:
		v.startSpinning(1)

	case gamedata.DashSpecialWeapon:
		v.velocity = v.velocity.Add(gmath.RadToVec(fireRotation - math.Pi).Mulf(200)).ClampLen(v.maxSpeed())
		v.scene.AddObject(newEffectNode(effectConfig{
			world:    v.world,
			pos:      v.pos.MoveInDirection(-20, fireRotation+math.Pi),
			layer:    slightlyAboveEffectLayer,
			speed:    veryFastEffectSpeed,
			image:    assets.ImageDashEffect,
			rotation: fireRotation + math.Pi/2,
		}))
	}

	w.specialCounter++
}

func (w *weaponSystem) Attack(charge float32) {
	v := w.vessel

	fireRotation := v.getFireRotation()

	switch w.design {
	case gamedata.SpinCannonWeapon:
		rotation := fireRotation + gmath.Rad(float64(w.attackCounter)*math.Pi/12)
		v.world.scene.AddObject(w.createSimpleProjectile(charge, rotation, w.design))

	case gamedata.IonCannonWeapon:
		v.world.scene.AddObject(w.createSimpleProjectile(charge, fireRotation, w.design))
		v.world.scene.AddObject(w.createSimpleProjectile(charge, fireRotation+0.3, w.design))
		v.world.scene.AddObject(w.createSimpleProjectile(charge, fireRotation-0.3, w.design))

	case gamedata.PulseLaserWeapon:
		v.world.scene.AddObject(w.createSimpleProjectile(charge, fireRotation, w.design))

	case gamedata.RearCannonWeapon:
		v.world.scene.AddObject(w.createSimpleProjectile(charge, fireRotation+math.Pi/2, w.design))
		v.world.scene.AddObject(w.createSimpleProjectile(charge, fireRotation+0.3+math.Pi/2, w.design))
		v.world.scene.AddObject(w.createSimpleProjectile(charge, fireRotation-0.3+math.Pi/2, w.design))
		v.world.scene.AddObject(w.createSimpleProjectile(charge, fireRotation-math.Pi/2, w.design))
		v.world.scene.AddObject(w.createSimpleProjectile(charge, fireRotation+0.3-math.Pi/2, w.design))
		v.world.scene.AddObject(w.createSimpleProjectile(charge, fireRotation-0.3-math.Pi/2, w.design))

	case gamedata.TwinCannonWeapon:
		if w.primaryCounter%24 > 11 {
			firePos := v.pos.MoveInDirection(10, fireRotation).Add((gmath.Vec{Y: 38}).Rotated(fireRotation))
			v.world.scene.AddObject(newProjectileNode(projectileConfig{
				target:    v.enemy,
				world:     v.world,
				weapon:    w.design,
				pos:       firePos,
				targetPos: firePos.MoveInDirection(w.design.AttackRange-20, fireRotation),
				charge:    charge,
			}))
		} else {
			firePos := v.pos.MoveInDirection(10, fireRotation).Add((gmath.Vec{Y: -38}).Rotated(fireRotation))
			v.world.scene.AddObject(newProjectileNode(projectileConfig{
				target:    v.enemy,
				world:     v.world,
				weapon:    w.design,
				pos:       firePos,
				targetPos: firePos.MoveInDirection(w.design.AttackRange-20, fireRotation),
				charge:    charge,
			}))
		}
	}

	w.attackCounter++
	w.primaryCounter++
}

func (w *weaponSystem) AltAttack(charge float32) {
	v := w.vessel

	fireRotation := v.getFireRotation()

	switch w.design {
	case gamedata.IonCannonWeapon:
		v.world.scene.AddObject(w.createSimpleProjectile(charge, fireRotation+0.15, w.design))
		v.world.scene.AddObject(w.createSimpleProjectile(charge, fireRotation-0.15, w.design))

	case gamedata.PulseLaserWeapon:
		firePos1 := v.pos.MoveInDirection(30, fireRotation).Add((gmath.Vec{Y: 16}).Rotated(fireRotation))
		firePos2 := v.pos.MoveInDirection(30, fireRotation).Add((gmath.Vec{Y: -16}).Rotated(fireRotation))
		for _, firePos := range [2]gmath.Vec{firePos1, firePos2} {
			v.world.scene.AddObject(newProjectileNode(projectileConfig{
				target:    v.enemy,
				world:     v.world,
				weapon:    w.design,
				pos:       firePos,
				targetPos: v.pos.MoveInDirection(w.design.AttackRange, fireRotation),
				charge:    charge,
			}))
		}

	case gamedata.RearCannonWeapon:
		v.world.scene.AddObject(w.createSimpleProjectile(charge, fireRotation+0.3+math.Pi/2+0.25, w.design))
		v.world.scene.AddObject(w.createSimpleProjectile(charge, fireRotation-0.3+math.Pi/2-0.25, w.design))
		v.world.scene.AddObject(w.createSimpleProjectile(charge, fireRotation+0.3+math.Pi/2+0.5, w.design))
		v.world.scene.AddObject(w.createSimpleProjectile(charge, fireRotation-0.3+math.Pi/2-0.5, w.design))
		v.world.scene.AddObject(w.createSimpleProjectile(charge, fireRotation+0.3-math.Pi/2+0.25, w.design))
		v.world.scene.AddObject(w.createSimpleProjectile(charge, fireRotation-0.3-math.Pi/2-0.25, w.design))
		v.world.scene.AddObject(w.createSimpleProjectile(charge, fireRotation+0.3-math.Pi/2+0.5, w.design))
		v.world.scene.AddObject(w.createSimpleProjectile(charge, fireRotation-0.3-math.Pi/2-0.5, w.design))

	case gamedata.SpinCannonWeapon:
		rotation := fireRotation + gmath.Rad(float64(w.altAttackCounter)*math.Pi-math.Pi/2)
		v.world.scene.AddObject(newProjectileNode(projectileConfig{
			target:    v.enemy,
			world:     v.world,
			weapon:    w.design,
			pos:       v.pos.MoveInDirection(60, rotation),
			targetPos: v.pos.MoveInDirection(w.design.AttackRange, fireRotation),
			charge:    charge,
		}))

	case gamedata.TwinCannonWeapon:
		if w.primaryCounter%24 > 11 {
			firePos := v.pos.MoveInDirection(10, fireRotation).Add((gmath.Vec{Y: 46}).Rotated(fireRotation))
			v.world.scene.AddObject(newProjectileNode(projectileConfig{
				target:    v.enemy,
				world:     v.world,
				weapon:    gamedata.TwinCannonSmallWeapon,
				pos:       firePos,
				targetPos: firePos.MoveInDirection(w.design.AttackRange-20, fireRotation),
				charge:    charge,
			}))
		} else {
			firePos := v.pos.MoveInDirection(10, fireRotation).Add((gmath.Vec{Y: -46}).Rotated(fireRotation))
			v.world.scene.AddObject(newProjectileNode(projectileConfig{
				target:    v.enemy,
				world:     v.world,
				weapon:    gamedata.TwinCannonSmallWeapon,
				pos:       firePos,
				targetPos: firePos.MoveInDirection(w.design.AttackRange-20, fireRotation),
				charge:    charge,
			}))
		}
	}

	w.altAttackCounter++
	w.primaryCounter++
}
