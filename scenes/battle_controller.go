package scenes

import (
	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/shmup-game/assets"
	"github.com/quasilyte/shmup-game/battle"
	"github.com/quasilyte/shmup-game/session"
	"github.com/quasilyte/shmup-game/viewport"
)

type BattleController struct {
	state  *session.State
	scene  *ge.Scene
	runner *battle.Runner
}

func NewBattleController(state *session.State) *BattleController {
	return &BattleController{
		state: state,
	}
}

func (c *BattleController) Init(scene *ge.Scene) {
	c.state.EventPlayerUpdate.Reset()

	scene.Audio().PlayMusic(assets.AudioMusic1)

	c.scene = scene

	worldRect := gmath.Rect{
		Max: gmath.Vec{X: 1024, Y: 1024 * 4},
	}

	bg := ge.NewTiledBackground(scene.Context())
	bg.LoadTilesetWithRand(scene.Context(), scene.Rand(), worldRect.Width(), worldRect.Height(), assets.ImageTileset1, assets.RawTiles1)

	stage := viewport.NewStage()
	stage.SetBackground(bg)

	c.runner = battle.NewRunner(battle.RunnerConfig{
		Session:   c.state,
		Stage:     stage,
		WorldRect: worldRect,
	})
	c.runner.Init(scene)
}

func (c *BattleController) Update(delta float64) {
	c.runner.Update(delta)
}
