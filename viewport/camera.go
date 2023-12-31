package viewport

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
)

type Camera struct {
	Stage *Stage

	Offset   gmath.Vec
	Rotation gmath.Rad

	// A layer that is always on top of everything else.
	// It's also position-independent.
	UI *UserInterfaceLayer

	width     float64
	height    float64
	WorldRect gmath.Rect

	Rect       gmath.Rect
	globalRect gmath.Rect

	screen    *ebiten.Image
	tmpScreen *ebiten.Image
}

func NewCamera(stage *Stage, world gmath.Rect, width, height float64) *Camera {
	cam := &Camera{
		Stage: stage,

		WorldRect: world,
		width:     world.Width(),
		height:    world.Height(),

		Rect: gmath.Rect{
			Min: gmath.Vec{},
			Max: gmath.Vec{X: width, Y: height},
		},

		UI: &UserInterfaceLayer{
			belowObjects: make([]ge.SceneGraphics, 0, 4),
			objects:      make([]ge.SceneGraphics, 0, 4),
			aboveObjects: make([]ge.SceneGraphics, 0, 4),
			Visible:      true,
		},

		screen:    ebiten.NewImage(int(width), int(height)),
		tmpScreen: ebiten.NewImage(int(width), int(height)),
	}

	return cam
}

func (c *Camera) IsDisposed() bool { return false }

func (c *Camera) Draw(screen *ebiten.Image) {
	c.globalRect = c.Rect
	c.globalRect.Min = c.Offset
	c.globalRect.Max = c.globalRect.Max.Add(c.Offset)

	c.screen.Clear()
	drawOffset := gmath.Vec{
		X: -c.Offset.X,
		Y: -c.Offset.Y,
	}
	c.drawLayer(c.screen, &c.Stage.belowObjects, drawOffset)
	c.drawLayer(c.screen, &c.Stage.objects, drawOffset)
	c.drawLayer(c.screen, &c.Stage.slightlyAboveObjects, drawOffset)
	c.drawLayer(c.screen, &c.Stage.aboveObjects, drawOffset)

	var options ebiten.DrawImageOptions
	options.GeoM.Translate(-c.Rect.Width()*0.5, -c.Rect.Height()*0.5)
	options.GeoM.Rotate(float64(c.Rotation.Normalized()))
	options.GeoM.Translate(c.Rect.Width()*0.5, c.Rect.Height()*0.5)

	if c.Stage.Shader != nil {
		width := c.screen.Bounds().Dx()
		height := c.screen.Bounds().Dy()
		c.tmpScreen.Clear()
		c.tmpScreen.DrawImage(c.screen, &options)
		var options2 ebiten.DrawRectShaderOptions
		options2.Images[0] = c.tmpScreen
		options2.Uniforms = c.Stage.ShaderParams
		screen.DrawRectShader(width, height, c.Stage.Shader, &options2)
	} else {
		screen.DrawImage(c.screen, &options)
	}

	// if c.Stage.Shader == nil {
	// 	screen.DrawImage(c.screen, &options)
	// } else {
	// 	width := c.screen.Bounds().Dx()
	// 	height := c.screen.Bounds().Dy()
	// 	var options2 ebiten.DrawRectShaderOptions
	// 	options2.GeoM = options.GeoM
	// 	options2.Images[0] = c.screen
	// 	options2.Uniforms = c.Stage.ShaderParams
	// 	screen.DrawRectShader(width, height, c.Stage.Shader, &options2)
	// }

	if c.UI.Visible {
		c.UI.belowObjects = drawSlice(screen, c.UI.belowObjects)
		c.UI.objects = drawSlice(screen, c.UI.objects)
		c.UI.aboveObjects = drawSlice(screen, c.UI.aboveObjects)
	}
}

func (c *Camera) AbsPos(screenPos gmath.Vec) gmath.Vec {
	return screenPos.Add(c.Offset)
}

func (c *Camera) ContainsPos(pos gmath.Vec) bool {
	globalRect := c.Rect
	globalRect.Min = c.Offset
	globalRect.Max = globalRect.Max.Add(c.Offset)
	return globalRect.Contains(pos)
}

func (c *Camera) checkBounds() {
	c.Offset.X = gmath.Clamp(c.Offset.X, 0, c.width-c.Rect.Width())
	c.Offset.Y = gmath.Clamp(c.Offset.Y, 0, c.height-c.Rect.Height())
}

func (c *Camera) Pan(delta gmath.Vec) {
	if delta.IsZero() {
		return
	}
	c.Offset = c.Offset.Add(delta)
	c.checkBounds()
}

func (c Camera) CenterPos() gmath.Vec {
	return c.Offset.Add(c.Rect.Center())
}

func (c *Camera) CenterOn(pos gmath.Vec) {
	c.Offset = pos.Sub(c.Rect.Center())
	c.checkBounds()
}

func (c *Camera) SetOffset(pos gmath.Vec) {
	c.Offset = pos
	c.checkBounds()
}

func (c *Camera) drawLayer(screen *ebiten.Image, l *layer, drawOffset gmath.Vec) {
	for _, s := range l.sprites {
		if c.isVisible(s.BoundsRect()) {
			s.DrawWithOffset(screen, drawOffset)
		}
	}

	if len(l.objects) != 0 {
		for _, o := range l.objects {
			if c.isVisible(o.BoundsRect()) {
				o.DrawWithOffset(screen, drawOffset)
			}
		}
	}
}

func (c *Camera) isVisible(objectRect gmath.Rect) bool {
	cameraRect := c.globalRect

	if objectRect.Max.X < cameraRect.Min.X {
		return false
	}
	if objectRect.Min.X > cameraRect.Max.X {
		return false
	}
	if objectRect.Max.Y < cameraRect.Min.Y {
		return false
	}
	if objectRect.Min.Y > cameraRect.Max.Y {
		return false
	}

	return true
}
