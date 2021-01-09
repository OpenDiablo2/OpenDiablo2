package d2term

import (
	"errors"
	"fmt"
	"image/color"
	"log"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
)

const (
	termCharWidth       = 6
	termCharHeight      = 16
	termCharDoubleWidth = termCharWidth * 2
	termRowCount        = 24
	termRowCountMax     = 32
	termColCountMax     = 128
	termAnimLength      = 0.5
)

const (
	darkGrey  = 0x2e3436b0
	lightGrey = 0x555753b0
	lightBlue = 0x3465a4b0
	yellow    = 0xfce94fb0
	red       = 0xcc0000b0
)

type termVis int

const (
	termVisHidden termVis = iota
	termVisShowing
	termVisShown
	termVisHiding
)

const (
	maxVisAnim = 1.0
	minVisAnim = 0.0
)

type termHistoryEntry struct {
	text     string
	category d2enum.TermCategory
}

type termActionEntry struct {
	action      interface{}
	description string
}

type terminal struct {
	outputHistory []termHistoryEntry
	outputIndex   int

	command        string
	commandHistory []string
	commandIndex   int

	lineCount int
	visState  termVis
	visAnim   float64

	bgColor      color.RGBA
	fgColor      color.RGBA
	infoColor    color.RGBA
	warningColor color.RGBA
	errorColor   color.RGBA

	actions map[string]termActionEntry
}

func (t *terminal) Advance(elapsed float64) error {
	switch t.visState {
	case termVisShowing:
		t.visAnim = math.Min(maxVisAnim, t.visAnim+elapsed/termAnimLength)
		if t.visAnim == maxVisAnim {
			t.visState = termVisShown
		}
	case termVisHiding:
		t.visAnim = math.Max(minVisAnim, t.visAnim-elapsed/termAnimLength)
		if t.visAnim == minVisAnim {
			t.visState = termVisHidden
		}
	}

	if !t.IsVisible() {
		return nil
	}

	return nil
}

func (t *terminal) OnKeyDown(event d2interface.KeyEvent) bool {
	if event.Key() == d2enum.KeyGraveAccent {
		t.toggleTerminal()
	}

	if !t.IsVisible() {
		return false
	}

	switch event.Key() {
	case d2enum.KeyEscape:
		t.command = ""
	case d2enum.KeyEnd:
		t.outputIndex = 0
	case d2enum.KeyHome:
		t.outputIndex = d2math.MaxInt(0, len(t.outputHistory)-t.lineCount)
	case d2enum.KeyPageUp:
		maxOutputIndex := d2math.MaxInt(0, len(t.outputHistory)-t.lineCount)
		if t.outputIndex += t.lineCount; t.outputIndex >= maxOutputIndex {
			t.outputIndex = maxOutputIndex
		}
	case d2enum.KeyPageDown:
		if t.outputIndex -= t.lineCount; t.outputIndex < 0 {
			t.outputIndex = 0
		}
	case d2enum.KeyUp, d2enum.KeyDown:
		t.handleControlKey(event.Key(), event.KeyMod())
	case d2enum.KeyEnter:
		t.processCommand()
	case d2enum.KeyBackspace:
		if len(t.command) > 0 {
			t.command = t.command[:len(t.command)-1]
		}
	}

	return true
}

func (t *terminal) processCommand() {
	if t.command == "" {
		return
	}

	n := 0

	for _, command := range t.commandHistory {
		if command != t.command {
			t.commandHistory[n] = command
			n++
		}
	}

	t.commandHistory = t.commandHistory[:n]
	t.commandHistory = append(t.commandHistory, t.command)

	t.Outputf(t.command)

	if err := t.Execute(t.command); err != nil {
		t.OutputErrorf(err.Error())
	}

	t.commandIndex = len(t.commandHistory) - 1
	t.command = ""
}

func (t *terminal) handleControlKey(eventKey d2enum.Key, keyMod d2enum.KeyMod) {
	switch eventKey {
	case d2enum.KeyUp:
		if keyMod == d2enum.KeyModControl {
			t.lineCount = d2math.MaxInt(0, t.lineCount-1)
		} else if len(t.commandHistory) > 0 {
			t.command = t.commandHistory[t.commandIndex]
			if t.commandIndex == 0 {
				t.commandIndex = len(t.commandHistory) - 1
			} else {
				t.commandIndex--
			}
		}
	case d2enum.KeyDown:
		if keyMod == d2enum.KeyModControl {
			t.lineCount = d2math.MinInt(t.lineCount+1, termRowCountMax)
		}
	}
}

func (t *terminal) toggleTerminal() {
	if t.visState == termVisHiding || t.visState == termVisHidden {
		t.Show()
	} else {
		t.Hide()
	}
}

