package levellogger

import "testing"

import "strconv"

func TestLogger_Info(t *testing.T) {
	var l Logger
	l.LogLevel = InfoLevel
	i, e := strconv.Atoi("1q")
	l.Info("Parser Error ", e)
	if i != 0 {
		t.Fail()
	}
}

func TestLogger_Debug(t *testing.T) {
	var l Logger
	l.LogLevel = DebugLevel
	i, e := strconv.Atoi("1q")
	l.Debug(e)
	if i != 0 {
		t.Fail()
	}
}

func TestLogger_Error(t *testing.T) {
	var l Logger
	i, e := strconv.Atoi("1q")
	l.Error("A big error: ", e)
	if i != 0 {
		t.Fail()
	}
}

func TestLogger_InfoAll(t *testing.T) {
	var l Logger
	l.LogLevel = AllLevel
	i, e := strconv.Atoi("1q")
	l.Info(e)
	if i != 0 {
		t.Fail()
	}
}

func TestLogger_DebugAll(t *testing.T) {
	var l Logger
	l.LogLevel = AllLevel
	i, e := strconv.Atoi("1q")
	l.Debug(e)
	if i != 0 {
		t.Fail()
	}
}

func TestLogger_ErrorAll(t *testing.T) {
	var l Logger
	i, e := strconv.Atoi("1q")
	l.Error(e)
	if i != 0 {
		t.Fail()
	}
}

func TestLogger_InfoOff(t *testing.T) {
	var l Logger
	i, e := strconv.Atoi("1q")
	l.Info(e)
	if i != 0 {
		t.Fail()
	}
}

func TestLogger_DebugOff(t *testing.T) {
	var l Logger
	i, e := strconv.Atoi("1q")
	l.Debug(e)
	if i != 0 {
		t.Fail()
	}
}

func TestLogger_ErrorOff(t *testing.T) {
	var l Logger
	i, e := strconv.Atoi("1q")
	l.Error(e)
	if i != 0 {
		t.Fail()
	}
}
