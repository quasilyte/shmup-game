package battle

import (
	"math"
	"strconv"

	"github.com/quasilyte/ge"
	"github.com/quasilyte/ge/input"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/shmup-game/assets"
	"github.com/quasilyte/shmup-game/controls"
	"github.com/quasilyte/shmup-game/viewport"
)

type humanPlayer struct {
	world  *battleState
	input  *input.Handler
	camera *viewport.Camera
	vessel *vesselNode

	pointerRotation gmath.Rad
	pointerPos      gmath.Vec
	pointer         *ge.Sprite
	distLabelPos    gmath.Vec
	distLabel       *ge.Label
}

func (p *humanPlayer) Init(scene *ge.Scene) {
	p.updateCamera()

	p.pointer = scene.NewSprite(assets.ImageTargetPointer)
	p.pointerPos = gmath.Vec{X: 1920 / 4, Y: 1080 / 4}
	p.pointer.Rotation = &p.pointerRotation
	p.pointer.Pos.Base = &p.pointerPos
	p.camera.UI.AddGraphicsAbove(p.pointer)

	p.distLabel = ge.NewLabel(assets.BitmapFont1)
	p.distLabelPos = gmath.Vec{X: 1920/4 - 32, Y: 1080/4 - 16}
	p.distLabel.Width = 64
	p.distLabel.Height = 32
	p.distLabel.AlignHorizontal = ge.AlignHorizontalCenter
	p.distLabel.AlignVertical = ge.AlignVerticalCenter
	p.distLabel.Pos.Base = &p.distLabelPos
	p.distLabel.Text = "99"
	p.camera.UI.AddGraphicsAbove(p.distLabel)
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

func (p *humanPlayer) CameraPos() gmath.Vec {
	return p.vessel.pos.MoveInDirection(164, p.vessel.rotation)
}

func (p *humanPlayer) updateCamera() {
	p.camera.Rotation = -math.Pi/2 - p.vessel.rotation.Normalized()

	p.camera.Offset = p.CameraPos().Sub(gmath.Vec{
		X: p.camera.Rect.Width() * 0.5,
		Y: p.camera.Rect.Height() * 0.5,
	})

	if p.vessel.enemy != nil {
		enemyDist := p.camera.CenterPos().DistanceTo(p.vessel.enemy.pos)
		if enemyDist > 270 {
			p.pointerRotation = p.camera.CenterPos().AngleToPoint(p.vessel.enemy.pos) - p.vessel.rotation - math.Pi/2
			p.pointerPos = (gmath.Vec{X: 1920 / 4, Y: 1080 / 4}).MoveInDirection(252, p.pointerRotation)
			p.distLabelPos = (gmath.Vec{X: 1920/4 - 32, Y: 1080/4 - 16}).MoveInDirection(228, p.pointerRotation)
			displayDist := int((enemyDist - 270) / 4)
			p.distLabel.Text = strconv.Itoa(displayDist)
			if displayDist >= 400 {
				p.distLabel.SetColorScaleRGBA(255, 100, 100, 255)
			} else {
				p.distLabel.SetColorScaleRGBA(255, 255, 255, 255)
			}
			p.pointer.Visible = true
			p.distLabel.Visible = true
		} else {
			p.pointer.Visible = false
			p.distLabel.Visible = false
		}
	}

	// p.camera.SetOffset(p.vessel.pos.MoveInDirection(180, p.vessel.rotation).Sub(gmath.Vec{
	// 	X: p.camera.Rect.Width() * 0.5,
	// 	Y: p.camera.Rect.Height() * 0.5,
	// }))
}
