package service

import (
	"checkin-everything/utils"
	"testing"
)

func TestSmzdm_Checkin(t *testing.T) {
	utils.EnableGlogForTesting()

	svc := newSmzdm("smzdm1", "device_id=2130706433159836439057149409a412b512c0b286a4128d5af5a6b7a4; r_sort_type=score; userId=user:4954308188|4954308188; smzdm_user_source=25645266EA1C5F488FF13DF6551AA30A; __ckguid=fcA4he3CqIfT1FSVLiqbCR5; __jsluid_s=148ff24cca57da7cea475ddbf82a28a0; _ga=GA1.2.1285720849.1598364393; __gads=ID=080855f848dcf65f:T=1623762753:S=ALNI_MZ2KAI40fTQIbaGg2VCXKwAyO-RJQ; _ga_09SRZM2FDD=GS1.1.1629908399.5.0.1629908409.0; homepage_sug=c; sess=AT-m3gzB7NH3J8dQ4twoYfXnTIPWYKcBdVYLhlsUTGhLSWcYIvkWQUaGHMPsm%2BUmF9iNMItCjBW%2B0Qa05LXLBo%2FNB6IP%2BB1zApXkN0OfV4owRevhYdoiQCpJsup; user=user%3A4954308188%7C4954308188; smzdm_id=4954308188; _zdmA.uid=ZDMA.UPLgkyDw6.1648126065.2419200; _zdmA.vid=*; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%224954308188%22%2C%22first_id%22%3A%2217425f119fbc71-0078e7f9f3e085-33504770-3686400-17425f119fcb50%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%2C%22%24latest_referrer%22%3A%22%22%2C%22%24latest_landing_page%22%3A%22https%3A%2F%2Fwww.smzdm.com%2F%22%7D%2C%22%24device_id%22%3A%2217425f119fbc71-0078e7f9f3e085-33504770-3686400-17425f119fcb50%22%7D; Hm_lvt_9b7ac3d38f30fe89ff0b8a0546904e58=1646288320,1647506888,1648126065; Hm_lpvt_9b7ac3d38f30fe89ff0b8a0546904e58=1648126065; footer_floating_layer=0; ad_date=24; bannerCounter=%5B%7B%22number%22%3A0%2C%22surplus%22%3A1%7D%2C%7B%22number%22%3A0%2C%22surplus%22%3A1%7D%2C%7B%22number%22%3A0%2C%22surplus%22%3A1%7D%2C%7B%22number%22%3A0%2C%22surplus%22%3A1%7D%2C%7B%22number%22%3A0%2C%22surplus%22%3A1%7D%2C%7B%22number%22%3A0%2C%22surplus%22%3A1%7D%5D; ad_json_feed=%7B%7D; amvid=69fa7bfd20a48d43e32c0db9c965ca5e; _zdmA.time=1648126067079.0.https%3A%2F%2Fwww.smzdm.com%2F")
	svc.Checkin()
}
