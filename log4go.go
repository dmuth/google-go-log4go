
package log4go

import "fmt"
import "strings"
import "time"

const (
	ErrorLevel = iota
	WarnLevel = iota
	InfoLevel = iota
	DebugLevel = iota
	TraceLevel = iota
)

//
// Our current logging level
//
var _level int

//
// Our starting timestamp
//
var startTime time.Time
var beenhere bool

//
// Do we want to display timestamps?
//
var displayTime bool


/**
* Set our logging level.
* @param {int} The level to set to.
*/
func SetLevel(level int) {
	_level = level
}


/**
* Set our logging level based on a string.
* @param {string} level The level to set. This is case-insensitive.
*/
func SetLevelString(level string) {
	level = strings.ToLower(level)
	if (level == "error") {
		SetLevel(ErrorLevel)
	} else if (level == "warn") {
		SetLevel(WarnLevel)
	} else if (level == "info") {
		SetLevel(InfoLevel)
	} else if (level == "debug") {
		SetLevel(DebugLevel)
	} else if (level == "trace") {
		SetLevel(TraceLevel)
	}

} // End of SetLevelString()


/**
* Retrieve our current logging level.
* @return {int} The current level
*/
func Level() int {
	return(_level)
}


/**
* Get our display time value
* @return {bool} True if we are displaying timestamps in messages
*/
func DisplayTime() bool {
	return (displayTime)
}


/**
* Set our display time setting
* @param {bool} display True if we want to display timestamps in messages
*/
func SetDisplayTime(display bool) {
	displayTime = display
}


/**
* Log an error
*/
func Error(message string) {
	print(ErrorLevel, "ERROR: " + message);
}


/**
* Log a warning
*/
func Warn(message string) {
	print(WarnLevel, "WARN: " + message)
}


/**
* Log info
*/
func Info(message string) {
	print(InfoLevel, "INFO: " + message)
}


/**
* Log debugigng
*/
func Debug(message string) {
	print(DebugLevel, "DEBUG: " + message)
}


/**
* Log a trace 
*/
func Trace(message string) {
	print (TraceLevel, "TRACE: " + message)
}


/**
* Central function for printing
*/
func print(level int, message string) {

	if (!beenhere) {
		startTime = time.Now()
		beenhere = true
	}

	if (_level >= level) {

		if (displayTime) {
			now := time.Now()
			elapsed := now.Sub(startTime)
			fmt.Printf("[%s] %s\n", elapsed, message)

		} else {
			fmt.Println(message)

		}

	}

} // End of print()



