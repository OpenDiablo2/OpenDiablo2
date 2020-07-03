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

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

)

// TermCategory applies styles to the lines in the  Terminal
type TermCategory d2interface.TermCategory

// Terminal Category types
const (
	TermCategoryNone    = TermCategory(d2interface.TermCategoryNone)
	TermCategoryInfo    = TermCategory(d2interface.TermCategoryInfo)
	TermCategoryWarning = TermCategory(d2interface.TermCategoryWarning)
	TermCategoryError   = TermCategory(d2interface.TermCategoryError)
)
const (
	termCharWidth   = 6
	termCharHeight  = 16
	termRowCount    = 24
	termRowCountMax = 32
	termColCountMax = 128
	termAnimLength  = 0.5
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
	category d2interface.TermCategory
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
	if event.Key() == d2interface.KeyGraveAccent {
		t.toggleTerminal()
	}

	if !t.IsVisible() {
		return false
	}

	switch event.Key() {
	case d2interface.KeyEscape:
		t.command = ""
	case d2interface.KeyEnd:
		t.outputIndex = 0
	case d2interface.KeyHome:
		t.outputIndex = d2common.MaxInt(0, len(t.outputHistory)-t.lineCount)
	case d2interface.KeyPageUp:
		maxOutputIndex := d2common.MaxInt(0, len(t.outputHistory)-t.lineCount)
		if t.outputIndex += t.lineCount; t.outputIndex >= maxOutputIndex {
			t.outputIndex = maxOutputIndex
		}
	case d2interface.KeyPageDown:
		if t.outputIndex -= t.lineCount; t.outputIndex < 0 {
			t.outputIndex = 0
		}
	case d2interface.KeyUp, d2interface.KeyDown:
		t.handleControlKey(event.Key(), event.KeyMod())
	case d2interface.KeyEnter:
		t.processCommand()
	case d2interface.KeyBackspace:
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

func (t *terminal) handleControlKey(eventKey d2interface.Key, keyMod d2interface.KeyMod) {
	switch eventKey {
	case d2interface.KeyUp:
		if keyMod == d2interface.KeyModControl {
			t.lineCount = d2common.MaxInt(0, t.lineCount-1)
		} else if len(t.commandHistory) > 0 {
			t.command = t.commandHistory[t.commandIndex]
			if t.commandIndex == 0 {
				t.commandIndex = len(t.commandHistory) - 1
			} else {
				t.commandIndex--
			}
		}
	case d2interface.KeyDown:
		if keyMod == d2interface.KeyModControl {
			t.lineCount = d2common.MinInt(t.lineCount+1, termRowCountMax)
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

		surface.PushTranslation(termCharWidth*2, outputHeight-(i+1)*termCharHeight)
		surface.DrawText(historyEntry.text)
		surface.PushTranslation(-termCharWidth*2, 0)

		switch historyEntry.category {
		case d2interface.TermCategoryInfo:
			surface.DrawRect(termCharWidth, termCharHeight, t.infoColor)
		case d2interface.TermCategoryWarning:
			surface.DrawRect(termCharWidth, termCharHeight, t.warningColor)
		case d2interface.TermCategoryError:
			surface.DrawRect(termCharWidth, termCharHeight, t.errorColor)
		}

		surface.Pop()
		surface.Pop()
	}

	surface.PushTranslation(0, outputHeight)
	surface.DrawRect(totalWidth, termCharHeight, t.fgColor)
	surface.DrawText("> " + t.command)
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

func (t *terminal) OutputRaw(text string, category d2interface.TermCategory) {
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
	t.OutputRaw(fmt.Sprintf(format, params...), d2interface.TermCategoryNone)
}

func (t *terminal) OutputInfof(format string, params ...interface{}) {
	t.OutputRaw(fmt.Sprintf(format, params...), d2interface.TermCategoryInfo)
}

func (t *terminal) OutputWarningf(format string, params ...interface{}) {
	t.OutputRaw(fmt.Sprintf(format, params...), d2interface.TermCategoryWarning)
}

func (t *terminal) OutputErrorf(format string, params ...interface{}) {
	t.OutputRaw(fmt.Sprintf(format, params...), d2interface.TermCategoryError)
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

func createTerminal() (*terminal, error) {
	terminal := &terminal{
		lineCount:    termRowCount,
		bgColor:      color.RGBA{R: 0x2e, G: 0x34, B: 0x36, A: 0xb0},
		fgColor:      color.RGBA{R: 0x55, G: 0x57, B: 0x53, A: 0xb0},
		infoColor:    color.RGBA{R: 0x34, G: 0x65, B: 0xa4, A: 0xb0},
		warningColor: color.RGBA{R: 0xfc, G: 0xe9, B: 0x4f, A: 0xb0},
		errorColor:   color.RGBA{R: 0xcc, A: 0xb0},
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
