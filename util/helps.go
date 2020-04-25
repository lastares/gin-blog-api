package util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func MD5(key string) string {

	for i := 0; i < 2; i++ {
		hash := md5.New()
		hash.Write([]byte(key))
		key = strings.ToUpper(hex.EncodeToString(hash.Sum(nil)))
	}
	return key
}
