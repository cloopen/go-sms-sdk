package cloopen

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Base64(str string) string {
	strbytes := []byte(str)
	return base64.StdEncoding.EncodeToString(strbytes)
}
func Base64URL(str string) string {
	strbytes := []byte(str)
	return base64.URLEncoding.EncodeToString(strbytes)
}
