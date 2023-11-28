package battle

import (
	"math"

	"github.com/quasilyte/gmath"
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

	switch w.special {
	case gamedata.HomingMissileSpecialWeapon:
		offset := homingMissileOffsets[w.specialCounter%3]
		firePos := v.pos.MoveInDirection(30, v.rotation).Add(offset.Rotated(v.rotation))
		missile := newProjectileNode(projectileConfig{
			target:    v.enemy,
			world:     v.world,
			weapon:    w.special.Base,
			pos:       firePos,
			targetPos: v.pos.MoveInDirection(w.special.Base.AttackRange, v.rotation),
			charge:    charge,
		})
		v.world.scene.AddObject(missile)
	}

	w.specialCounter++
}

func (w *weaponSystem) Attack(charge float32) {
	v := w.vessel

	switch w.design {
	case gamedata.SpinCannonWeapon:
		rotation := v.rotation + gmath.Rad(float64(w.attackCounter)*math.Pi/12)
		v.world.scene.AddObject(w.createSimpleProjectile(charge, rotation, w.design))

	case gamedata.IonCannonWeapon:
		v.world.scene.AddObject(w.createSimpleProjectile(charge, v.rotation, w.design))
		v.world.scene.AddObject(w.createSimpleProjectile(charge, v.rotation+0.3, w.design))
		v.world.scene.AddObject(w.createSimpleProjectile(charge, v.rotation-0.3, w.design))
	}

	w.attackCounter++
}

func (w *weaponSystem) AltAttack(charge float32) {
	v := w.vessel

	switch w.design {
	case gamedata.IonCannonWeapon:
		v.world.scene.AddObject(w.createSimpleProjectile(charge, v.rotation+0.15, w.design))
		v.world.scene.AddObject(w.createSimpleProjectile(charge, v.rotation-0.15, w.design))

	case gamedata.SpinCannonWeapon:
		rotation := v.rotation + gmath.Rad(float64(w.altAttackCounter)*math.Pi-math.Pi/2)
		v.world.scene.AddObject(newProjectileNode(projectileConfig{
			target:    v.enemy,
			world:     v.world,
			weapon:    w.design,
			pos:       v.pos.MoveInDirection(60, rotation),
			targetPos: v.pos.MoveInDirection(w.design.AttackRange, v.rotation),
			charge:    charge,
		}))
	}

	w.altAttackCounter++
}
