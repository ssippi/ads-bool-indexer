package tools

import (
	"crypto/md5"
	"encoding/json"
)

func GetMd5(data interface{}) string {
	bytes, _ := json.Marshal(data)
	res := md5.Sum(bytes)
	md5 := string(res[:])
	return md5
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
