package session

import (
	"github.com/quasilyte/ge/input"
	"github.com/quasilyte/gsignal"
	"github.com/quasilyte/shmup-game/eui"
	"github.com/quasilyte/xm"
)

type State struct {
	UIResources *eui.Resources

	Input *input.Handler

	Settings Settings

	EventPlayerUpdate gsignal.Event[xm.StreamEvent]
}

type Settings struct {
	Weapon          int
	Vessel          int
	Special         int
	Difficulty      int
	SoundLevel      int
	MusicLevel      int
	LevelsAvailable int
	Hints           int
}
