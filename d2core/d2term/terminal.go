package d2term

import (
	"errors"
	"fmt"
	"image/color"
	"log"
	"math"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
)

const (
	charWidth       = 6
	charHeight      = 16
	charDoubleWidth = charWidth * 2
	rowCount        = 24
	rowCountMax     = 32
	colCountMax     = 128
	animLength      = 0.5
)

const (
	darkGrey  = 0x2e3436b0
	lightGrey = 0x555753b0
	lightBlue = 0x3465a4b0
	yellow    = 0xfce94fb0
	red       = 0xcc0000b0
)

type visibility int

const (
	visHidden visibility = iota
	visShowing
	visShown
	visHiding
)

const (
	maxVisAnim = 1.0
	minVisAnim = 0.0
)

type historyEntry struct {
	text     string
	category d2enum.TermCategory
}

type commandEntry struct {
	description string
	arguments   []string
	fn          func([]string) error
}

// Terminal handles the in-game terminal
type Terminal struct {
	outputHistory []historyEntry
	outputIndex   int

	command        string
	commandHistory []string
	commandIndex   int

	lineCount int
	visState  visibility
	visAnim   float64

	bgColor      color.RGBA
	fgColor      color.RGBA
	infoColor    color.RGBA
	warningColor color.RGBA
	errorColor   color.RGBA

	commands map[string]commandEntry
}

// NewTerminal creates and returns a terminal
func NewTerminal() (*Terminal, error) {
	term := &Terminal{
		lineCount:    rowCount,
		bgColor:      d2util.Color(darkGrey),
		fgColor:      d2util.Color(lightGrey),
		infoColor:    d2util.Color(lightBlue),
		warningColor: d2util.Color(yellow),
		errorColor:   d2util.Color(red),
		commands:     make(map[string]commandEntry),
	}

	term.Infof("::: OpenDiablo2 Terminal :::")
	term.Infof("type \"ls\" for a list of commands")

	if err := term.Bind("ls", "list available commands", nil, term.commandList); err != nil {
		return nil, err
	}

	if err := term.Bind("clear", "clear terminal", nil, term.commandClear); err != nil {
		return nil, err
	}

	return term, nil
}

// Bind binds commands to the terminal
func (t *Terminal) Bind(name, description string, arguments []string, fn func(args []string) error) error {
	if name == "" || description == "" {
		return fmt.Errorf("missing name or description")
	}

	if _, ok := t.commands[name]; ok {
		t.Warningf("rebinding command with name: %s", name)
	}

	t.commands[name] = commandEntry{description, arguments, fn}

	return nil
}

// Unbind unbinds commands from the terminal
func (t *Terminal) Unbind(names ...string) error {
	for _, name := range names {
		delete(t.commands, name)
	}

	return nil
}

// Advance advances the terminal animation
func (t *Terminal) Advance(elapsed float64) error {
	switch t.visState {
	case visShowing:
		t.visAnim = math.Min(maxVisAnim, t.visAnim+elapsed/animLength)
		if t.visAnim == maxVisAnim {
			t.visState = visShown
		}
	case visHiding:
		t.visAnim = math.Max(minVisAnim, t.visAnim-elapsed/animLength)
		if t.visAnim == minVisAnim {
			t.visState = visHidden
		}
	}

	if !t.Visible() {
		return nil
	}

	return nil
}

// OnKeyDown handles key down in the terminal
func (t *Terminal) OnKeyDown(event d2interface.KeyEvent) bool {
	if event.Key() == d2enum.KeyGraveAccent {
		t.toggle()
	}

	if !t.Visible() {
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

func (t *Terminal) processCommand() {
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

	t.Printf(t.command)

	if err := t.Execute(t.command); err != nil {
		t.Errorf(err.Error())
	}

	t.commandIndex = len(t.commandHistory) - 1
	t.command = ""
}

func (t *Terminal) handleControlKey(eventKey d2enum.Key, keyMod d2enum.KeyMod) {
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
			t.lineCount = d2math.MinInt(t.lineCount+1, rowCountMax)
		}
	}
}

