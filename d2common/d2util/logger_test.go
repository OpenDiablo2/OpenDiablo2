package d2util

import (
	"fmt"
	"testing"
)

type testWriter struct {
	data []byte
}

func (tw *testWriter) Write(msg []byte) (int, error) {
	tw.data = msg

	return len(msg), nil
}

func Test_logger_SetLevel(t *testing.T) {
	l := NewLogger()
	l.Writer = &testWriter{}

	tests := []struct {
		level LogLevel
	}{
		{LogLevelNone},
		{LogLevelError},
		{LogLevelWarning},
		{LogLevelInfo},
		{LogLevelDebug},
	}

	for idx := range tests {
		targetLevel := tests[idx].level
		l.SetLevel(targetLevel)

		if l.level != targetLevel {
			t.Error("unexpected log level")
		}
	}
}

func Test_logger_LogLevels(t *testing.T) {
	l := NewLogger()
	w := &testWriter{}
	l.Writer = w

	noMessage := ""
	message := "test"
	expectedError := fmt.Sprintf(LogFmtError, message)
	expectedWarning := fmt.Sprintf(LogFmtWarning, message)
	expectedInfo := fmt.Sprintf(LogFmtInfo, message)
	expectedDebug := fmt.Sprintf(LogFmtDebug, message)

	// for each log level we set, we will use different log methods (info, warning, etc) and check
	// what the output in the writer is (clearing the writer data before each test)
	tests := []struct {
		logLevel LogLevel
		expect   map[LogLevel]string
	}{
		{LogLevelDebug, map[LogLevel]string{
			LogLevelError:   expectedError,
			LogLevelWarning: expectedWarning,
			LogLevelInfo:    expectedInfo,
			LogLevelDebug:   expectedDebug,
		}},
		{LogLevelInfo, map[LogLevel]string{
			LogLevelError:   expectedError,
			LogLevelWarning: expectedWarning,
			LogLevelInfo:    expectedInfo,
			LogLevelDebug:   noMessage,
		}},
		{LogLevelWarning, map[LogLevel]string{
			LogLevelError:   expectedError,
			LogLevelWarning: expectedWarning,
			LogLevelInfo:    noMessage,
			LogLevelDebug:   noMessage,
		}},
		{LogLevelError, map[LogLevel]string{
			LogLevelError:   expectedError,
			LogLevelWarning: noMessage,
			LogLevelInfo:    noMessage,
			LogLevelDebug:   noMessage,
		}},
		{LogLevelNone, map[LogLevel]string{
			LogLevelError:   noMessage,
			LogLevelWarning: noMessage,
			LogLevelInfo:    noMessage,
			LogLevelDebug:   noMessage,
		}},
	}

	for idx := range tests {
		level := tests[idx].logLevel
		l.SetLevel(level)

		for levelTry, msgExpect := range tests[idx].expect {
			w.data = make([]byte, 0)

			switch levelTry {
			case LogLevelError:
				l.Error(message)
			case LogLevelWarning:
				l.Warning(message)
			case LogLevelInfo:
				l.Info(message)
			case LogLevelDebug:
				l.Debug(message)
			}

			msgGot := string(w.data)

			if len(msgGot) > 0 && len(msgExpect) < 1 {
				t.Errorf("logger printed when it should not have")
			}

			if len(msgGot) < 1 && len(msgExpect) > 0 {
				t.Errorf("logger didnt print when expected")
			}
		}
	}
}
