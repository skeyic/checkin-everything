package service

import (
	"checkin-everything/config"
	"github.com/golang/glog"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"sync"
	"time"
)

var (
	TheMaster  = newMaster()
	configPath = config.Config.DataFolder + "/config.json"
)

func Start() {
	//for _, smzdmCookie := range SmzdmCookies {
	//	TheMaster.Add(newSmzdm(smzdmCookie["name"], smzdmCookie["cookie"]))
	//}
	TheMaster.Start()
}

type checkinSvc interface {
	Checkin() error
}

type master struct {
	lock *sync.RWMutex
	svcs []checkinSvc
}

func newMaster() *master {
	return &master{
		lock: &sync.RWMutex{},
	}
}

func (m *master) Add(svc checkinSvc) {
	m.lock.Lock()
	m.svcs = append(m.svcs, svc)
	m.lock.Unlock()
}

func (m *master) Load() {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(any(err))
	}

	for _, cookie := range gjson.Get(string(data), "data.smzdm").Array() {
		var (
			name   = gjson.Get(cookie.String(), "name").String()
			cookie = gjson.Get(cookie.String(), "cookie").String()
		)
		m.Add(newSmzdm(name, cookie))
		glog.V(4).Infof("Add name: %s", name)
	}
}

func (m *master) Start() {
	m.Load()
	go m.process()
}

func (m *master) process() {
	ticker := time.NewTicker(12 * time.Hour)
	startChan := make(chan time.Time, 1)
	go func() {
		time.Sleep(30 * time.Second)
		startChan <- time.Now()
	}()

	doFunc := func(a time.Time) {
		glog.V(4).Infof("Do Func at %s", a)
		for _, svc := range m.svcs {
			go func(svc checkinSvc) {
				svc.Checkin()
			}(svc)
		}
	}

	for {
		select {
		case a := <-ticker.C:
			doFunc(a)
		case a := <-startChan:
			doFunc(a)
		}
	}
}
