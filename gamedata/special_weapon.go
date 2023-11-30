package gamedata

import (
	"github.com/quasilyte/shmup-game/assets"
)

type SpecialWeaponDesign struct {
	EnergyCost float64

	Base *WeaponDesign
}

var DashSpecialWeapon = &SpecialWeaponDesign{
	EnergyCost: 15,
}

var SpinningShieldSpecialWeapon = &SpecialWeaponDesign{
	EnergyCost: 45,
}

var MegaBombSpecialWeapon = &SpecialWeaponDesign{
	EnergyCost: 25,
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

var MineSpecialWeapon = &SpecialWeaponDesign{
	Base: &WeaponDesign{
		Damage:                Damage{HP: 30},
		AttackRange:           300,
		ProjectileSpeed:       20,
		ExplosionRange:        6,
		ProjectileImage:       assets.ImageMine,
		ProjectileSpawnEffect: assets.ImageMineSpawn,
		CollisionRange:        5,
		ProjectileRotateSpeed: 1,
		IgnoreChargeColor:     true,
	},
}
