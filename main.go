package main

import (
	"checkin-everything/service"
	"flag"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()

	service.Start()

	glog.V(4).Info("Check In Everything Service started")
	<-make(chan struct{}, 1)
}
