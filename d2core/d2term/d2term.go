package d2term

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
)

// Create creates and initializes the terminal
func Create() (d2interface.Terminal, error) {
	term, err := createTerminal()
	if err != nil {
		return nil, err
	}

	if err := d2input.BindHandlerWithPriority(term, d2interface.PriorityHigh); err != nil {
		return nil, err
	}

	return term, nil
}