func (t *terminal) OnKeyChars(event d2interface.KeyCharsEvent) bool {
	if !t.IsVisible() {
		return false
	}

	var handled bool

	for _, c := range event.Chars() {
		if c != '`' {
			t.command += string(c)
			handled = true
		}
	}

	return handled
}

func (t *terminal) Render(surface d2interface.Surface) error {
	if !t.IsVisible() {
		return nil
	}

	totalWidth, _ := surface.GetSize()
	outputHeight := t.lineCount * termCharHeight
	totalHeight := outputHeight + termCharHeight

	offset := -int((1.0 - easeInOut(t.visAnim)) * float64(totalHeight))
	surface.PushTranslation(0, offset)

	surface.DrawRect(totalWidth, outputHeight, t.bgColor)

	for i := 0; i < t.lineCount; i++ {
		historyIndex := len(t.outputHistory) - i - t.outputIndex - 1
		if historyIndex < 0 {
			break
		}

		historyEntry := t.outputHistory[historyIndex]

		surface.PushTranslation(termCharDoubleWidth, outputHeight-(i+1)*termCharHeight)
		surface.DrawTextf(historyEntry.text)
		surface.PushTranslation(-termCharDoubleWidth, 0)

		switch historyEntry.category {
		case d2enum.TermCategoryInfo:
			surface.DrawRect(termCharWidth, termCharHeight, t.infoColor)
		case d2enum.TermCategoryWarning:
			surface.DrawRect(termCharWidth, termCharHeight, t.warningColor)
		case d2enum.TermCategoryError:
			surface.DrawRect(termCharWidth, termCharHeight, t.errorColor)
		}

		surface.Pop()
		surface.Pop()
	}

	surface.PushTranslation(0, outputHeight)
	surface.DrawRect(totalWidth, termCharHeight, t.fgColor)
	surface.DrawTextf("> " + t.command)
	surface.Pop()

	surface.Pop()

	return nil
}

func (t *terminal) Execute(command string) error {
	params := parseCommand(command)
	if len(params) == 0 {
		return errors.New("invalid command")
	}

	actionName := params[0]
	actionParams := params[1:]

	actionEntry, ok := t.actions[actionName]
	if !ok {
		return errors.New("action not found")
	}

	actionType := reflect.TypeOf(actionEntry.action)
	if actionType.Kind() != reflect.Func {
		return errors.New("action is not a function")
	}

	if len(actionParams) != actionType.NumIn() {
		return errors.New("action requires different argument count")
	}

	paramValues, err := parseActionParams(actionType, actionParams)
	if err != nil {
		return err
	}

	actionValue := reflect.ValueOf(actionEntry.action)
	actionReturnValues := actionValue.Call(paramValues)

	if actionReturnValueCount := len(actionReturnValues); actionReturnValueCount > 0 {
		t.OutputInfof("function returned %d values:", actionReturnValueCount)

		for _, actionReturnValue := range actionReturnValues {
			t.OutputInfof("%v: %s", actionReturnValue.Interface(), actionReturnValue.String())
		}
	}

	return nil
}

func parseActionParams(actionType reflect.Type, actionParams []string) ([]reflect.Value, error) {
	var paramValues []reflect.Value

	for i := 0; i < actionType.NumIn(); i++ {
		actionParam := actionParams[i]

		switch actionType.In(i).Kind() {
		case reflect.String:
			paramValues = append(paramValues, reflect.ValueOf(actionParam))
		case reflect.Int:
			value, err := strconv.ParseInt(actionParam, 10, 64)
			if err != nil {
				return nil, err
			}

			paramValues = append(paramValues, reflect.ValueOf(int(value)))
		case reflect.Uint:
			value, err := strconv.ParseUint(actionParam, 10, 64)
			if err != nil {
				return nil, err
			}

			paramValues = append(paramValues, reflect.ValueOf(uint(value)))
		case reflect.Float64:
			value, err := strconv.ParseFloat(actionParam, 64)
			if err != nil {
				return nil, err
			}

			paramValues = append(paramValues, reflect.ValueOf(value))
		case reflect.Bool:
			value, err := strconv.ParseBool(actionParam)
			if err != nil {
				return nil, err
			}

			paramValues = append(paramValues, reflect.ValueOf(value))
		default:
			return nil, errors.New("action has unsupported arguments")
		}
	}

	return paramValues, nil
}

func (t *terminal) OutputRaw(text string, category d2enum.TermCategory) {
	var line string

	for _, word := range strings.Split(text, " ") {
		if len(line) > 0 {
			line += " "
		}

		lineLength := len(line)
		wordLength := len(word)

		if lineLength+wordLength >= termColCountMax {
			t.outputHistory = append(t.outputHistory, termHistoryEntry{line, category})
			line = word
		} else {
			line += word
		}
	}

	t.outputHistory = append(t.outputHistory, termHistoryEntry{line, category})
}

