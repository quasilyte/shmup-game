package controls

import (
	"github.com/quasilyte/ge/input"
)

const (
	ActionUnknown input.Action = iota

	ActionSpecial
	ActionMoveTurbo
	ActionRotateLeft
	ActionRotateRight
	ActionStrafe

	ActionMenuBack
)
