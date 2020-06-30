package d2term

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type terminalLogger struct {
	terminal *terminal
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
		tl.terminal.OutputErrorf(line)
	case strings.Index(lineLower, "warning") > 0:
		tl.terminal.OutputWarningf(line)
	default:
		tl.terminal.Outputf(line)
	}

	return tl.writer.Write(p)
}

func (tl *terminalLogger) BindToTerminal(t *terminal) {
	tl.terminal = t
}
