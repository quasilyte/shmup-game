package viewport

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/ge"
)

type UserInterfaceLayer struct {
	belowObjects []ge.SceneGraphics
	objects      []ge.SceneGraphics
	aboveObjects []ge.SceneGraphics

	Visible bool
}

func (l *UserInterfaceLayer) AddGraphicsBelow(o ge.SceneGraphics) {
	l.belowObjects = append(l.belowObjects, o)
}

func (l *UserInterfaceLayer) AddGraphics(o ge.SceneGraphics) {
	l.objects = append(l.objects, o)
}

func (l *UserInterfaceLayer) AddGraphicsAbove(o ge.SceneGraphics) {
	l.aboveObjects = append(l.aboveObjects, o)
}

func drawSlice(dst *ebiten.Image, slice []ge.SceneGraphics) []ge.SceneGraphics {
	live := slice[:0]
	for _, o := range slice {
		if o.IsDisposed() {
			continue
		}
		o.Draw(dst)
		live = append(live, o)
	}
	return live
}
