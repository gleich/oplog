package main

import "go.mattglei.ch/tlog"

var Cache = tlog.Group("cache", &struct {
	MarshalResponse tlog.Task
	Marshal         struct {
		JSON tlog.Task
		CSV  tlog.Task
	}
}{})

func main() {
	Cache.MarshalResponse.Info("hello world!")
}
