package battle

import (
	"math"

	"github.com/quasilyte/ge"
	"github.com/quasilyte/ge/input"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/shmup-game/controls"
	"github.com/quasilyte/shmup-game/viewport"
)

type humanPlayer struct {
	world  *battleState
	input  *input.Handler
	camera *viewport.Camera
	vessel *vesselNode
}

func (p *humanPlayer) Init(scene *ge.Scene) {
	p.updateCamera()
}

func (p *humanPlayer) IsDisposed() bool {
	return false
}

func (p *humanPlayer) Update(delta float64) {
	p.vessel.orders.rotateLeft = p.input.ActionIsPressed(controls.ActionRotateLeft)
	p.vessel.orders.rotateRight = p.input.ActionIsPressed(controls.ActionRotateRight)
	p.vessel.orders.turbo = p.input.ActionIsPressed(controls.ActionMoveTurbo)
	p.vessel.orders.strafe = p.input.ActionIsPressed(controls.ActionStrafe)

	p.updateCamera()
}

func (p *humanPlayer) updateCamera() {
	p.camera.Rotation = -math.Pi/2 - p.vessel.rotation.Normalized()

	p.camera.SetOffset(p.vessel.pos.MoveInDirection(164, p.vessel.rotation).Sub(gmath.Vec{
		X: p.camera.Rect.Width() * 0.5,
		Y: p.camera.Rect.Height() * 0.5,
	}))

	// p.camera.SetOffset(p.vessel.pos.MoveInDirection(180, p.vessel.rotation).Sub(gmath.Vec{
	// 	X: p.camera.Rect.Width() * 0.5,
	// 	Y: p.camera.Rect.Height() * 0.5,
	// }))
}
