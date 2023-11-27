package battle

import (
	"github.com/quasilyte/shmup-game/gamedata"
)

type weaponSystem struct {
	vessel *vesselNode
	design *gamedata.WeaponDesign
}

func newWeaponSystem(world *battleState, vessel *vesselNode, design *gamedata.WeaponDesign) *weaponSystem {
	return &weaponSystem{
		vessel: vessel,
		design: design,
	}
}

func (w *weaponSystem) Attack(charge float32) {
	v := w.vessel

	switch w.design {
	case gamedata.IonCannonWeapon:
		v.world.scene.AddObject(newProjectileNode(projectileConfig{
			target:    v.enemy,
			world:     v.world,
			weapon:    gamedata.IonCannonWeapon,
			pos:       v.pos.MoveInDirection(30, v.rotation),
			targetPos: v.pos.MoveInDirection(300, v.rotation),
			charge:    charge,
		}))
		v.scene.AddObject(newProjectileNode(projectileConfig{
			target:    v.enemy,
			world:     v.world,
			weapon:    gamedata.IonCannonWeapon,
			pos:       v.pos.MoveInDirection(30, v.rotation+0.30),
			targetPos: v.pos.MoveInDirection(300, v.rotation+0.30),
			charge:    charge,
		}))
		v.scene.AddObject(newProjectileNode(projectileConfig{
			target:    v.enemy,
			world:     v.world,
			weapon:    gamedata.IonCannonWeapon,
			pos:       v.pos.MoveInDirection(30, v.rotation-0.30),
			targetPos: v.pos.MoveInDirection(300, v.rotation-0.30),
			charge:    charge,
		}))
	}
}

func (w *weaponSystem) AltAttack(charge float32) {
	v := w.vessel

	switch w.design {
	case gamedata.IonCannonWeapon:
		v.scene.AddObject(newProjectileNode(projectileConfig{
			target:    v.enemy,
			world:     v.world,
			weapon:    gamedata.IonCannonWeapon,
			pos:       v.pos.MoveInDirection(30, v.rotation+0.15),
			targetPos: v.pos.MoveInDirection(300, v.rotation+0.15),
			charge:    charge,
		}))
		v.scene.AddObject(newProjectileNode(projectileConfig{
			target:    v.enemy,
			world:     v.world,
			weapon:    gamedata.IonCannonWeapon,
			pos:       v.pos.MoveInDirection(30, v.rotation-0.15),
			targetPos: v.pos.MoveInDirection(300, v.rotation-0.15),
			charge:    charge,
		}))
	}
}
