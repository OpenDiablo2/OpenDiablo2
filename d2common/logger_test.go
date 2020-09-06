package d2common

import (
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
	l := &Logger{Writer: &testWriter{}}

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

func Test_logger_Debug(t *testing.T) {

}

func Test_logger_Error(t *testing.T) {

}

func Test_logger_Info(t *testing.T) {

}

func Test_logger_Warning(t *testing.T) {

}
