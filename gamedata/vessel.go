package gamedata

import (
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/shmup-game/assets"
)

type VesselDesign struct {
	Image resource.ImageID

	HP     float64
	Energy float64

	Size float64

	Acceleration     float64
	Speed            float64
	StrafeSpeed      float64
	RotationMaxSpeed gmath.Rad
}

var InterceptorDesign1 = &VesselDesign{
	Image: assets.ImageInterceptor1,

	HP:     100,
	Energy: 100,
	Size:   16,

	Acceleration:     150,
	Speed:            320,
	StrafeSpeed:      300,
	RotationMaxSpeed: 3.5,
}

var InterceptorDesign2 = &VesselDesign{
	Image: assets.ImageInterceptor2,

	HP:     130,
	Energy: 80,
	Size:   20,

	Acceleration:     200,
	Speed:            280,
	StrafeSpeed:      250,
	RotationMaxSpeed: 3.0,
}

var BossVessel1 = &VesselDesign{
	Image: assets.ImageBoss1,

	HP:   500,
	Size: 26,

	Acceleration:     100,
	Speed:            120,
	StrafeSpeed:      120,
	RotationMaxSpeed: 2.5,
}
