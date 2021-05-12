package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(data string) string  {
	obj := sha256.Sum256([]byte(data))
	return hex.EncodeToString(obj[:])
}

func Sha1(data string) string {
	obj := sha1.New()
	obj.Write([]byte(data))
	return hex.EncodeToString(obj.Sum([]byte("")))
}

func Hmac(key, data string) string {
	obj := hmac.New(md5.New, []byte(key))
	obj.Write([]byte(data))
	return hex.EncodeToString(obj.Sum([]byte("")))
}

func Md5(data string) string {
	obj := md5.New()
	obj.Write([]byte(data))
	return hex.EncodeToString(obj.Sum([]byte("")))
}
