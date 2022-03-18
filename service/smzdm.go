package service

import (
	"checkin-everything/utils"
	"github.com/golang/glog"
	"net/http"
)

type Smzdm struct {
	name   string
	cookie string
}

func newSmzdm(name, cookie string) checkinSvc {
	return &Smzdm{
		name:   name,
		cookie: cookie,
	}
}

func (s *Smzdm) URL() string {
	return "https://zhiyou.smzdm.com/user/checkin/jsonp_checkin"
}

func (s *Smzdm) headers(req *http.Request) {
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Host", "zhiyou.smzdm.com")
	req.Header.Set("Referer", "https://www.smzdm.com/")
	req.Header.Set("Sec-Fetch-Dest", "script")
	req.Header.Set("Sec-Fetch-Mode", "no-cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36")
	req.Header.Set("Cookie", s.cookie)
}

func (s *Smzdm) Checkin() (err error) {
	defer func() {
		if err != nil {
			utils.SendAlertV2("CheckIn Failed for "+s.name, "Error: "+err.Error())
		} else {
			utils.SendAlertV2("CheckIn Successfully for "+s.name, "")
		}
	}()
	rCode, rBody, rErr := utils.SendRequestCustom(http.MethodPost, s.URL(), nil, s.headers)
	if rErr != nil {
		glog.Errorf("SMZDM CHECK IN failed, code: %d, body: %s, error: %v", rCode, rBody, rErr)
		err = rErr
		return
	}

	return
}
