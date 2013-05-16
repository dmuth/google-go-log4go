
package log4go

import "testing"


func TestMain(t *testing.T) {

	var level int
	SetLevel(ErrorLevel)
	level = Level()
	if (level != ErrorLevel) {
		t.Errorf("Level %d != %d", level, Error)
	}

	SetLevel(WarnLevel)
	level = Level()
	if (level != WarnLevel) {
		t.Errorf("Level %d != %d", level, Error)
	}

} // End of TestMain()


func ExampleError() {
	SetLevel(ErrorLevel)
	Error("test")
	Warn("test")
	// Output:
	// ERROR: test
}

func ExampleWarn() {
	SetLevel(WarnLevel)
	Error("test")
	Warn("test")
	Info("test")
	// Output:
	// ERROR: test
	// WARN: test
}

func ExampleInfo() {
	SetLevel(InfoLevel)
	Error("test")
	Warn("test")
	Info("test")
	Debug("test")
	// Output:
	// ERROR: test
	// WARN: test
	// INFO: test
}

func ExampleDebug() {
	SetLevel(DebugLevel)
	Error("test")
	Warn("test")
	Info("test")
	Debug("test")
	Trace("test")
	// Output:
	// ERROR: test
	// WARN: test
	// INFO: test
	// DEBUG: test
}

func ExampleTrace() {
	SetLevel(TraceLevel)
	Error("test")
	Warn("test")
	Info("test")
	Debug("test")
	Trace("test")
	// Output:
	// ERROR: test
	// WARN: test
	// INFO: test
	// DEBUG: test
	// TRACE: test
}


