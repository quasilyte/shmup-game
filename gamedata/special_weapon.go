package gamedata

import (
	"github.com/quasilyte/shmup-game/assets"
)

type SpecialWeaponDesign struct {
	Base *WeaponDesign
}

var HomingMissileSpecialWeapon = &SpecialWeaponDesign{
	Base: &WeaponDesign{
		AttackRange:         500,
		ProjectileSpeed:     250,
		ExplosionRange:      26,
		ProjectileImage:     assets.ImageHomingMissile,
		ProjectileExplosion: assets.ImageMissileExplosion,
		ProjectileHoming:    140,
		CanCollide:          true,
		IgnoreChargeColor:   true,
	},
}
