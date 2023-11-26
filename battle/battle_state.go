package battle

import (
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/shmup-game/viewport"
)

type battleState struct {
	stage *viewport.Stage
	rect  gmath.Rect
	human *humanPlayer
}
