package battle

import (
	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/shmup-game/assets"
)

type vesselTrailNode struct {
	world    *battleState
	sprite   *ge.Sprite
	frame    int
	pos      gmath.Vec
	rotation gmath.Rad
}

func newVesselTrailNode(world *battleState, frame int, pos gmath.Vec, rotation gmath.Rad) *vesselTrailNode {
	return &vesselTrailNode{
		world:    world,
		frame:    frame,
		pos:      pos,
		rotation: rotation,
	}
}

func (t *vesselTrailNode) Init(scene *ge.Scene) {
	t.sprite = scene.NewSprite(assets.ImageTrailEffect)
	t.sprite.Pos.Base = &t.pos
	t.sprite.Rotation = &t.rotation
	t.sprite.FrameOffset.X = float64(t.frame) * t.sprite.FrameWidth
	t.world.stage.AddSpriteBelow(t.sprite)
}

func (t *vesselTrailNode) IsDisposed() bool {
	return t.sprite.IsDisposed()
}

func (t *vesselTrailNode) Dispose() {
	t.sprite.Dispose()
}

func (t *vesselTrailNode) Update(delta float64) {
	t.sprite.SetAlpha(t.sprite.GetAlpha() - float32(delta))
	if t.sprite.GetAlpha() < 0.1 {
		t.Dispose()
	}
}