// OnKeyChars handles char key in terminal
func (t *Terminal) OnKeyChars(event d2interface.KeyCharsEvent) bool {
	if !t.Visible() {
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

// Render renders the terminal
func (t *Terminal) Render(surface d2interface.Surface) error {
	if !t.Visible() {
		return nil
	}

	totalWidth, _ := surface.GetSize()
	outputHeight := t.lineCount * charHeight
	totalHeight := outputHeight + charHeight

	offset := -int((1 - easeInOut(t.visAnim)) * float64(totalHeight))
	surface.PushTranslation(0, offset)

	surface.DrawRect(totalWidth, outputHeight, t.bgColor)

	for i := 0; i < t.lineCount; i++ {
		historyIndex := len(t.outputHistory) - i - t.outputIndex - 1
		if historyIndex < 0 {
			break
		}

		entry := t.outputHistory[historyIndex]

		surface.PushTranslation(charDoubleWidth, outputHeight-(i+1)*charHeight)
		surface.DrawTextf(entry.text)
		surface.PushTranslation(-charDoubleWidth, 0)

		switch entry.category {
		case d2enum.TermCategoryInfo:
			surface.DrawRect(charWidth, charHeight, t.infoColor)
		case d2enum.TermCategoryWarning:
			surface.DrawRect(charWidth, charHeight, t.warningColor)
		case d2enum.TermCategoryError:
			surface.DrawRect(charWidth, charHeight, t.errorColor)
		}

		surface.Pop()
		surface.Pop()
	}

	surface.PushTranslation(0, outputHeight)
	surface.DrawRect(totalWidth, charHeight, t.fgColor)
	surface.DrawTextf("> " + t.command)
	surface.Pop()

	surface.Pop()

	return nil
}

// Execute executes a command with arguments
func (t *Terminal) Execute(command string) error {
	params := parseCommand(command)
	if len(params) == 0 {
		return errors.New("invalid command")
	}

	name := params[0]
	args := params[1:]

	entry, ok := t.commands[name]
	if !ok {
		return errors.New("command not found")
	}

	if len(args) != len(entry.arguments) {
		return errors.New("command requires different argument count")
	}

	if err := entry.fn(args); err != nil {
		return err
	}

	return nil
}

// Rawf writes a raw message to the terminal
func (t *Terminal) Rawf(category d2enum.TermCategory, format string, params ...interface{}) {
	text := fmt.Sprintf(format, params...)
	lines := d2util.SplitIntoLinesWithMaxWidth(text, colCountMax)

	for _, line := range lines {
		// removes color token (this token ends with [0m )
		l := strings.Split(line, "\033[0m")
		line = l[len(l)-1]

		t.outputHistory = append(t.outputHistory, historyEntry{line, category})
	}
}

// Printf writes a message to the terminal
func (t *Terminal) Printf(format string, params ...interface{}) {
	t.Rawf(d2enum.TermCategoryNone, format, params...)
}

// Infof writes a warning message to the terminal
func (t *Terminal) Infof(format string, params ...interface{}) {
	t.Rawf(d2enum.TermCategoryInfo, format, params...)
}

// Warningf writes a warning message to the terminal
func (t *Terminal) Warningf(format string, params ...interface{}) {
	t.Rawf(d2enum.TermCategoryWarning, format, params...)
}

// Errorf writes a error message to the terminal
func (t *Terminal) Errorf(format string, params ...interface{}) {
	t.Rawf(d2enum.TermCategoryError, format, params...)
}

// Clear clears the terminal
func (t *Terminal) Clear() {
	t.outputHistory = nil
	t.outputIndex = 0
}

// Visible returns visible state
func (t *Terminal) Visible() bool {
	return t.visState != visHidden
}

// Hide hides the terminal
func (t *Terminal) Hide() {
	if t.visState != visHidden {
		t.visState = visHiding
	}
}

// Show shows the terminal
func (t *Terminal) Show() {
	if t.visState != visShown {
		t.visState = visShowing
	}
}

func (t *Terminal) toggle() {
	if t.visState == visHiding || t.visState == visHidden {
		t.Show()
		return
	}

	t.Hide()
}

// BindLogger binds a log.Writer to the output
func (t *Terminal) BindLogger() {
	log.SetOutput(&terminalLogger{writer: log.Writer(), terminal: t})
}

func easeInOut(t float64) float64 {
	t *= 2
	if t < 1 {
		return 0.5 * t * t * t * t
	}

	t -= 2

	// nolint:gomnd // constant
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
