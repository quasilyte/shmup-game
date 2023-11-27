package gamedata

import (
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/shmup-game/assets"
)

type Damage struct {
	HP float64
}

type WeaponDesign struct {
	AttackRange         float64
	ExplosionRange      float64
	ProjectileSpeed     float64
	ProjectileImage     resource.ImageID
	ProjectileExplosion resource.ImageID
	ImpactSound         resource.AudioID
}

var TestWeapon = &WeaponDesign{
	AttackRange:         320,
	ProjectileSpeed:     600,
	ExplosionRange:      10,
	ProjectileImage:     assets.ImageLaserProjectile1,
	ProjectileExplosion: assets.ImageLaserExplosion1,
	ImpactSound:         assets.AudioLaser1Impact,
}
