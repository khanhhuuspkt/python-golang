package main

import (
    "log"
//"log/syslog"
    "os"
)

func main() {
    // log to custom file
    LOG_FILE := "/tmp/myapp_log"
    // open log file
    logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        log.Panic(err)
    }
    defer logFile.Close()

    // Set log out put and enjoy :)
    log.SetOutput(logFile)

    // optional: log date-time, filename, and line number
    log.SetFlags(log.Lshortfile | log.LstdFlags)

    log.Println("hk hk Logging to custom file")
}
