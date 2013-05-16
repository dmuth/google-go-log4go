Introducing Log4go
==================

Still being new to Google Go, I wanted a logging facility.  I didn't 
have much luck finding one, so I decided to write a quick and dirty
implementation of Log4j.

Installation
============

For Beginners
-------------
    GOPATH=`pwd` go get   

For Nerds
---------
    go get
    
The latter example assumes you already have `$GOPATH` set.  `$HOME/golibs/` isn't a bad place.

Usage
=====

    import log "github.com/dmuth/google-go-log4go"
    
    func main() {
        log.SetLevel(log.DebugLevel)
        log.SetDisplayTime(true)
        log.Info("foobar")
    }

And you can expect to see output like this:

    [301.119ms] INFO: foobar

Testing
========

`go test ./src/github.com/dmuth/google-go-log4go/`

You should see output like this:

`ok      github.com/dmuth/google-go-log4go       0.009s`

Contact Me
==========

This is my first ever Google Go package.  Chances are that I've done like over 9,000 things the wrong way.  Please drop me a line if you liked the package, thought it sucked, want to discuss the merits of velociraptors as housepets, or whatever.  

My email address is doug.muth@gmail.com.  Additional contact methods can be found at http://www.dmuth.org/contact  Thanks!

