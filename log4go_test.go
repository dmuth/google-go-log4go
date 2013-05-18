
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


/**
* Central function to log things that is called by the other testing 
* functions.
*/
func logStuff() {
	Error("test")
	Warn("test")
	Info("test")
	Debug("test")
	Trace("test")
}

/**
* Test SetLevel()
*/
func ExampleError() {
	SetLevel(ErrorLevel)
	logStuff()
	// Output:
	// ERROR: test
}

func ExampleWarn() {
	SetLevel(WarnLevel)
	logStuff()
	// Output:
	// ERROR: test
	// WARN: test
}

func ExampleInfo() {
	SetLevel(InfoLevel)
	logStuff()
	// Output:
	// ERROR: test
	// WARN: test
	// INFO: test
}

func ExampleDebug() {
	SetLevel(DebugLevel)
	logStuff()
	// Output:
	// ERROR: test
	// WARN: test
	// INFO: test
	// DEBUG: test
}

func ExampleTrace() {
	SetLevel(TraceLevel)
	logStuff()
	// Output:
	// ERROR: test
	// WARN: test
	// INFO: test
	// DEBUG: test
	// TRACE: test
}


/**
* Test SetLevelString()
*/
func ExampleErrorString() {
	SetLevel(DebugLevel)
	SetLevelString("error")
	logStuff()
	// Output:
	// ERROR: test
}

func ExampleWarnString() {
	SetLevel(ErrorLevel)
	SetLevelString("warn")
	logStuff()
	// Output:
	// ERROR: test
	// WARN: test
}

func ExampleInfoString() {
	SetLevel(ErrorLevel)
	SetLevelString("INFO")
	logStuff()
	// Output:
	// ERROR: test
	// WARN: test
	// INFO: test
}

func ExampleDebugString() {
	SetLevel(ErrorLevel)
	SetLevelString("Debug")
	logStuff()
	// Output:
	// ERROR: test
	// WARN: test
	// INFO: test
	// DEBUG: test
}

func ExampleTraceString() {
	SetLevel(ErrorLevel)
	SetLevelString("trace")
	logStuff()
	// Output:
	// ERROR: test
	// WARN: test
	// INFO: test
	// DEBUG: test
	// TRACE: test
}



