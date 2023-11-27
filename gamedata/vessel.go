package gamedata

import (
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/shmup-game/assets"
)

type VesselDesign struct {
	Image resource.ImageID

	HP float64

	Size float64

	Acceleration  float64
	Speed         float64
	StrafeSpeed   float64
	RotationSpeed gmath.Rad
}

var InterceptorDesign1 = &VesselDesign{
	Image: assets.ImageInterceptor1,

	HP:   100,
	Size: 20,

	Acceleration:  200,
	Speed:         400,
	StrafeSpeed:   300,
	RotationSpeed: 4,
}

var BossVessel1 = &VesselDesign{
	Image: assets.ImageBoss1,

	HP:   500,
	Size: 28,

	Acceleration:  120,
	Speed:         200,
	StrafeSpeed:   200,
	RotationSpeed: 3,
}
