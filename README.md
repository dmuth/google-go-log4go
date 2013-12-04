## Introducing Log4go

Still being new to Google Go, I wanted a logging facility.  I didn't 
have much luck finding one, so I decided to write a quick and dirty
implementation of Log4j.

### Installation

Installation is easy. First make sure you have your [GOPATH](http://golang.org/doc/code.html#GOPATH) setup properly. Then use `go get` to download and compile this package.

    go get github.com/dmuth/google-go-log4go

Also, make sure your $GOPATH environment variable is set up properly: `export GOPATH=$HOME/golib` or similar

### Usage

    import log "github.com/dmuth/google-go-log4go"
    
    func main() {
        log.SetLevel(log.DebugLevel)
        log.SetDisplayTime(true)
        log.Info("foobar")
        
        //
        // Strings work, too.
        //
        log.SetLevelString("info")
        log.Info("baz")
        
        //
        // Strings are even case-insensitive!
        //
        log.SetLevelString("INFO")
        log.Info("blurfl")
        
        //
        // I herd u liek Printf()
        //
        log.Infof("Number of widgets: %d", 1001)
        
    }

And you can expect to see output like this:

![log4go screenshot](https://raw.github.com/dmuth/google-go-log4go/master/docs/log4go_screenshot.png)

(Colored text is created with to Meng Zhang's excellent Terminal package: https://github.com/wsxiaoys/terminal)

### Escaping malicious input

It was pointed out to me that if things like invalid logins, SQL 
injection attempts, etc. are logged, an would-be attacker could pass in
the backspace character (ASCII 8) to "hide" output on a casual glance.

This package tries to protect against that by escaping anything with an
ASCII value of less than 32.  This includes carriage returns (ASCII 13)
and newlines (ASCII 10).

Here is an example of an attacker trying to use backspaces, and what 
this module would print out instead:

    ERROR: Failed login: BlackHat[0x08][0x08][0x08][0x08][0x08][0x08][0x08][0x08]WhiteHat


### Testing

`go test ./src/github.com/dmuth/google-go-log4go/`

You should see output like this:

`ok      github.com/dmuth/google-go-log4go       0.009s`

### Contact Me

This is my first ever Google Go package.  Chances are that I've done like over 9,000 things the wrong way.  Please drop me a line if you liked the package, thought it sucked, want to discuss the merits of velociraptors as housepets, or whatever.  

My email address is doug.muth@gmail.com.  Additional contact methods can be found at http://www.dmuth.org/contact  Thanks!

