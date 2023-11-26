package battle

import (
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/ge"
	"github.com/quasilyte/ge/gesignal"
	"github.com/quasilyte/gmath"
)

type effectLayer int

const (
	normalEffectLayer effectLayer = iota
	slightlyAboveEffectLayer
	aboveEffectLayer
)

type effectSpeed int

const (
	normalEffectSpeed effectSpeed = iota
	fastEffectSpeed
	veryFastEffectSpeed
)

type effectNode struct {
	pos        gmath.Vec
	image      resource.ImageID
	anim       *ge.Animation
	layer      effectLayer
	speed      effectSpeed
	colorScale ge.ColorScale
	world      *battleState
	rotates    bool
	noFlip     bool

	rotation gmath.Rad

	EventCompleted gesignal.Event[gesignal.Void]
}

type effectConfig struct {
	world    *battleState
	pos      gmath.Vec
	layer    effectLayer
	speed    effectSpeed
	image    resource.ImageID
	rotation gmath.Rad
}

func newEffectNode(config effectConfig) *effectNode {
	return &effectNode{
		pos:        config.pos,
		image:      config.image,
		layer:      config.layer,
		speed:      config.speed,
		world:      config.world,
		colorScale: ge.ColorScale{1, 1, 1, 1},
		rotation:   config.rotation,
	}
}

func (e *effectNode) Init(scene *ge.Scene) {
	var sprite *ge.Sprite
	if e.anim == nil {
		sprite = scene.NewSprite(e.image)
		sprite.Pos.Base = &e.pos
	} else {
		sprite = e.anim.Sprite()
	}
	sprite.Rotation = &e.rotation
	sprite.SetColorScale(e.colorScale)
	if !e.noFlip {
		sprite.FlipHorizontal = scene.Rand().Bool()
	}
	switch e.layer {
	case aboveEffectLayer:
		e.world.stage.AddGraphicsAbove(sprite)
	case slightlyAboveEffectLayer:
		e.world.stage.AddGraphicsSlightlyAbove(sprite)
	default:
		e.world.stage.AddGraphics(sprite)
	}
	if e.anim == nil {
		e.anim = ge.NewAnimation(sprite, -1)
	}
	switch e.speed {
	case fastEffectSpeed:
		e.anim.SetSecondsPerFrame(0.035)
	case veryFastEffectSpeed:
		e.anim.SetSecondsPerFrame(0.025)
	}
}

func (e *effectNode) IsDisposed() bool {
	return e.anim.IsDisposed()
}

func (e *effectNode) Dispose() {
	e.anim.Sprite().Dispose()
}

func (e *effectNode) Update(delta float64) {
	if e.anim.Tick(delta) {
		e.EventCompleted.Emit(gesignal.Void{})
		e.Dispose()
		return
	}
	if e.rotates {
		e.rotation += gmath.Rad(delta * 2)
	}
}
