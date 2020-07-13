package d2input

import (
	"errors"
)

var (
	// ErrHasReg shows the input system already has a registered handler
	ErrHasReg = errors.New("input system already has provided handler")
	// ErrNotReg shows the input system has no registered handler
	ErrNotReg = errors.New("input system does not have provided handler")
)
