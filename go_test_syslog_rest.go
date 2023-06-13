package main

import (
    "bufio"
    "fmt"
    "net/http"
    "log"
    "log/syslog"
)

func main() {

    resp, err := http.Get("https://official-joke-api.appspot.com/random_joke")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("Response status:", resp.Status)
    syslogger, err := syslog.New(syslog.LOG_INFO, "syslog_example")
    if err != nil {
        log.Fatalln(err)
    }

    log.SetOutput(syslogger)
    log.Println("Response status:" + resp.Status)
    //fmt.Println(resp.Body);
    scanner := bufio.NewScanner(resp.Body)
    for i := 0; scanner.Scan() && i < 5; i++ {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

}