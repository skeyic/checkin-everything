package service

import (
	"bytes"
	"checkin-everything/utils"
	"crypto/md5"
	"fmt"
	"github.com/golang/glog"
	"github.com/tidwall/gjson"
	"math/rand"
	"net/http"
	"time"
)

type Genshin struct {
	name   string
	cookie string
}

func newGenshin(name, cookie string) checkinSvc {
	return &Genshin{
		name:   name,
		cookie: cookie,
	}
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(97, 122))
	}
	return string(bytes)
}

func getDS() string {
	var (
		salt = "1OUn34iIy84ypu9cpXyun2VaQ2zuFeLm"
	)

	timestamp := time.Now().Unix()
	randomString := randomString(6)
	dsString := fmt.Sprintf("salt=%s&t=%d&r=%s", salt, timestamp, randomString)
	md5String := fmt.Sprintf("%x", md5.Sum([]byte(dsString)))
	glog.V(4).Infof("SALT: %s, TIMESTAMP: %d, RANDOM_STRING: %s, MD5: %s", salt, timestamp, randomString, md5String)
	return fmt.Sprintf("%d,%s,%s", timestamp, randomString, md5String)
}

func (s *Genshin) UidURL() string {
	return "https://api-takumi.mihoyo.com/binding/api/getUserGameRolesByCookie?game_biz=hk4e_cn"
}

func (s *Genshin) SignURL() string {
	return "https://api-takumi.mihoyo.com/event/bbs_sign_reward/sign"
}

func (s *Genshin) uidHeaders(req *http.Request) {
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Referer", "https://webstatic.mihoyo.com/bbs/event/signin-ys/index.html?bbs_auth_required=true&act_id=e202009291139501&utm_source=bbs&utm_medium=mys&utm_campaign=icon")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 14_0_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) miHoYoBBS/2.1.0")
	req.Header.Set("Cookie", s.cookie)
}

func (s *Genshin) signHeaders(req *http.Request) {
	s.uidHeaders(req)
	req.Header.Set("DS", getDS())
	req.Header.Set("x-rpc-device_id", "7ab3bc70b846186b9da1e816e6c6f08d")
	req.Header.Set("x-rpc-client_type", "4")
	req.Header.Set("x-rpc-app_version", "2.33.1")
}

func (s *Genshin) Checkin() (err error) {
	defer func() {
		if err != nil {
			utils.SendAlertV2("CheckIn Genshin Failed for "+s.name, "Error: "+err.Error())
		} else {
			utils.SendAlertV2("CheckIn Genshin Successfully for "+s.name, "")
		}
	}()
	rCode, rBody, rErr := utils.SendRequestCustom(http.MethodGet, s.UidURL(), nil, s.uidHeaders)
	if rErr != nil {
		glog.Errorf("GENSHINE Get Uid failed, code: %d, body: %s, error: %v", rCode, rBody, rErr)
		err = rErr
		return
	}

	uid := gjson.Get(rBody, "data.list.0.game_uid")
	data := fmt.Sprintf(`{"act_id": "e202009291139501", "region": "cn_gf01", "uid": %s}`, uid)
	rCode, rBody, rErr = utils.SendRequestCustom(http.MethodPost, s.SignURL(), bytes.NewBufferString(data), s.signHeaders)
	if rErr != nil {
		glog.Errorf("GENSHINE Checkin failed, code: %d, body: %s, error: %v", rCode, rBody, rErr)
		err = rErr
		return
	}
	glog.V(4).Infof("BODY: %s", rBody)

	return
}
