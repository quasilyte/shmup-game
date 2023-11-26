package assets

import (
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/ge"

	_ "image/png"
)

func registerImageResources(ctx *ge.Context) {
	imageResources := map[resource.ImageID]resource.ImageInfo{
		ImageInterceptor1: {Path: "image/vessel/interceptor1.png", FrameWidth: 64},

		ImageLaserProjectile1: {Path: "image/projectile/laser1.png"},

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

	ImageLaserProjectile1

	ImageTileset1

	ImageBattleOverlay
)
