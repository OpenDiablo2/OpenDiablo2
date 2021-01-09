package d2term

import (
	"sort"
)

func (t *Terminal) commandList([]string) error {
	names := make([]string, 0, len(t.commands))
	for name := range t.commands {
		names = append(names, name)
	}

	sort.Strings(names)
	t.Infof("available actions (%d):", len(names))

	for _, name := range names {
		entry := t.commands[name]
		if entry.arguments != nil {
			t.Infof("%s: %s; %v", name, entry.description, entry.arguments)
			continue
		}

		t.Infof("%s: %s", name, entry.description)
	}

	return nil
}

func (t *Terminal) commandClear([]string) error {
	t.Clear()

	return nil
}
