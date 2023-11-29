package gamedata

import (
	"github.com/quasilyte/shmup-game/assets"
)

type SpecialWeaponDesign struct {
	EnergyCost float64

	Base *WeaponDesign
}

var DashSpecialWeapon = &SpecialWeaponDesign{
	EnergyCost: 20,
}

var MegaBombSpecialWeapon = &SpecialWeaponDesign{
	EnergyCost: 30,
	Base: &WeaponDesign{
		Damage:              Damage{HP: 20},
		AttackRange:         340,
		ProjectileSpeed:     400,
		ExplosionRange:      30,
		ProjectileImage:     assets.ImageMegaBomb,
		ProjectileExplosion: assets.ImageMegaBombImpact,
		IgnoreChargeColor:   true,
	},
}

var HomingMissileSpecialWeapon = &SpecialWeaponDesign{
	Base: &WeaponDesign{
		Damage:                Damage{HP: 10},
		AttackRange:           500,
		ProjectileSpeed:       250,
		ExplosionRange:        26,
		ProjectileImage:       assets.ImageHomingMissile,
		ProjectileExplosion:   assets.ImageMissileExplosion,
		ProjectileSpawnEffect: assets.ImageMissileSpawn,
		ProjectileHoming:      140,
		CollisionRange:        1,
		IgnoreChargeColor:     true,
	},
}
