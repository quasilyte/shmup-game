package controls

import (
	"github.com/quasilyte/ge/input"
)

const (
	ActionUnknown input.Action = iota

	ActionMoveTurbo
	ActionRotateLeft
	ActionRotateRight
	ActionStrafe
)
