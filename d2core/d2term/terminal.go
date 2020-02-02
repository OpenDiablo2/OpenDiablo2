package d2term

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"image/color"
	"io"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

const (
	termCharWidth   = 6
	termCharHeight  = 16
	termRowCount    = 24
	termRowCountMax = 32
	termColCountMax = 128
	termAnimLength  = 0.5
)

type termCategory int

const (
	termCategoryNone termCategory = iota
	termCategoryInfo
	termCategoryWarning
	termCategoryError
)

type termVis int

const (
	termVisHidden termVis = iota
	termVisShowing
	termVisShown
	termVisHiding
)

var (
	termBgColor      = color.RGBA{R: 0x2e, G: 0x34, B: 0x36, A: 0xb0}
	termFgColor      = color.RGBA{R: 0x55, G: 0x57, B: 0x53, A: 0xb0}
	termInfoColor    = color.RGBA{R: 0x34, G: 0x65, B: 0xa4, A: 0xb0}
	termWarningColor = color.RGBA{R: 0xfc, G: 0xe9, B: 0x4f, A: 0xb0}
	termErrorColor   = color.RGBA{R: 0xcc, A: 0xb0}
)

type termHistroyEntry struct {
	text     string
	category termCategory
}

type termActionEntry struct {
	action      interface{}
	description string
}

type terminal struct {
	outputHistory []termHistroyEntry
	outputIndex   int

	command        string
	commandHistory []string
	commandIndex   int

	lineCount int
	visState  termVis
	visAnim   float64

	actions map[string]termActionEntry
}

func createTerminal() (*terminal, error) {
	terminal := &terminal{
		lineCount: termRowCount,
		actions:   make(map[string]termActionEntry),
	}

	terminal.outputInfo("::: OpenDiablo2 Terminal :::")
	terminal.outputInfo("type \"ls\" for a list of actions")

	terminal.bindAction("ls", "list available actions", func() {
		var names []string
		for name := range terminal.actions {
			names = append(names, name)
		}

		sort.Strings(names)

		terminal.outputInfo("available actions (%d):", len(names))
		for _, name := range names {
			entry := terminal.actions[name]
			terminal.outputInfo("%s: %s; %s", name, entry.description, reflect.TypeOf(entry.action).String())
		}
	})
	terminal.bindAction("clear", "clear terminal", func() {
		terminal.outputClear()
	})

	return terminal, nil
}

func (t *terminal) advance(elapsed float64) error {
	switch t.visState {
	case termVisShowing:
		t.visAnim = math.Min(1.0, t.visAnim+elapsed/termAnimLength)
		if t.visAnim == 1.0 {
			t.visState = termVisShown
		}
	case termVisHiding:
		t.visAnim = math.Max(0.0, t.visAnim-elapsed/termAnimLength)
		if t.visAnim == 0.0 {
			t.visState = termVisHidden
		}
	}

	if !t.isVisible() {
		return nil
	}

	return nil
}

func (t *terminal) OnKeyDown(event d2input.KeyEvent) bool {
	maxOutputIndex := d2common.MaxInt(0, len(t.outputHistory)-t.lineCount)

	if t.visState == termVisHiding || t.visState == termVisHidden {
		if event.Key == d2input.KeyGraveAccent {
			t.show()
			return true
		}

		// When terminal is not visible, only d2input.KeyGraveAccent is consumed.
		return false
	}

	if event.Key == d2input.KeyGraveAccent {
		t.hide()
		return true
	}

	if event.Key == d2input.KeyEscape {
		t.command = ""
		return true
	}

	if event.Key == d2input.KeyHome {
		t.outputIndex = maxOutputIndex
		return true
	}

	if event.Key == d2input.KeyEnd {
		t.outputIndex = 0
		return true
	}

	if event.Key == d2input.KeyPageUp {
		if t.outputIndex += t.lineCount; t.outputIndex >= maxOutputIndex {
			t.outputIndex = maxOutputIndex
		}

		return true
	}

	if event.Key == d2input.KeyPageDown {
		if t.outputIndex -= t.lineCount; t.outputIndex < 0 {
			t.outputIndex = 0
		}

		return true
	}

	if event.Key == d2input.KeyUp {
		if event.KeyMod == d2input.KeyModControl {
			t.lineCount = d2common.MaxInt(0, t.lineCount-1)
		} else if len(t.commandHistory) > 0 {
			t.command = t.commandHistory[t.commandIndex]
			if t.commandIndex == 0 {
				t.commandIndex = len(t.commandHistory) - 1
			} else {
				t.commandIndex--
			}
		}

		return true
	}

	if event.Key == d2input.KeyDown && event.KeyMod == d2input.KeyModControl {
		t.lineCount = d2common.MinInt(t.lineCount+1, termRowCountMax)
		return true
	}

	if event.Key == d2input.KeyEnter && len(t.command) > 0 {
		var commandHistory []string
		for _, command := range t.commandHistory {
			if command != t.command {
				commandHistory = append(commandHistory, command)
			}
		}

		t.commandHistory = append(commandHistory, t.command)

		t.output(t.command)
		if err := t.execute(t.command); err != nil {
			t.outputError(err.Error())
		}

		t.commandIndex = len(t.commandHistory) - 1
		t.command = ""

		return true
	}

	if event.Key == d2input.KeyBackspace && len(t.command) > 0 {
		t.command = t.command[:len(t.command)-1]
		return true
	}

	// When terminal is visible, all keys are consumed.
	return true
}

func (t *terminal) OnKeyChars(event d2input.KeyCharsEvent) bool {
	var handled bool
	for _, c := range event.Chars {
		if c != '`' {
			t.command += string(c)
			handled = true
		}
	}

	return handled
}

