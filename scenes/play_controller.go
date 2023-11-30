package scenes

import (
	"github.com/ebitenui/ebitenui/widget"
	"github.com/quasilyte/ge"
	"github.com/quasilyte/shmup-game/assets"
	"github.com/quasilyte/shmup-game/controls"
	"github.com/quasilyte/shmup-game/eui"
	"github.com/quasilyte/shmup-game/session"
)

type PlayController struct {
	state *session.State
	scene *ge.Scene
}

func NewPlayController(state *session.State) *PlayController {
	return &PlayController{state: state}
}

func (c *PlayController) Init(scene *ge.Scene) {
	scene.Audio().ContinueMusic(assets.AudioMusicMenu)

	c.scene = scene

	root := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			StretchHorizontal: true,
		})),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()))

	rowContainer := eui.NewRowLayoutContainerWithMinWidth(320, 8, nil)
	root.AddChild(rowContainer)

	rowContainer.AddChild(eui.NewCenteredLabel("Play", assets.BitmapFont2))

	// rowContainer.AddChild(eui.NewSelectButton(eui.SelectButtonConfig{
	// 	Resources: c.state.UIResources,
	// 	Input:     c.state.Input,
	// 	Value:     &c.state.Settings.SelectedMusic,
	// 	Label:     "Music Track",
	// 	ValueNames: []string{
	// 		gamedata.Music1.Name,
	// 		gamedata.Music2.Name,
	// 		gamedata.Music3.Name,
	// 	},
	// }))

	rowContainer.AddChild(eui.NewSelectButton(eui.SelectButtonConfig{
		Resources: c.state.UIResources,
		Input:     c.state.Input,
		Value:     &c.state.Settings.Weapon,
		Label:     "Weapon",
		ValueNames: []string{
			"Pulse Laser",
			"Rear Cannon",
			"Frag Cannon",
			"Twin Cannon",
		},
	}))

	rowContainer.AddChild(eui.NewSelectButton(eui.SelectButtonConfig{
		Resources: c.state.UIResources,
		Input:     c.state.Input,
		Value:     &c.state.Settings.Vessel,
		Label:     "Vessel",
		ValueNames: []string{
			"Interceptor",
			"Vindicator",
		},
	}))

	rowContainer.AddChild(eui.NewSelectButton(eui.SelectButtonConfig{
		Resources: c.state.UIResources,
		Input:     c.state.Input,
		Value:     &c.state.Settings.Special,
		Label:     "Special",
		ValueNames: []string{
			"Dash",
			"Mega Bomb",
			"Spin Shield",
		},
	}))

	rowContainer.AddChild(eui.NewSelectButton(eui.SelectButtonConfig{
		Resources: c.state.UIResources,
		Input:     c.state.Input,
		Value:     &c.state.Settings.Difficulty,
		Label:     "Difficulty",
		ValueNames: []string{
			"Easy",
			"Normal",
			"Hard",
			"Nightmare",
		},
	}))

	panel := eui.NewPanelWithPadding(c.state.UIResources, 300, 0, widget.NewInsetsSimple(12))
	rowContainer.AddChild(panel)

	grid := eui.NewGridContainer(3, widget.GridLayoutOpts.Spacing(24, 4),
		widget.GridLayoutOpts.Stretch(nil, nil))

	panel1rows := eui.NewRowLayoutContainerWithMinWidth(320, 8, nil)
	panel.AddChild(panel1rows)
	panel1rows.AddChild(eui.NewCenteredLabel("Select Level", assets.BitmapFont1))
	panel1rows.AddChild(grid)

	levelLabels := []string{"I", "II", "III", "IV", "V", "VI"}
	for i := range levelLabels {
		levelID := i
		b := eui.NewButtonWithConfig(c.state.UIResources, eui.ButtonConfig{
			MinWidth: 120,
			Text:     levelLabels[i],
			OnClick: func() {
				c.scene.Context().SaveGameData("save", c.state.Settings)
				c.scene.Context().ChangeScene(NewBattleController(c.state, levelID))
			},
		})
		grid.AddChild(b)
	}

	rowContainer.AddChild(eui.NewButton(c.state.UIResources, "BACK", func() {
		c.back()
	}))

	initUI(scene, root)
}

func (c *PlayController) Update(delta float64) {
	if c.state.Input.ActionIsJustReleased(controls.ActionMenuBack) {
		c.back()
	}
}

func (c *PlayController) back() {
	c.scene.Context().SaveGameData("save", c.state.Settings)
	c.scene.Context().ChangeScene(NewMainMenuController(c.state))
}
