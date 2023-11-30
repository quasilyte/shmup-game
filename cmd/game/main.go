package main

import (
	"fmt"
	"time"

	"github.com/quasilyte/ge"
	"github.com/quasilyte/ge/input"
	"github.com/quasilyte/shmup-game/assets"
	"github.com/quasilyte/shmup-game/controls"
	"github.com/quasilyte/shmup-game/eui"
	"github.com/quasilyte/shmup-game/scenes"
	"github.com/quasilyte/shmup-game/session"
	"github.com/quasilyte/xm"
)

func main() {
	ctx := ge.NewContext(ge.ContextConfig{
		FixedDelta: true,
	})
	ctx.Rand.SetSeed(time.Now().Unix())
	ctx.GameName = "tunefire"
	ctx.WindowTitle = "TuneFire"
	ctx.WindowWidth = 1920 / 2
	ctx.WindowHeight = 1080 / 2
	ctx.FullScreen = true

	state := &session.State{
		Settings: getDefaultSettings(),
	}

	ctx.Loader.OpenAssetFunc = assets.MakeOpenAssetFunc(ctx)
	assets.RegisterResources(ctx, assets.Config{
		PlayerListener: func(e xm.StreamEvent) {
			state.EventPlayerUpdate.Emit(e)
		},
	})

	state.UIResources = eui.PrepareResources(ctx.Loader)

	if err := ctx.LoadGameData("save", &state.Settings); err != nil {
		fmt.Printf("can't load game data: %v", err)
		state.Settings = getDefaultSettings()
		ctx.SaveGameData("save", state.Settings)
	}

	keymap := input.Keymap{
		controls.ActionSpecial:     {input.KeyDown, input.KeyS, input.KeySpace},
		controls.ActionMoveTurbo:   {input.KeyUp, input.KeyW, input.KeyGamepadUp},
		controls.ActionRotateLeft:  {input.KeyLeft, input.KeyA, input.KeyGamepadLeft},
		controls.ActionRotateRight: {input.KeyRight, input.KeyD, input.KeyGamepadRight},
		controls.ActionStrafe:      {input.KeySpace, input.KeyShift},

		controls.ActionMenuBack: {input.KeyEscape},
	}
	state.Input = ctx.Input.NewHandler(0, keymap)

	if err := ge.RunGame(ctx, scenes.NewMainMenuController(state)); err != nil {
		panic(err)
	}
}

func getDefaultSettings() session.Settings {
	return session.Settings{
		SoundLevel:      3,
		MusicLevel:      3,
		Difficulty:      0, // For game jams, easy is a good default
		LevelsAvailable: 1,
	}
}
