package battle

import (
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/shmup-game/assets"
)

func multiplyColorScale(cs ge.ColorScale, v float32) ge.ColorScale {
	cs.R *= v
	cs.G *= v
	cs.B *= v
	return cs
}

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

func playSound(world *battleState, pos gmath.Vec, id resource.AudioID) {
	if !world.human.camera.ContainsPos(pos) {
		return
	}

	numSamples := assets.NumSamples(id)
	if numSamples == 1 {
		world.scene.Audio().PlaySound(id)
	} else {
		soundIndex := world.scene.Rand().IntRange(0, numSamples-1)
		sound := resource.AudioID(int(id) + soundIndex)
		world.scene.Audio().PlaySound(sound)
	}
}