func (t *terminal) Outputf(format string, params ...interface{}) {
	t.OutputRaw(fmt.Sprintf(format, params...), d2enum.TermCategoryNone)
}

func (t *terminal) OutputInfof(format string, params ...interface{}) {
	t.OutputRaw(fmt.Sprintf(format, params...), d2enum.TermCategoryInfo)
}

func (t *terminal) OutputWarningf(format string, params ...interface{}) {
	t.OutputRaw(fmt.Sprintf(format, params...), d2enum.TermCategoryWarning)
}

func (t *terminal) OutputErrorf(format string, params ...interface{}) {
	t.OutputRaw(fmt.Sprintf(format, params...), d2enum.TermCategoryError)
}

func (t *terminal) OutputClear() {
	t.outputHistory = nil
	t.outputIndex = 0
}

func (t *terminal) IsVisible() bool {
	return t.visState != termVisHidden
}

func (t *terminal) Hide() {
	if t.visState != termVisHidden {
		t.visState = termVisHiding
	}
}

func (t *terminal) Show() {
	if t.visState != termVisShown {
		t.visState = termVisShowing
	}
}

func (t *terminal) BindAction(name, description string, action interface{}) error {
	actionType := reflect.TypeOf(action)
	if actionType.Kind() != reflect.Func {
		return errors.New("action is not a function")
	}

	for i := 0; i < actionType.NumIn(); i++ {
		switch actionType.In(i).Kind() {
		case reflect.String:
		case reflect.Int:
		case reflect.Uint:
		case reflect.Float64:
		case reflect.Bool:
		default:
			return errors.New("action has unsupported arguments")
		}
	}

	t.actions[name] = termActionEntry{action, description}

	return nil
}

func (t *terminal) BindLogger() {
	log.SetOutput(&terminalLogger{writer: log.Writer(), terminal: t})
}

func (t *terminal) UnbindAction(name string) error {
	delete(t.actions, name)
	return nil
}

func easeInOut(t float64) float64 {
	t *= 2
	if t < 1 {
		return 0.5 * t * t * t * t
	}

	t -= 2

	return -0.5 * (t*t*t*t - 2)
}

func parseCommand(command string) []string {
	var (
		quoted bool
		escape bool
		param  string
		params []string
	)

	for _, c := range command {
		switch c {
		case '"':
			if escape {
				param += string(c)
				escape = false
			} else {
				quoted = !quoted
			}
		case ' ':
			if quoted {
				param += string(c)
			} else if len(param) > 0 {
				params = append(params, param)
				param = ""
			}
		case '\\':
			if escape {
				param += string(c)
				escape = false
			} else {
				escape = true
			}
		default:
			param += string(c)
		}
	}

	if len(param) > 0 {
		params = append(params, param)
	}

	return params
}

func rgbaColor(rgba uint32) color.RGBA {
	result := color.RGBA{}
	a, b, g, r := 0, 1, 2, 3
	byteWidth := 8
	byteMask := 0xff

	for idx := 0; idx < 4; idx++ {
		shift := idx * byteWidth
		component := uint8(rgba>>shift) & uint8(byteMask)

		switch idx {
		case a:
			result.A = component
		case b:
			result.B = component
		case g:
			result.G = component
		case r:
			result.R = component
		}
	}

	return result
}

func createTerminal() (*terminal, error) {
	terminal := &terminal{
		lineCount:    termRowCount,
		bgColor:      rgbaColor(darkGrey),
		fgColor:      rgbaColor(lightGrey),
		infoColor:    rgbaColor(lightBlue),
		warningColor: rgbaColor(yellow),
		errorColor:   rgbaColor(red),
		actions:      make(map[string]termActionEntry),
	}

	terminal.OutputInfof("::: OpenDiablo2 Terminal :::")
	terminal.OutputInfof("type \"ls\" for a list of actions")

	err := terminal.BindAction("ls", "list available actions", func() {
		var names []string
		for name := range terminal.actions {
			names = append(names, name)
		}

		sort.Strings(names)

		terminal.OutputInfof("available actions (%d):", len(names))
		for _, name := range names {
			entry := terminal.actions[name]
			terminal.OutputInfof("%s: %s; %s", name, entry.description, reflect.TypeOf(entry.action).String())
		}
	})
	if err != nil {
		return nil, fmt.Errorf("failed to bind the '%s' action, err: %w", "ls", err)
	}

	err = terminal.BindAction("clear", "clear terminal", func() {
		terminal.OutputClear()
	})
	if err != nil {
		return nil, fmt.Errorf("failed to bind the '%s' action, err: %w", "clear", err)
	}

	return terminal, nil
}
