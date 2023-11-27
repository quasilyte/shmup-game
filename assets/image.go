package assets

import (
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/ge"

	_ "image/png"
)

func registerImageResources(ctx *ge.Context) {
	imageResources := map[resource.ImageID]resource.ImageInfo{
		ImageBoss1: {Path: "image/vessel/boss1.png", FrameWidth: 64},

		ImageInterceptor1: {Path: "image/vessel/interceptor1.png", FrameWidth: 64},

		ImageLaserProjectile1:     {Path: "image/projectile/laser1.png"},
		ImageSpinCannonProjectile: {Path: "image/projectile/spin_cannon.png"},

		ImageLaserExplosion1:     {Path: "image/effect/laser1_impact.png", FrameWidth: 11},
		ImageSpinCannonExplosion: {Path: "image/effect/spin_cannon_impact.png", FrameWidth: 26},

		ImageDashEffect: {Path: "image/effect/dash.png", FrameWidth: 30},

		ImageTileset1: {Path: "image/landscape/tiles1.png"},

		ImageBattleOverlay: {Path: "image/ui/battle_hud.png"},
	}

	for id, res := range imageResources {
		ctx.Loader.ImageRegistry.Set(id, res)
		ctx.Loader.LoadImage(id)
	}
}

const (
	ImageNone resource.ImageID = iota

	ImageInterceptor1

	ImageBoss1

	ImageLaserProjectile1
	ImageSpinCannonProjectile

	ImageLaserExplosion1
	ImageSpinCannonExplosion

	ImageDashEffect

	ImageTileset1

	ImageBattleOverlay
)
