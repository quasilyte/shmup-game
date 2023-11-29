package battle

import (
	"time"
)

type Result struct {
	Victory bool

	DodgePoints     float64
	PressurePenalty float64
	TimePlayed      time.Duration
}
