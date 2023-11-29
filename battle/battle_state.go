package battle

import (
	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/shmup-game/viewport"
)

type battleState struct {
	scene  *ge.Scene
	stage  *viewport.Stage
	rect   gmath.Rect
	human  *humanPlayer
	bot    botPlayer
	result *Result

	playerDamageMultiplier float64
	botDamageMultiplier    float64
}

type botPlayer interface {
	GetVessel() *vesselNode
}
