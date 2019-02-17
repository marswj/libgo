package libgo

import "time"

//unix时间戳格式化
func TimeStr(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

//unix时间戳格式化
func TimeStr2(t int64) string {
	t1 := time.Unix(t, 0)
	return t1.Format("2006-01-02 15:04:05")
}

//unix时间戳
func TimeUnix() int64 {
	return time.Now().Unix()
}
