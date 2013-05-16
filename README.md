Introducing Log4go
==================

Still being new to Google Go, I wanted a logging facility.  I didn't 
have much luck finding one, so I decided to write a quick and dirty
implementation of Log4j.

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