func (t *terminal) render(surface d2render.Surface) error {
	if !t.isVisible() {
		return nil
	}

	totalWidth, _ := surface.GetSize()
	outputHeight := t.lineCount * termCharHeight
	totalHeight := outputHeight + termCharHeight

	offset := -int((1.0 - easeInOut(t.visAnim)) * float64(totalHeight))
	surface.PushTranslation(0, offset)

	surface.DrawRect(totalWidth, outputHeight, termBgColor)

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
		case termCategoryInfo:
			surface.DrawRect(termCharWidth, termCharHeight, termInfoColor)
		case termCategoryWarning:
			surface.DrawRect(termCharWidth, termCharHeight, termWarningColor)
		case termCategoryError:
			surface.DrawRect(termCharWidth, termCharHeight, termErrorColor)
		}
		surface.Pop()
		surface.Pop()
	}

	surface.PushTranslation(0, outputHeight)
	surface.DrawRect(totalWidth, termCharHeight, termFgColor)
	surface.DrawText("> " + t.command)
	surface.Pop()

	surface.Pop()

	return nil
}

func (t *terminal) execute(command string) error {
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

	var paramValues []reflect.Value
	for i := 0; i < actionType.NumIn(); i++ {
		actionParam := actionParams[i]
		switch actionType.In(i).Kind() {
		case reflect.String:
			paramValues = append(paramValues, reflect.ValueOf(actionParam))
		case reflect.Int:
			value, err := strconv.ParseInt(actionParam, 10, 64)
			if err != nil {
				return err
			}
			paramValues = append(paramValues, reflect.ValueOf(int(value)))
		case reflect.Uint:
			value, err := strconv.ParseUint(actionParam, 10, 64)
			if err != nil {
				return err
			}
			paramValues = append(paramValues, reflect.ValueOf(uint(value)))
		case reflect.Float64:
			value, err := strconv.ParseFloat(actionParam, 64)
			if err != nil {
				return err
			}
			paramValues = append(paramValues, reflect.ValueOf(value))
		case reflect.Bool:
			value, err := strconv.ParseBool(actionParam)
			if err != nil {
				return err
			}
			paramValues = append(paramValues, reflect.ValueOf(value))
		default:
			return errors.New("action has unsupported arguments")
		}
	}

	actionValue := reflect.ValueOf(actionEntry.action)
	actionReturnValues := actionValue.Call(paramValues)

	if actionReturnValueCount := len(actionReturnValues); actionReturnValueCount > 0 {
		t.outputInfo("function returned %d values:", actionReturnValueCount)
		for _, actionReturnValue := range actionReturnValues {
			t.outputInfo("%v: %s", actionReturnValue.Interface(), actionReturnValue.String())
		}
	}

	return nil
}

func (t *terminal) outputRaw(text string, category termCategory) error {
	var line string
	for _, word := range strings.Split(text, " ") {
		if len(line) > 0 {
			line += " "
		}

		lineLength := len(line)
		wordLength := len(word)

		if lineLength+wordLength >= termColCountMax {
			t.outputHistory = append(t.outputHistory, termHistroyEntry{line, category})
			line = word
		} else {
			line += word
		}
	}

	t.outputHistory = append(t.outputHistory, termHistroyEntry{line, category})
	return nil
}

func (t *terminal) output(format string, params ...interface{}) error {
	return t.outputRaw(fmt.Sprintf(format, params...), termCategoryNone)
}

func (t *terminal) outputInfo(format string, params ...interface{}) error {
	return t.outputRaw(fmt.Sprintf(format, params...), termCategoryInfo)
}

func (t *terminal) outputWarning(format string, params ...interface{}) error {
	return t.outputRaw(fmt.Sprintf(format, params...), termCategoryWarning)
}

func (t *terminal) outputError(format string, params ...interface{}) error {
	return t.outputRaw(fmt.Sprintf(format, params...), termCategoryError)
}

func (t *terminal) outputClear() {
	t.outputHistory = nil
	t.outputIndex = 0
}

func (t *terminal) isVisible() bool {
	return t.visState != termVisHidden
}

func (t *terminal) hide() {
	if t.visState != termVisHidden {
		t.visState = termVisHiding
	}
}

func (t *terminal) show() {
	if t.visState != termVisShown {
		t.visState = termVisShowing
	}
}

func (t *terminal) bindAction(name, description string, action interface{}) error {
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
			break
		default:
			return errors.New("action has unsupported arguments")
		}
	}

	t.actions[name] = termActionEntry{action, description}
	return nil
}

func (t *terminal) unbindAction(name string) error {
	delete(t.actions, name)
	return nil
}

type terminalLogger struct {
	buffer bytes.Buffer
	writer io.Writer
}

func (t *terminalLogger) Write(p []byte) (int, error) {
	n, err := t.buffer.Write(p)
	if err != nil {
		return n, err
	}

	reader := bufio.NewReader(&t.buffer)
	termBytes, _, err := reader.ReadLine()
	if err != nil {
		return n, err
	}

	line := string(termBytes[:])
	lineLower := strings.ToLower(line)

	if strings.Index(lineLower, "error") > 0 {
		OutputError(line)
	} else if strings.Index(lineLower, "warning") > 0 {
		OutputWarning(line)
	} else {
		Output(line)
	}

	return t.writer.Write(p)
}

func easeInOut(t float64) float64 {
	t *= 2
	if t < 1 {
		return 0.5 * t * t * t * t
	} else {
		t -= 2
		return -0.5 * (t*t*t*t - 2)
	}
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
