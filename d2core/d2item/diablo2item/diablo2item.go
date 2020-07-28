package diablo2item

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"

// NewProperty creates a property
func NewProperty(code string, values ...int) *Property {
	record := d2datadict.Properties[code]

	if record == nil {
		return nil
	}

	result := &Property{
		record:      record,
		inputParams: values,
	}

	return result.init()
}
