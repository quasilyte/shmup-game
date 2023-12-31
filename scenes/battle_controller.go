package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/shmup-game/assets"
	"github.com/quasilyte/shmup-game/battle"
	"github.com/quasilyte/shmup-game/controls"
	"github.com/quasilyte/shmup-game/gamedata"
	"github.com/quasilyte/shmup-game/session"
	"github.com/quasilyte/shmup-game/viewport"
)

type BattleController struct {
	state   *session.State
	levelID int
	scene   *ge.Scene
	runner  *battle.Runner
}

func NewBattleController(state *session.State, levelID int) *BattleController {
	return &BattleController{
		state:   state,
		levelID: levelID,
	}
}

func (c *BattleController) Init(scene *ge.Scene) {
	levelMusic := []int{
		0, 1, 2,
		0, 1, 2,
	}

	c.state.EventPlayerUpdate.Reset()
	scene.Audio().PauseCurrentMusic()

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

	music := gamedata.MusicList[levelMusic[c.levelID]]

	c.runner = battle.NewRunner(battle.RunnerConfig{
		LevelID:        c.levelID,
		Session:        c.state,
		Stage:          stage,
		Music:          music,
		SectorSize:     sectorSize,
		SectorTextures: textures,
	})
	c.runner.Init(scene)

	c.runner.EventBattleOver.Connect(nil, func(result battle.Result) {
		c.leaveScene(NewResultsController(c.state, result))
	})

	scene.Audio().PlayMusic(music.AudioID)
}

func (c *BattleController) leaveScene(next ge.SceneController) {
	c.scene.Audio().PauseCurrentMusic()
	c.state.EventPlayerUpdate.Reset()
	c.scene.Context().ChangeScene(next)
}

func (c *BattleController) Update(delta float64) {
	if c.state.Input.ActionIsJustReleased(controls.ActionMenuBack) {
		c.leaveScene(NewPlayController(c.state))
		return
	}
	c.runner.Update(delta)
}
