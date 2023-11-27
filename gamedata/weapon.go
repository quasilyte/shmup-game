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
	ProjectileImage       resource.ImageID
	ProjectileExplosion   resource.ImageID
	ImpactSound           resource.AudioID
}

var IonCannonWeapon = &WeaponDesign{
	AttackRange:         320,
	ProjectileSpeed:     700,
	ExplosionRange:      10,
	ProjectileImage:     assets.ImageLaserProjectile1,
	ProjectileExplosion: assets.ImageLaserExplosion1,
	ImpactSound:         assets.AudioLaser1Impact,
}

var SpinCannonWeapon = &WeaponDesign{
	AttackRange:           260,
	ProjectileSpeed:       400,
	ExplosionRange:        20,
	ProjectileImage:       assets.ImageSpinCannonProjectile,
	ProjectileExplosion:   assets.ImageSpinCannonExplosion,
	ProjectileRotateSpeed: 26,
	// ImpactSound:         assets.AudioLaser1Impact,
}
