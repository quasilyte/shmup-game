package battle

import (
	"math"
	"strconv"

	"github.com/quasilyte/ge"
	"github.com/quasilyte/ge/input"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/shmup-game/assets"
	"github.com/quasilyte/shmup-game/controls"
	"github.com/quasilyte/shmup-game/gamedata"
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
	energyRegen     float64
	hints           bool
}

func (p *humanPlayer) Init(scene *ge.Scene) {
	p.updateCamera(0)

	p.pointer = scene.NewSprite(assets.ImageTargetPointer)
	p.pointerPos = gmath.Vec{X: 1920 / 4, Y: 1080 / 4}
	p.pointer.Rotation = &p.pointerRotation
	p.pointer.Pos.Base = &p.pointerPos
	p.pointer.Visible = false
	p.camera.UI.AddGraphicsAbove(p.pointer)

	p.distLabel = ge.NewLabel(assets.BitmapFont1)
	p.distLabelPos = gmath.Vec{X: 1920/4 - 32, Y: 1080/4 - 16}
	p.distLabel.Width = 64
	p.distLabel.Height = 32
	p.distLabel.AlignHorizontal = ge.AlignHorizontalCenter
	p.distLabel.AlignVertical = ge.AlignVerticalCenter
	p.distLabel.Pos.Base = &p.distLabelPos
	p.distLabel.Text = "???"
	p.distLabel.Visible = false
	p.camera.UI.AddGraphicsAbove(p.distLabel)

	{
		pos := gmath.Vec{X: 182, Y: 49}
		hpBar := newValueBar(p.camera, pos, &p.vessel.hp, p.vessel.design.HP, true)
		scene.AddObject(hpBar)
	}
	{
		pos := gmath.Vec{X: 182 + 486, Y: 49}
		energyBar := newValueBar(p.camera, pos, &p.vessel.energy, p.vessel.design.Energy, false)
		energyBar.threshold = p.vessel.weapon.special.EnergyCost
		scene.AddObject(energyBar)
	}

	p.energyRegen = p.vessel.design.Energy * 0.05

	if p.hints {
		strafeHint := ge.NewRect(scene.Context(), 196, 40)
		strafeHint.Centered = false
		strafeHint.Pos.Offset = gmath.Vec{X: 32, Y: 32}
		strafeHint.OutlineWidth = 2
		strafeHint.OutlineColorScale.SetRGBA(0, 0, 0, 0xff)
		strafeHint.FillColorScale.SetRGBA(0x01, 0x02, 0x09, 0xff)
		p.camera.UI.AddGraphicsAbove(strafeHint)

		strafeHintLabel := ge.NewLabel(assets.BitmapFont1)
		strafeHintLabel.Width = strafeHint.Width
		strafeHintLabel.Height = strafeHint.Height
		strafeHintLabel.AlignHorizontal = ge.AlignHorizontalCenter
		strafeHintLabel.AlignVertical = ge.AlignVerticalCenter
		strafeHintLabel.SetColorScaleRGBA(0x4c, 0xac, 0x3e, 0xff)
		strafeHintLabel.Text = "Hold SHIFT to strafe"
		strafeHintLabel.Pos.Offset = strafeHint.Pos.Offset
		p.camera.UI.AddGraphicsAbove(strafeHintLabel)
	}

	if p.hints {
		specialHint := ge.NewRect(scene.Context(), 196, 40)
		specialHint.Centered = false
		specialHint.Pos.Offset = gmath.Vec{X: scene.Context().ScreenWidth - 32 - 196, Y: 32}
		specialHint.OutlineWidth = 2
		specialHint.OutlineColorScale.SetRGBA(0, 0, 0, 0xff)
		specialHint.FillColorScale.SetRGBA(0x01, 0x02, 0x09, 0xff)
		p.camera.UI.AddGraphicsAbove(specialHint)

		specialHintLabel := ge.NewLabel(assets.BitmapFont1)
		specialHintLabel.Width = specialHint.Width
		specialHintLabel.Height = specialHint.Height
		specialHintLabel.AlignHorizontal = ge.AlignHorizontalCenter
		specialHintLabel.AlignVertical = ge.AlignVerticalCenter
		specialHintLabel.SetColorScaleRGBA(0x4c, 0xac, 0x3e, 0xff)
		specialHintLabel.Text = "Press SPACE to activate\nSpecial"
		specialHintLabel.Pos.Offset = specialHint.Pos.Offset
		p.camera.UI.AddGraphicsAbove(specialHintLabel)
	}
}

func (p *humanPlayer) IsDisposed() bool {
	return false
}

func (p *humanPlayer) Update(delta float64) {
	p.vessel.orders.rotateLeft = p.input.ActionIsPressed(controls.ActionRotateLeft)
	p.vessel.orders.rotateRight = p.input.ActionIsPressed(controls.ActionRotateRight)
	p.vessel.orders.turbo = p.input.ActionIsPressed(controls.ActionMoveTurbo)
	p.vessel.orders.strafe = p.input.ActionIsPressed(controls.ActionStrafe)
	p.vessel.orders.specialFire = p.input.ActionIsJustPressed(controls.ActionSpecial)

	if p.world.difficulty <= 1 {
		p.vessel.hp = gmath.ClampMax(p.vessel.hp+delta, p.vessel.design.HP)
	}

	p.updateCamera(delta)
}

func (p *humanPlayer) CameraPos() gmath.Vec {
	return p.vessel.pos.MoveInDirection(164, p.vessel.rotation)
}

func (p *humanPlayer) updateCamera(delta float64) {
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
			if displayDist <= 100 {
				p.distLabel.SetColorScaleRGBA(255, 255, 255, 255)
			} else if displayDist <= 200 {
				p.distLabel.SetColorScaleRGBA(255, 255, 100, 255)
				p.world.result.PressurePenalty += 0.05
			} else {
				p.distLabel.SetColorScaleRGBA(255, 100, 100, 255)
				p.world.result.PressurePenalty += 0.3
			}
			switch p.world.difficulty {
			case 3: // hard (take damage in red zone)
				if displayDist > 200 {
					p.vessel.OnDamage(gamedata.Damage{HP: 0.1})
				}
			case 4: // nightmare (take damage in yellow zone)
				if displayDist > 100 {
					p.vessel.OnDamage(gamedata.Damage{HP: 0.15})
				}
			}
			p.pointer.Visible = true
			p.distLabel.Visible = true
		} else {
			p.pointer.Visible = false
			p.distLabel.Visible = false
			p.vessel.energy = gmath.ClampMax(p.vessel.energy+(p.energyRegen*delta), p.vessel.design.Energy)
		}
	}

	// p.camera.SetOffset(p.vessel.pos.MoveInDirection(180, p.vessel.rotation).Sub(gmath.Vec{
	// 	X: p.camera.Rect.Width() * 0.5,
	// 	Y: p.camera.Rect.Height() * 0.5,
	// }))
}
