package d2term

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
)

// Initialize creates and initializes the terminal
func Initialize() (d2interface.Terminal, error) {
	term, err := createTerminal()
	if err != nil {
		return nil, err
	}

	if err := d2input.BindHandlerWithPriority(term, d2input.PriorityHigh); err != nil {
		return nil, err
	}

	return term, nil
}
