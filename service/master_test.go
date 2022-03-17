package service

import (
	"checkin-everything/utils"
	"github.com/golang/glog"
"github.com/tidwall/gjson"
	"io/ioutil"
	"testing"
)

func Test_master_Load(t *testing.T) {
	utils.EnableGlogForTesting()
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(any(err))
	}

	for _, cookie := range gjson.Get(string(data), "data.smzdm").Array() {
		glog.V(4).Infof("BODY: %s", cookie)
	}
}
