package battle

import (
	"math"

	"github.com/quasilyte/gmath"
)

func correctedPos(sector gmath.Rect, pos gmath.Vec, pad float64) gmath.Vec {
	if pos.X < (pad + sector.Min.X) {
		pos.X = pad + sector.Min.X
	} else if pos.X > (sector.Max.X - pad) {
		pos.X = sector.Max.X - pad
	}
	if pos.Y < (pad + sector.Min.Y) {
		pos.Y = pad + sector.Min.Y
	} else if pos.Y > (sector.Max.Y - pad) {
		pos.Y = sector.Max.Y - pad
	}
	return pos
}

func randomRectPos(rng *gmath.Rand, sector gmath.Rect) gmath.Vec {
	return gmath.Vec{
		X: rng.FloatRange(sector.Min.X, sector.Max.X),
		Y: rng.FloatRange(sector.Min.Y, sector.Max.Y),
	}
}

func angleDelta(a, b gmath.Rad) gmath.Rad {
	d1 := gmath.RadToVec(a)
	d2 := gmath.RadToVec(b)
	cross := d1.X*d2.Y - d1.Y*d2.X
	if cross == -0 {
		return 0
	}
	if cross < 0 {
		return gmath.Rad(math.Acos(d1.Dot(d2)))
	}
	if cross > 0 {
		return gmath.Rad(-math.Acos(d1.Dot(d2)))
	}
	return -math.Pi
}
