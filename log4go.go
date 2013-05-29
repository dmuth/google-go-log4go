
package log4go

import "fmt"
import "strings"
import "regexp"
import "time"


import "github.com/wsxiaoys/terminal/color"


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
// Since I can't initialize a boolean to true, I'm going to make this an 
// int with the following values:
// 0 - uninitialized (defaults to true)
// 1 - false
// 2 - true
//
// I'm not sure if this is the best method, as I am still learning Go, 
// but it'll work for now!
//
var displayTime int


//
// Same thing with color.  Do we want it?
//
// 0 - uninitialized (defaults to true)
// 1 - false
// 2 - true
//
var useColor int


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
	if (displayTime == 0 || displayTime == 2) {
		return(true)
	}

	return(false)

}


/**
* Set our display time setting
* @param {bool} display True if we want to display timestamps in messages
*/
func SetDisplayTime(display bool) {
	if (display) {
		displayTime = 2
	} else {
		displayTime = 1
	}
}


/**
* Get our use color value
* @return {bool} True if we are using color in messages
*/
func UseColor() bool {
	if (useColor == 0 || useColor == 2) {
		return(true)
	}

	return(false)

}


/**
* Set our use color settings
* @param {bool} use_color True if we want to use color in messages
*/
func SetUseColor(use_color bool) {
	if (use_color) {
		useColor = 2
	} else {
		useColor = 1
	}
}


/**
* Log an error
*/
func Error(message string) {
	print(ErrorLevel, "ERROR: " + message);
}

/**
* This and all of the other "LevelF()" functions are wrappers for Printf.
*/
func Errorf(message string, v ...interface{}) {
	printf(ErrorLevel, "ERROR: " + message, v...)
}


/**
* Log a warning
*/
func Warn(message string) {
	print(WarnLevel, "WARN: " + message)
}


/**
* This and all of the other "LevelF()" functions are wrappers for Printf.
*/
func Warnf(message string, v ...interface{}) {
	printf(WarnLevel, "WARN: " + message, v...)
}


/**
* Log info
*/
func Info(message string) {
	print(InfoLevel, "INFO: " + message)
}


/**
* This and all of the other "LevelF()" functions are wrappers for Printf.
*/
func Infof(message string, v ...interface{}) {
	printf(InfoLevel, "INFO: " + message, v...)
}


/**
* Log debugigng
*/
func Debug(message string) {
	print(DebugLevel, "DEBUG: " + message)
}


/**
* This and all of the other "LevelF()" functions are wrappers for Printf.
*/
func Debugf(message string, v ...interface{}) {
	printf(DebugLevel, "DEBUG: " + message, v...)
}


/**
* Log a trace 
*/
func Trace(message string) {
	print (TraceLevel, "TRACE: " + message)
}


/**
* This and all of the other "LevelF()" functions are wrappers for Printf.
*/
func Tracef(message string, v ...interface{}) {
	printf(TraceLevel, "TRACE: " + message, v...)
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

		//
		// Escape newlines and carriage returns
		//
		message = strings.Replace(message, "\n", "\\n", -1)
		message = strings.Replace(message, "\r", "\\r", -1)


		//
		// Escape anything else that is below ASCII 32 (space)
		// This is done because backspace (ASCII 8) is especially evil,
		// as an attrack could inject backspaces to hide their tracks
		//
		re, _ := regexp.Compile("([\x00-\x1f])")
		message = re.ReplaceAllStringFunc(message, func(in string) string {
			return(fmt.Sprintf("[0x%x]", in))
			})

		//
		// If displayTime is uninitialized, default to true
		//
		if (displayTime == 0 || displayTime == 2) {
			now := time.Now()
			elapsed := now.Sub(startTime)
			if (UseColor()) {
				color_string := getColor(level)
				color.Printf(color_string + "[%s] %s\n", elapsed, message)
			} else {
				fmt.Printf("[%s] %s\n", elapsed, message)
			}

		} else {
			if (UseColor()) {
				color_string := getColor(level)
				color.Println(color_string + message)
			} else {
				fmt.Println(message)
			}

		}

	}

} // End of print()


/**
* Get a color code for a specific log level.
*
* @param {integer} level The reporting level
* @return {string} The string for the color code to return.
*/
func getColor(level int) (retval string) {

	retval = ""
	if (level == ErrorLevel) {
		retval = "@{!r}"
	} else if (level == WarnLevel) {
		retval = "@{!y}"
	} else if (level == InfoLevel) {
		retval = "@{!g}"
	} else if (level == DebugLevel) {
		retval = "@{!c}"
	} else if (level == TraceLevel) {
		retval = "@{!b}"
	}

	return(retval)

} // End of getcolor()


/**
* Fancy wrapper for Printf.
*/
func printf(level int, message string, v ...interface{}) {
	message_out := fmt.Sprintf(message, v...)
	print(level, message_out)
}



