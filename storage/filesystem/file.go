package filesystem

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

//var (
//	hash     map[string]string
//	hashName string
//)

//func getExtension() {
//
//}

// 自动生成文件名
func generateHashName(rule string) string {
	if rule == "" {
		nowTime := time.Now()
		today := nowTime.Format("20060102")
		timeUnixNano := nowTime.UnixNano()
		m := md5.Sum([]byte(strconv.FormatInt(timeUnixNano, 10)))
		md5Name := hex.EncodeToString(m[:])

		return today + "/" + md5Name
	}

	return rule
}
