package gamedata

import (
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/shmup-game/assets"
)

type Damage struct {
	HP float64
}

type WeaponDesign struct {
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
	AttackRange:           280,
	ProjectileSpeed:       700,
	ExplosionRange:        12,
	ProjectileImage:       assets.ImageLaserProjectile1,
	ProjectileExplosion:   assets.ImageLaserExplosion1,
	ProjectileSpawnEffect: assets.ImageLaserExplosion1,
	ImpactSound:           assets.AudioLaser1Impact,
}

var SpinCannonWeapon = &WeaponDesign{
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
