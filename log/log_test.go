package log

import "testing"

func TestLogInfo(t *testing.T) {
	Info("hello info")
}

func TestLogDebug(t *testing.T) {
	Debug("hello debug")
}

func TestLogWaring(t *testing.T) {
	Waring("hello Waring")
}

func TestLogError(t *testing.T) {
	Error("hello", "Error")
}
