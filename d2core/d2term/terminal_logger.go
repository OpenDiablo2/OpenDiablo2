package d2term

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type terminalLogger struct {
	terminal *Terminal
	buffer   bytes.Buffer
	writer   io.Writer
}

func (tl *terminalLogger) Write(p []byte) (int, error) {
	n, err := tl.buffer.Write(p)
	if err != nil {
		return n, err
	}

	reader := bufio.NewReader(&tl.buffer)
	termBytes, _, err := reader.ReadLine()

	if err != nil {
		return n, err
	}

	line := string(termBytes)
	lineLower := strings.ToLower(line)

	switch {
	case strings.Index(lineLower, "error") > 0:
		tl.terminal.Errorf(line)
	case strings.Index(lineLower, "warning") > 0:
		tl.terminal.Errorf(line)
	default:
		tl.terminal.Printf(line)
	}

	return tl.writer.Write(p)
}

func (tl *terminalLogger) BindToTerminal(t *Terminal) {
	tl.terminal = t
}
