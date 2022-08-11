package service

import (
	"checkin-everything/utils"
	"crypto/md5"
	"fmt"
	"github.com/golang/glog"
	"github.com/tidwall/gjson"
	"net/http"
	"strings"
	"testing"
)

var cookies = "_MHYUUID=28088536-c698-46e5-85d5-07719bf9b1c0; _ga=GA1.2.888577559.1630052488; UM_distinctid=181cd284905886-0ecb184677133e-1b3c2165-1fa400-181cd284906985; _gid=GA1.2.20698955.1659967711; CNZZDATA1275023096=79157469-1657002140-https%253A%252F%252Fys.mihoyo.com%252F%7C1659965998; smidV2=20220808220851d145f63135cec1fd964658821cf71af7007531810c97bdb50; .thumbcache_a5f2da7236017eb7e922ea0d742741d5=rqn4irEr1ivsZuziop2qOieoJhlh6GOQhghPe673rc6rFm36TzmVIeSHmfyVv8A/yl7pgiqu21Oq5oFX9SKUag%3D%3D; ltoken=ghNexwFOzb7NqO9gZODu9lBtbmF6g1MUTv5TGJwt; ltuid=277541725; cookie_token=MExHfHiUzHneroxz9hOm1zjmWFTeOeCzedzqmmq3; account_id=277541725#_MHYUUID=3d4b85d9-60fe-49d2-8790-a80ec5a2cd4f; _ga=GA1.2.713370200.1659967403; _gid=GA1.2.488516915.1659967403; UM_distinctid=1827dc4a9d25ab-0005463b703122-26021a51-384000-1827dc4a9d3768; smidV2=2022080822033706718dd76b4174363b6e731eb6cd05d50017c6ae4ef537510; .thumbcache_a5f2da7236017eb7e922ea0d742741d5=iz6uJh5v5vP/neSYKNbroyaJFqMeVzyZ8nsti/8GmF5L2fras8vYRvw/uY7a55jMA83bsxRU8SUskW0RfsOlSQ%3D%3D; CNZZDATA1275023096=1135566375-1659965998-https%253A%252F%252Fbbs.mihoyo.com%252F%7C1659965998; ltoken=6aPfKkQy4pFV55O1q8v8dICWcqQQXtrhBCzHnqGQ; ltuid=337236739; cookie_token=ULhnwGeR5ISkYFQc7eSvpqRP9UDDH2nG8f9KDmQM; account_id=337236739; _gat=1#_MHYUUID=f980dc79-acec-4a09-9bf8-496067a25af8; mi18nLang=zh-cn; _ga_XR5VD06Z8Y=GS1.1.1644071419.8.1.1644071604.0; _ga_Q3LKDGYS1J=GS1.1.1644212727.1.1.1644212806.0; _ga_6ZKT513CTT=GS1.1.1647996153.1.1.1647996240.0; _ga_HKTGWLY8PN=GS1.1.1648351199.2.1.1648351336.0; _ga_R8CG4VZ69C=GS1.1.1650864331.6.1.1650864379.0; ltuid=277345273; _ga_KJ6J9V9VZQ=GS1.1.1651472150.1.0.1651472150.0; _ga_55SMHPM22L=GS1.1.1651507230.3.1.1651507411.0; _ga_9TTX3TE5YL=GS1.1.1657071345.1.0.1657071345.0; _ga_N90YRBX6FW=GS1.1.1657295268.5.1.1657295446.0; UM_distinctid=181f26bf89e5d3-0c1770943e092d-26021b51-384000-181f26bf89f5d7; _ga=GA1.2.965744188.1641718264; _ga_VNK94KKWTF=GS1.1.1658986579.5.1.1658986620.0; _gid=GA1.2.751365916.1659885425; smidV2=2022080723170508e291f29cac5f8104b1061dc9978ca8009f9d02358698ad0; .thumbcache_a5f2da7236017eb7e922ea0d742741d5=Cor5SAmGYawGSXFvh8SukAtGXyWEDhNoWIcuY9RYFK8lfUwFrW7knIs+YHKK49FeaTiSyf0BE3c/2T51SQd9cg%3D%3D; CNZZDATA1275023096=1177002171-1641716930-https%253A%252F%252Fwww.baidu.com%252F%7C1659965998; _gat=1; ltoken=6EdqT9qvYJJq3qHhln37sgkVch2dPhxpXwawYshH; cookie_token=vvoZRA99JEnmLU6v4ZATju73cbrsQwTQFrPS7DGX; account_id=277345273; _dd_s=logs=1&id=caa24d25-5280-4ec8-a99d-d8c4157934b7&created=1659968122389&expire=1659969045384'"
var cookie1 = strings.Split(cookies, "#")[0]

func TestGenshin_Checkin(t *testing.T) {
	utils.EnableGlogForTesting()

	svc := newGenshin("genshin", cookie1)
	svc.Checkin()
}

func Test_get_uid(t *testing.T) {
	utils.EnableGlogForTesting()

	s := newGenshin("genshin", cookie1).(*Genshin)

	rCode, rBody, rErr := utils.SendRequestCustom(http.MethodGet, s.UidURL(), nil, s.uidHeaders)
	if rErr != nil {
		glog.Errorf("GENSHINE Get Uid failed, code: %d, body: %s, error: %v", rCode, rBody, rErr)
		return
	}

	glog.V(4).Infof("RBODY: %s", rBody)

	uid := gjson.Get(rBody, "data.list.0.game_uid")
	glog.V(4).Infof("UID: %s", uid)
}

func Test_get_uid_from_body(t *testing.T) {
	utils.EnableGlogForTesting()
	var (
		rBody = `{"retcode":0,"message":"OK","data":{"list":[{"game_biz":"hk4e_cn","region":"cn_gf01","game_uid":"175258477","nickname":"夏一跳","level":58,"is_chosen":true,"region_name":"天空岛","is_official":true}]}}`
	)

	uid := gjson.Get(rBody, "data.list.0.game_uid")
	glog.V(4).Infof("UID: %s", uid)
}

func Test_GetSalt(t *testing.T) {
	utils.EnableGlogForTesting()
	var (
		salt = "1OUn34iIy84ypu9cpXyun2VaQ2zuFeLm"
	)

	//timestamp := time.Now().Unix()
	//randomString := randomString(6)
	timestamp := 1660200119
	randomString := "floqyh"
	dsString := fmt.Sprintf("salt=%s&t=%d&r=%s", salt, timestamp, randomString)
	md5String := fmt.Sprintf("%x", md5.Sum([]byte(dsString)))
	glog.V(4).Infof("SALT: %s, TIMESTAMP: %d, RANDOM_STRING: %s, MD5: %s", salt, timestamp, randomString, md5String)
}
