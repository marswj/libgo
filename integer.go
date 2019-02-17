package libgo

import "strconv"

func IntToStr(i int) string {
	return strconv.Itoa(i)
}

func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}
