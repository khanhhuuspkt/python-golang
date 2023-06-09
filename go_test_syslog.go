package main

import (
    "log"
    "log/syslog"
)

func main() {
    syslogger, err := syslog.New(syslog.LOG_INFO, "syslog_example")
    if err != nil {
        log.Fatalln(err)
    }

    log.SetOutput(syslogger)
    log.Println("Log entry")


    syslogger1, err1 := syslog.New(syslog.LOG_CRIT, "syslog_example")
    if err1 != nil {
        log.Fatalln(err1)
    }
    // Set log out put and enjoy :)
    log.SetOutput(logFile)

    // optional: log date-time, filename, and line number
    log.SetFlags(log.Lshortfile | log.LstdFlags)

    log.SetOutput(syslogger1)
    log.Println("Log entry LOG_CRIT")
}
