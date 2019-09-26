package libgo

import (
	"os"
	"strconv"
	"strings"

	"github.com/gofrs/uuid"
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

//文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
