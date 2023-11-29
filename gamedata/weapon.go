package gamedata

import (
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/shmup-game/assets"
)

type Damage struct {
	HP float64
}

type WeaponDesign struct {
	Damage                Damage
	AttackRange           float64
	ExplosionRange        float64
	ProjectileSpeed       float64
	ProjectileRotateSpeed float64
	ProjectileHoming      float64
	ProjectileImage       resource.ImageID
	ProjectileExplosion   resource.ImageID
	ProjectileSpawnEffect resource.ImageID
	ImpactSound           resource.AudioID
	CollisionRange        float64
	IgnoreChargeColor     bool
}

var IonCannonWeapon = &WeaponDesign{
	Damage:                Damage{HP: 10},
	AttackRange:           280,
	ProjectileSpeed:       700,
	ExplosionRange:        12,
	ProjectileImage:       assets.ImageLaserProjectile1,
	ProjectileExplosion:   assets.ImageLaserExplosion1,
	ProjectileSpawnEffect: assets.ImageLaserExplosion1,
	ImpactSound:           assets.AudioLaser1Impact,
}

var PulseLaserWeapon = &WeaponDesign{
	Damage:                Damage{HP: 5},
	AttackRange:           220,
	ProjectileSpeed:       680,
	ExplosionRange:        10,
	ProjectileImage:       assets.ImageLaserProjectile2,
	ProjectileExplosion:   assets.ImageLaserExplosion2,
	ProjectileSpawnEffect: assets.ImageLaserExplosion2,
	ImpactSound:           assets.AudioLaser1Impact,
	CollisionRange:        2,
}

var RearCannonWeapon = &WeaponDesign{
	Damage:                Damage{HP: 7},
	AttackRange:           240,
	ProjectileSpeed:       800,
	ExplosionRange:        4,
	ProjectileImage:       assets.ImageRearCannonProjectile,
	ProjectileExplosion:   assets.ImageRearCannonExplosion,
	ProjectileSpawnEffect: assets.ImageRearCannonExplosion,
	ImpactSound:           assets.AudioLaser1Impact,
	CollisionRange:        2,
}

var SpinCannonWeapon = &WeaponDesign{
	Damage:                Damage{HP: 15},
	AttackRange:           420,
	ProjectileSpeed:       280,
	ExplosionRange:        20,
	ProjectileImage:       assets.ImageSpinCannonProjectile,
	ProjectileExplosion:   assets.ImageSpinCannonExplosion,
	ProjectileSpawnEffect: assets.ImageSpinCannonSpawn,
	ProjectileRotateSpeed: 26,
	CollisionRange:        4,
	// ImpactSound:         assets.AudioLaser1Impact,
}
