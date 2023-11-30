package assets

import (
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/ge"

	_ "image/png"
)

func registerImageResources(ctx *ge.Context) {
	imageResources := map[resource.ImageID]resource.ImageInfo{
		ImageBoss1: {Path: "image/vessel/boss1.png", FrameWidth: 64},
		ImageBoss2: {Path: "image/vessel/boss2.png", FrameWidth: 64},

		ImageInterceptor1: {Path: "image/vessel/interceptor1.png", FrameWidth: 64},
		ImageInterceptor2: {Path: "image/vessel/interceptor2.png", FrameWidth: 64},

		ImageLaserProjectile1:          {Path: "image/projectile/laser1.png"},
		ImageLaserProjectile2:          {Path: "image/projectile/laser2.png"},
		ImageRearCannonProjectile:      {Path: "image/projectile/rear_cannon.png"},
		ImageTwinCannonProjectile:      {Path: "image/projectile/twin_cannon.png"},
		ImageTwinCannonSmallProjectile: {Path: "image/projectile/twin_cannon_small.png"},
		ImageSpinCannonProjectile:      {Path: "image/projectile/spin_cannon.png"},
		ImagePhotonCannonProjectile:    {Path: "image/projectile/photon_cannon.png"},
		ImageHomingMissile:             {Path: "image/projectile/homing_missile.png"},
		ImageMegaBomb:                  {Path: "image/projectile/mega_bomb.png"},
		ImageMine:                      {Path: "image/projectile/mine.png"},

		ImageLaserExplosion1:          {Path: "image/effect/laser1_impact.png", FrameWidth: 11},
		ImageLaserExplosion2:          {Path: "image/effect/laser2_impact.png", FrameWidth: 10},
		ImageRearCannonExplosion:      {Path: "image/effect/rear_cannon_impact.png", FrameWidth: 11},
		ImageTwinCannonExplosion:      {Path: "image/effect/twin_cannon_impact.png", FrameWidth: 32},
		ImageTwinCannonSmallExplosion: {Path: "image/effect/twin_cannon_impact_small.png", FrameWidth: 24},
		ImageSpinCannonExplosion:      {Path: "image/effect/spin_cannon_impact.png", FrameWidth: 26},
		ImageSpinCannonSpawn:          {Path: "image/effect/spin_cannon_spawn.png", FrameWidth: 24},
		ImagePhotonCannonExplosion:    {Path: "image/effect/photon_cannon_impact.png", FrameWidth: 15},
		ImageMissileExplosion:         {Path: "image/effect/missile_impact.png", FrameWidth: 24},
		ImageMissileSpawn:             {Path: "image/effect/missile_spawn.png", FrameWidth: 15},
		ImageExplosionSmoke:           {Path: "image/effect/explosion_smoke.png", FrameWidth: 32},
		ImageFireExplosion:            {Path: "image/effect/fire_explosion.png", FrameWidth: 32},
		ImageMegaBombImpact:           {Path: "image/effect/mega_bomb_impact.png", FrameWidth: 32},
		ImageTrailEffect:              {Path: "image/effect/trail.png", FrameWidth: 10},
		ImageMineSpawn:                {Path: "image/effect/mine_spawn.png", FrameWidth: 24},

		ImageDashEffect: {Path: "image/effect/dash.png", FrameWidth: 30},

		ImageTileset1: {Path: "image/landscape/tiles1.png"},

		ImageBattleOverlay:   {Path: "image/ui/battle_hud.png"},
		ImageTargetPointer:   {Path: "image/ui/pointer.png"},
		ImageBattleBarHP:     {Path: "image/ui/hp_bar.png"},
		ImageBattleBarEnergy: {Path: "image/ui/energy_bar.png"},

		ImageMenuBg: {Path: "image/ui/menu_bg.png"},

		ImageUIButtonDisabled:      {Path: "image/ebitenui/button-disabled.png"},
		ImageUIButtonIdle:          {Path: "image/ebitenui/button-idle.png"},
		ImageUIButtonHover:         {Path: "image/ebitenui/button-hover.png"},
		ImageUIButtonPressed:       {Path: "image/ebitenui/button-pressed.png"},
		ImageUISelectButtonIdle:    {Path: "image/ebitenui/select-button-idle.png"},
		ImageUISelectButtonHover:   {Path: "image/ebitenui/select-button-hover.png"},
		ImageUISelectButtonPressed: {Path: "image/ebitenui/select-button-pressed.png"},
		ImageUIPanelIdle:           {Path: "image/ebitenui/panel-idle.png"},
	}

	for id, res := range imageResources {
		ctx.Loader.ImageRegistry.Set(id, res)
		ctx.Loader.LoadImage(id)
	}
}

const (
	ImageNone resource.ImageID = iota

	ImageInterceptor1
	ImageInterceptor2

	ImageBoss1
	ImageBoss2

	ImageLaserProjectile1
	ImageLaserProjectile2
	ImageRearCannonProjectile
	ImageTwinCannonProjectile
	ImageTwinCannonSmallProjectile
	ImageSpinCannonProjectile
	ImagePhotonCannonProjectile
	ImageHomingMissile
	ImageMegaBomb
	ImageMine

	ImageLaserExplosion1
	ImageLaserExplosion2
	ImageRearCannonExplosion
	ImageTwinCannonExplosion
	ImageTwinCannonSmallExplosion
	ImageSpinCannonExplosion
	ImageSpinCannonSpawn
	ImagePhotonCannonExplosion
	ImageMissileExplosion
	ImageMissileSpawn
	ImageExplosionSmoke
	ImageFireExplosion
	ImageMegaBombImpact
	ImageTrailEffect
	ImageMineSpawn

	ImageDashEffect

	ImageTileset1

	ImageBattleOverlay
	ImageTargetPointer
	ImageBattleBarHP
	ImageBattleBarEnergy

	ImageMenuBg

	ImageUIButtonDisabled
	ImageUIButtonIdle
	ImageUIButtonHover
	ImageUIButtonPressed
	ImageUISelectButtonIdle
	ImageUISelectButtonHover
	ImageUISelectButtonPressed
	ImageUIPanelIdle
)
