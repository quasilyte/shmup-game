package gamedata

import (
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/shmup-game/assets"
)

type WeaponDesign struct {
	AttackRange         float64
	ProjectileSpeed     float64
	ProjectileImage     resource.ImageID
	ProjectileExplosion resource.ImageID
}

var TestWeapon = &WeaponDesign{
	AttackRange:         320,
	ProjectileSpeed:     600,
	ProjectileImage:     assets.ImageLaserProjectile1,
	ProjectileExplosion: assets.ImageLaserExplosion1,
}
