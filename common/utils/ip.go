package utils

import (
	"encoding/json"
	"gin-vben-admin/common"
	"io/ioutil"
	"net/http"
)

func GetLocation(ip string) interface{} {
	if ip == "127.0.0.1" || ip == "localhost"||ip=="::1" {
		return "内部IP"
	}
	resp, err := http.Get(common.CONFIG.BaiduMap.IpLocationUrl + common.CONFIG.BaiduMap.AK + "&ip="+ip+"&coor=bd09ll")
	if err != nil {
		panic(err)

	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)

	m := make(map[string]interface{})

	err = json.Unmarshal(s, &m)
	if err != nil {
		return "未知位置"
	}
	if m["address"] == "" ||m["address"]==nil{
		return "未知位置"
	}
	return m["address"]
}