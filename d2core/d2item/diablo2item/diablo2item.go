package diablo2item

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"

func NewItem(codes ...string) *Item {
	var item *Item

	var common, set, unique string

	var prefixes, suffixes []string

	for _, code := range codes {
		if found := d2datadict.CommonItems[code]; found != nil {
			common = code
			continue
		}

		if found := d2datadict.SetItems[code]; found != nil {
			set = code
			continue
		}

		if found := d2datadict.UniqueItems[code]; found != nil {
			unique = code
			continue
		}

		if found := d2datadict.MagicPrefix[code]; found != nil {
			if prefixes == nil {
				prefixes = make([]string, 0)
			}

			prefixes = append(prefixes, code)

			continue
		}

		if found := d2datadict.MagicSuffix[code]; found != nil {
			if suffixes == nil {
				suffixes = make([]string, 0)
			}

			suffixes = append(suffixes, code)

			continue
		}
	}

	if common != "" { // we will at least have a regular item
		item = &Item{CommonCode: common}

		if set != "" { // it's a set item
			item.SetItemCode = set
			return item.init()
		}

		if unique != "" { // it's a unique item
			item.UniqueCode = unique
			return item.init()
		}

		if prefixes != nil {
			if len(prefixes) > 0 { // it's a magic or rare item
				item.PrefixCodes = prefixes
			}
		}

		if suffixes != nil {
			if len(suffixes) > 0 { // it's a magic or rare item
				item.SuffixCodes = suffixes
			}
		}

		return item.init()
	}

	return nil
}

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
