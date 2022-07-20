package tools

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func MobileEncrypt(mobile string) string {
	if mobile == "" || len(mobile) != 11 {
		return mobile
	}
	var n strings.Builder
	n.WriteString(mobile[:3])
	n.WriteString("****")
	n.WriteString(mobile[7:])
	return n.String()
}
