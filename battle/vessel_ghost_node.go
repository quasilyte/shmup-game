package battle

import (
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
)

type vesselGhostNode struct {
	world    *battleState
	sprite   *ge.Sprite
	img      resource.ImageID
	pos      gmath.Vec
	rotation gmath.Rad
}

func newVesselGhostNode(world *battleState, img resource.ImageID, pos gmath.Vec, rotation gmath.Rad) *vesselGhostNode {
	return &vesselGhostNode{
		world:    world,
		img:      img,
		pos:      pos,
		rotation: rotation,
	}
}

func (g *vesselGhostNode) Init(scene *ge.Scene) {
	g.sprite = scene.NewSprite(g.img)
	g.sprite.SetColorScale(ge.ColorScale{R: 0.45, G: 0.6, B: 0.6, A: 0.9})
	g.sprite.Pos.Base = &g.pos
	g.sprite.Rotation = &g.rotation
	g.world.stage.AddSprite(g.sprite)
}

func (g *vesselGhostNode) IsDisposed() bool {
	return g.sprite.IsDisposed()
}

func (g *vesselGhostNode) Dispose() {
	g.sprite.Dispose()
}

func (g *vesselGhostNode) Update(delta float64) {
	g.sprite.SetAlpha(g.sprite.GetAlpha() - float32(2*delta))
	if g.sprite.GetAlpha() < 0.05 {
		g.Dispose()
	}
}
