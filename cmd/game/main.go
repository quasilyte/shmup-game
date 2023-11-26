package main

import (
	"time"

	"github.com/quasilyte/ge"
	"github.com/quasilyte/ge/input"
	"github.com/quasilyte/shmup-game/assets"
	"github.com/quasilyte/shmup-game/controls"
	"github.com/quasilyte/shmup-game/scenes"
	"github.com/quasilyte/shmup-game/session"
	"github.com/quasilyte/xm"
)

func main() {
	ctx := ge.NewContext(ge.ContextConfig{
		FixedDelta: true,
	})
	ctx.Rand.SetSeed(time.Now().Unix())
	ctx.GameName = "dogfight"
	ctx.WindowTitle = "Dogfight"
	ctx.WindowWidth = 1920 / 2
	ctx.WindowHeight = 1080 / 2
	ctx.FullScreen = true

	state := &session.State{}

	ctx.Loader.OpenAssetFunc = assets.MakeOpenAssetFunc(ctx)
	assets.RegisterResources(ctx, assets.Config{
		PlayerListener: func(e xm.StreamEvent) {
			state.EventPlayerUpdate.Emit(e)
		},
	})

	keymap := input.Keymap{
		controls.ActionMoveTurbo:   {input.KeyUp, input.KeyW, input.KeyGamepadUp},
		controls.ActionRotateLeft:  {input.KeyLeft, input.KeyA, input.KeyGamepadLeft},
		controls.ActionRotateRight: {input.KeyRight, input.KeyD, input.KeyGamepadRight},
		controls.ActionStrafe:      {input.KeySpace, input.KeyShift},
	}
	state.Input = ctx.Input.NewHandler(0, keymap)

	if err := ge.RunGame(ctx, scenes.NewBattleController(state)); err != nil {
		panic(err)
	}
}
