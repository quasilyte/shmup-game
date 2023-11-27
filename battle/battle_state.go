package battle

import (
	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/shmup-game/viewport"
)

type battleState struct {
	scene *ge.Scene
	stage *viewport.Stage
	rect  gmath.Rect
	human *humanPlayer
	bot   botPlayer
}

type botPlayer interface{}
