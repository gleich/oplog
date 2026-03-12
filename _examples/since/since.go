package main

import (
	"time"

	"go.mattglei.ch/tlog"
)

func main() {
	start := time.Now()
	time.Sleep(2 * time.Second)
	tlog.Task("waiting-2-seconds").InfoSince("waited 2 seconds!", start)
}
