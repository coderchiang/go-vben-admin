package utils

import (
	"time"
)

func TimeParseStr(t time.Time, format string) (res string ){
	switch format {
	case "FULL" :
		res=t.Format("2006-01-02 15:04:05")
	case  "YMD":
      res=t.Format("2006/01/02")
	case  "HMS":
      res=t.Format("15:04:05")
	}
	return
}


func StrParseTime(t string, format string) (ret time.Time ) {
	var res string
	switch format {
	case "FULL":
		res = "2006-01-02 15:04:05"
	case "YMD":
		res = "2006/01/02"
	case "HMS":
		res = "15:04:05"
	}
	ret, _ = time.ParseInLocation(res, t, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	return
}

