package viewport

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/ge"
)

type Stage struct {
	bg *ge.TiledBackground

	LayerContainer

	Shader       *ebiten.Shader
	ShaderParams map[string]any
}

func NewStage() *Stage {
	return &Stage{}
}

func (stage *Stage) SetBackground(bg *ge.TiledBackground) {
	stage.bg = bg
}

func (stage *Stage) Update() {
	stage.belowObjects.filter()
	stage.objects.filter()
	stage.slightlyAboveObjects.filter()
	stage.aboveObjects.filter()
}
