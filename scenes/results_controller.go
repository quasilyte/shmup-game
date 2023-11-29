package scenes

import (
	"strconv"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/shmup-game/assets"
	"github.com/quasilyte/shmup-game/battle"
	"github.com/quasilyte/shmup-game/eui"
	"github.com/quasilyte/shmup-game/session"
)

type ResultsController struct {
	scene  *ge.Scene
	state  *session.State
	result battle.Result
}

func NewResultsController(state *session.State, result battle.Result) *ResultsController {
	return &ResultsController{
		state:  state,
		result: result,
	}
}

func (c *ResultsController) Init(scene *ge.Scene) {
	c.scene = scene

	root := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			StretchHorizontal: true,
		})),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()))

	rowContainer := eui.NewRowLayoutContainerWithMinWidth(320, 8, nil)
	root.AddChild(rowContainer)

	title := "Defeat"
	if c.result.Victory {
		title = "Victory"
	}
	rowContainer.AddChild(eui.NewCenteredLabel(title, assets.BitmapFont2))

	panel := eui.NewPanelWithPadding(c.state.UIResources, 300, 0, widget.NewInsetsSimple(24))
	rowContainer.AddChild(panel)

	grid := eui.NewGridContainer(2, widget.GridLayoutOpts.Spacing(24, 4),
		widget.GridLayoutOpts.Stretch([]bool{true, false}, nil))
	panel.AddChild(grid)

	score := c.calcScore()

	lines := [][2]string{
		{"Time played", formatDurationCompact(c.result.TimePlayed)},
		{"Distance penalty", strconv.Itoa(int(c.result.PressurePenalty))},
		{"Perfect dodges", strconv.Itoa(int(c.result.DodgePoints))},
		{"Score", strconv.Itoa(score)},
		{"Rank", c.calcRank(score)},
	}
	for _, pair := range lines {
		label := pair[0]
		value := pair[1]
		grid.AddChild(eui.NewLabel(label, assets.BitmapFont1))
		grid.AddChild(eui.NewLabel(value, assets.BitmapFont1))
	}

	rowContainer.AddChild(eui.NewButton(c.state.UIResources, "OK", func() {
		scene.Context().ChangeScene(NewMainMenuController(c.state))
	}))

	initUI(scene, root)
}

func (c *ResultsController) Update(delta float64) {}

func (c *ResultsController) calcRank(score int) string {
	if !c.result.Victory {
		return "none"
	}

	switch {
	case score >= 2000:
		return "S+"
	case score >= 1500:
		return "S"
	case score >= 1200:
		return "A"
	case score >= 1000:
		return "B"
	case score >= 600:
		return "C"
	case score >= 300:
		return "D"
	case score >= 100:
		return "E"
	default:
		return "F"
	}
}

func (c *ResultsController) calcScore() int {
	if !c.result.Victory {
		return 0
	}

	score := 1000.0

	badTimeSec := 8.0 * 60.0
	timeSec := c.result.TimePlayed.Seconds()
	timeMultiplier := gmath.ClampMin(1.0-(timeSec*(1/badTimeSec)), 0.0001)

	score += 3.0 * float64(c.result.DodgePoints)
	score -= float64(int(c.result.PressurePenalty))

	return gmath.ClampMin(int(score*timeMultiplier), 1)
}
