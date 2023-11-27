package battle

import (
	"math"

	"github.com/quasilyte/gmath"
	"github.com/quasilyte/shmup-game/gamedata"
)

type weaponSystem struct {
	vessel *vesselNode
	design *gamedata.WeaponDesign

	attackCounter    int
	altAttackCounter int
}

func newWeaponSystem(world *battleState, vessel *vesselNode, design *gamedata.WeaponDesign) *weaponSystem {
	return &weaponSystem{
		vessel: vessel,
		design: design,
	}
}

func (w *weaponSystem) createSimpleProjectile(charge float32, rotation gmath.Rad) *projectileNode {
	v := w.vessel
	return newProjectileNode(projectileConfig{
		target:    v.enemy,
		world:     v.world,
		weapon:    w.design,
		pos:       v.pos.MoveInDirection(30, rotation),
		targetPos: v.pos.MoveInDirection(w.design.AttackRange, rotation),
		charge:    charge,
	})
}

func (w *weaponSystem) Attack(charge float32) {
	v := w.vessel

	switch w.design {
	case gamedata.SpinCannonWeapon:
		rotation := v.rotation + gmath.Rad(float64(w.attackCounter)*math.Pi/12)
		v.world.scene.AddObject(w.createSimpleProjectile(charge, rotation))

	case gamedata.IonCannonWeapon:
		v.world.scene.AddObject(w.createSimpleProjectile(charge, v.rotation))
		v.world.scene.AddObject(w.createSimpleProjectile(charge, v.rotation+0.3))
		v.world.scene.AddObject(w.createSimpleProjectile(charge, v.rotation-0.3))
	}

	w.attackCounter++
}

func (w *weaponSystem) AltAttack(charge float32) {
	v := w.vessel

	switch w.design {
	case gamedata.IonCannonWeapon:
		v.world.scene.AddObject(w.createSimpleProjectile(charge, v.rotation+0.15))
		v.world.scene.AddObject(w.createSimpleProjectile(charge, v.rotation-0.15))

	case gamedata.SpinCannonWeapon:
		rotation := v.rotation + gmath.Rad(float64(w.altAttackCounter)*math.Pi/2)
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
