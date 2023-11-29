package battle

import (
	"time"
)

type Result struct {
	Victory bool

	Difficulty      int
	Health          float64
	DodgePoints     float64
	PressurePenalty float64
	TimePlayed      time.Duration
}
