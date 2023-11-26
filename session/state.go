package session

import (
	"github.com/quasilyte/ge/input"
	"github.com/quasilyte/gsignal"
	"github.com/quasilyte/xm"
)

type State struct {
	Input *input.Handler

	EventPlayerUpdate gsignal.Event[xm.StreamEvent]
}
