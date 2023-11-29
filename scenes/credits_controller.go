package scenes

import (
	"strings"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/quasilyte/ge"
	"github.com/quasilyte/shmup-game/assets"
	"github.com/quasilyte/shmup-game/controls"
	"github.com/quasilyte/shmup-game/eui"
	"github.com/quasilyte/shmup-game/session"
)

type CreditsController struct {
	state *session.State
	scene *ge.Scene
}

func NewCreditsController(state *session.State) *CreditsController {
	return &CreditsController{state: state}
}

func (c *CreditsController) Init(scene *ge.Scene) {
	c.scene = scene

	root := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			StretchHorizontal: true,
		})),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()))

	rowContainer := eui.NewRowLayoutContainerWithMinWidth(320, 8, nil)
	root.AddChild(rowContainer)

	rowContainer.AddChild(eui.NewCenteredLabel("Credits", assets.BitmapFont2))

	panel := eui.NewPanelWithPadding(c.state.UIResources, 300, 0, widget.NewInsetsSimple(24))
	rowContainer.AddChild(panel)

	text := strings.Join([]string{
		"A game by quasilyte.",
		"Made for a Game Off 2023 game jam in ~4 days.",
		"",
		"Special thanks to Drozerix, the music author.",
		"",
		"This game is made with Ebitengine,",
		"a game engine for Go programming language.",
	}, "\n")

	panel.AddChild(eui.NewLabel(text, assets.BitmapFont1))

	rowContainer.AddChild(eui.NewButton(c.state.UIResources, "BACK", func() {
		c.back()
	}))

	initUI(scene, root)
}

func (c *CreditsController) Update(delta float64) {
	if c.state.Input.ActionIsJustReleased(controls.ActionMenuBack) {
		c.back()
	}
}

func (c *CreditsController) back() {
	c.scene.Context().ChangeScene(NewMainMenuController(c.state))
}
