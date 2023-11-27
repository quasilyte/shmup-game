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
	cs.A = 1
	switch {
	case charge >= 0.95:
		// Red is the main color.
		cs.R = 0.05 + (1.4 * charge)
		cs.G = 0.05 + (charge * 0.7)
		cs.B = 0.05 + (charge * 0.7)
	case charge >= 0.7:
		// Blue is the main color.
		cs.R = 0.3 + (charge * 0.7)
		cs.G = 0.3 + (charge * 0.7)
		cs.B = 0.3 + (1.4 * charge)
	case charge >= 0.2:
		// Green is the main color.
		cs.R = 0.8 + (charge * 0.7)
		cs.G = 0.8 + (1.4 * charge)
		cs.B = 0.8 + (charge * 0.7)
	default:
		// It's gray.
		cs.R = 0.8 + charge
		cs.G = 0.8 + charge
		cs.B = 0.8 + charge
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
