package controls

import (
	"github.com/quasilyte/ge/input"
)

const (
	ActionUnknown input.Action = iota

	ActionMoveAccelerate
	ActionMoveDecelerate
	ActionRotateLeft
	ActionRotateRight
	ActionStrafe
)
