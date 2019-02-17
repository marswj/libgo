package libgo

import (
	"github.com/satori/go.uuid"
	"strconv"
	"strings"
)

func StrToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

func StrToFloat64(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

//去除两端空白
func StrTrim(data string) string {
	return strings.TrimSpace(data)
}

//uuid len=36
func UUIDStr() string {
	u1, _ := uuid.NewV1()
	return u1.String()
}

//uuid len=32 无"-"
func UUIDStr2() string {
	u1, _ := uuid.NewV1()
	return strings.Replace(u1.String(), "-", "", -1)
}