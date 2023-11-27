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
	var cs ge.ColorScale
	cs.A = float32(gmath.ClampMax(float64(charge)*1.5, 1))
	switch {
	case charge >= 0.95:
		// Red is the main color.
		cs.R = 1.4 * charge
		cs.G = charge * 0.6
		cs.B = charge * 0.6
	case charge >= 0.7:
		// Blue is the main color.
		cs.R = charge * 0.6
		cs.G = charge * 0.6
		cs.B = 1.4 * charge
	case charge >= 0.2:
		// Green is the main color.
		cs.R = charge * 0.6
		cs.G = 1.4 * charge
		cs.B = charge * 0.6
	default:
		// It's dark gray.
		cs.R = charge * 0.3
		cs.G = charge * 0.3
		cs.B = charge * 0.3
	}
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
