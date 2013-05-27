
package log4go

import "testing"

//import "fmt"


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
	Errorf("test: %s, %s, %d", "foo", "bar", 123)
	Warn("test")
	Warnf("test: %s, %s, %d", "foo", "bar", 123)
	Info("test")
	Infof("test: %s, %s, %d", "foo", "bar", 123)
	Debug("test")
	Debugf("test: %s, %s, %d", "foo", "bar", 123)
	Trace("test")
	Tracef("test: %s, %s, %d", "foo", "bar", 123)
}

/**
* Test SetLevel()
*/
func ExampleError() {
	SetLevel(ErrorLevel)
	logStuff()
	// Output:
	// ERROR: test
	// ERROR: test: foo, bar, 123
}

func ExampleWarn() {
	SetLevel(WarnLevel)
	logStuff()
	// Output:
	// ERROR: test
	// ERROR: test: foo, bar, 123
	// WARN: test
	// WARN: test: foo, bar, 123
}

func ExampleInfo() {
	SetLevel(InfoLevel)
	logStuff()
	// Output:
	// ERROR: test
	// ERROR: test: foo, bar, 123
	// WARN: test
	// WARN: test: foo, bar, 123
	// INFO: test
	// INFO: test: foo, bar, 123
}

func ExampleDebug() {
	SetLevel(DebugLevel)
	logStuff()
	// Output:
	// ERROR: test
	// ERROR: test: foo, bar, 123
	// WARN: test
	// WARN: test: foo, bar, 123
	// INFO: test
	// INFO: test: foo, bar, 123
	// DEBUG: test
	// DEBUG: test: foo, bar, 123
}

func ExampleTrace() {
	SetLevel(TraceLevel)
	logStuff()
	// Output:
	// ERROR: test
	// ERROR: test: foo, bar, 123
	// WARN: test
	// WARN: test: foo, bar, 123
	// INFO: test
	// INFO: test: foo, bar, 123
	// DEBUG: test
	// DEBUG: test: foo, bar, 123
	// TRACE: test
	// TRACE: test: foo, bar, 123
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
	// ERROR: test: foo, bar, 123
}

func ExampleWarnString() {
	SetLevel(ErrorLevel)
	SetLevelString("warn")
	logStuff()
	// Output:
	// ERROR: test
	// ERROR: test: foo, bar, 123
	// WARN: test
	// WARN: test: foo, bar, 123
}

func ExampleInfoString() {
	SetLevel(ErrorLevel)
	SetLevelString("INFO")
	logStuff()
	// Output:
	// ERROR: test
	// ERROR: test: foo, bar, 123
	// WARN: test
	// WARN: test: foo, bar, 123
	// INFO: test
	// INFO: test: foo, bar, 123
}

func ExampleDebugString() {
	SetLevel(ErrorLevel)
	SetLevelString("Debug")
	logStuff()
	// Output:
	// ERROR: test
	// ERROR: test: foo, bar, 123
	// WARN: test
	// WARN: test: foo, bar, 123
	// INFO: test
	// INFO: test: foo, bar, 123
	// DEBUG: test
	// DEBUG: test: foo, bar, 123
}

func ExampleTraceString() {
	SetLevel(ErrorLevel)
	SetLevelString("trace")
	logStuff()
	// Output:
	// ERROR: test
	// ERROR: test: foo, bar, 123
	// WARN: test
	// WARN: test: foo, bar, 123
	// INFO: test
	// INFO: test: foo, bar, 123
	// DEBUG: test
	// DEBUG: test: foo, bar, 123
	// TRACE: test
	// TRACE: test: foo, bar, 123
}

func ExampleMultiLine() {
	SetLevel(ErrorLevel)
	SetLevelString("trace")
	Error("test\ntest2")
	Error("test3\rtest4")
	// Output:
	// ERROR: test\ntest2
	// ERROR: test3\rtest4
}


func ExampleBackspaceAndMore() {
	SetLevelString("info")
	Error("test" + string(0) + "test2" + string(31) + "test3" + string(8) + string(2)  + "test4")
	//
	// Output: 
	// ERROR: test[0x00]test2[0x1f]test3[0x08][0x02]test4
}


