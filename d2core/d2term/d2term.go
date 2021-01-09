package d2term

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// New creates and initializes the terminal
func New(inputManager d2interface.InputManager) (d2interface.Terminal, error) {
	term, err := createTerminal()
	if err != nil {
		return nil, err
	}

	if err := inputManager.BindHandlerWithPriority(term, d2enum.PriorityHigh); err != nil {
		return nil, err
	}

	return term, nil
}
