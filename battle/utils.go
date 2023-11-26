package battle

import (
	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
)

func calculateColorScale(charge float32) ge.ColorScale {
	if charge == 1 {
		return ge.ColorScale{1, 1, 1, 1}
	}
	var cs ge.ColorScale
	cs.A = float32(gmath.ClampMax(float64(charge)*1.5, 1))
	cs.G = 1 + charge*0.5
	cs.R = 1 - charge*0.5
	cs.B = 3 * (1 - charge)
	return cs
}
