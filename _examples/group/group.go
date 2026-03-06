package main

import "go.mattglei.ch/oplog"

var Cache = oplog.Group("cache", &struct {
	MarshalResponse oplog.Op
	Marshal         struct {
		JSON oplog.Op
		CSV  oplog.Op
	}
}{})

func main() {
	Cache.MarshalResponse.Info("hello world!")
}
