package scenes

import (
	"fmt"
	"time"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/quasilyte/ge"
	"github.com/quasilyte/shmup-game/assets"
	"github.com/quasilyte/shmup-game/eui"
)

func initUI(scene *ge.Scene, root *widget.Container) {
	bg := scene.NewSprite(assets.ImageMenuBg)
	bg.Centered = false
	scene.AddGraphics(bg)

	uiObject := eui.NewSceneObject(root)
	scene.AddGraphics(uiObject)
	scene.AddObject(uiObject)
}

func formatDurationCompact(d time.Duration) string {
	d = d.Round(time.Second)
	hours := d / time.Hour
	d -= hours * time.Hour
	minutes := d / time.Minute
	d -= minutes * time.Minute
	seconds := d / time.Second
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
