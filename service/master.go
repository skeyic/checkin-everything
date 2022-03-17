package service

import (
	"checkin-everything/config"
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
	for _, smzdmCookie := range SmzdmCookies {
		TheMaster.Add(newSmzdm(smzdmCookie))
	}
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
		m.Add(newSmzdm(cookie.String()))
	}
}

func (m *master) Start() {
	m.Load()
	go m.process()
}

func (m *master) process() {
	ticker := time.NewTimer(6 * time.Hour)
	for {
		select {
		case <-ticker.C:
			for _, svc := range m.svcs {
				go svc.Checkin()
			}
		}
	}
}
