package filesystem

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
)

var (
	hash     map[string]string
	hashName string
)

func getExtension() {

}

// 自动生成文件名
func generateHashName(rule string) string {
	if hashName == "" {
		timeUnix := time.Now().Unix()
		today := time.Now().Format("2006-01-02")

		md5Name := fmt.Sprintf("%x", md5.Sum([]byte(strconv.FormatInt(timeUnix, 10))))

		hashName = today + "/" + md5Name
	}

	return hashName
}
