package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/shmup-game/assets"
	"github.com/quasilyte/shmup-game/battle"
	"github.com/quasilyte/shmup-game/gamedata"
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
	scene.Audio().PauseCurrentMusic()

	c.state.EventPlayerUpdate.Reset()

	c.scene = scene

	sectorSize := gmath.Vec{X: 960, Y: 640}
	textures := make([]*ebiten.Image, 16)

	for i := range textures {
		bg := ge.NewTiledBackground(scene.Context())
		bg.LoadTilesetWithRand(scene.Context(), scene.Rand(), sectorSize.X, sectorSize.Y, assets.ImageTileset1, assets.RawTiles1)
		img := ebiten.NewImage(int(sectorSize.X), int(sectorSize.Y))
		bg.Draw(img)
		textures[i] = img
	}

	stage := viewport.NewStage()

	shader := scene.Context().Loader.LoadShader(assets.ShaderCRT).Data
	stage.Shader = shader

	c.runner = battle.NewRunner(battle.RunnerConfig{
		Session:        c.state,
		Stage:          stage,
		Music:          gamedata.Music1,
		SectorSize:     sectorSize,
		SectorTextures: textures,
	})
	c.runner.Init(scene)

	scene.Audio().PlayMusic(assets.AudioMusic1)
}

func (c *BattleController) Update(delta float64) {
	c.runner.Update(delta)
}
