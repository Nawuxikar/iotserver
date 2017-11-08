package utils

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

func StrToMD5v1(s string) string{
	md5ctx := md5.New()
	md5ctx.Write([]byte(s))
	cipher :=md5ctx.Sum([]byte("wanglewo"))
	return hex.EncodeToString(cipher)
}

func StrToMD5v2(s string) string{
	md5ctx := md5.New()
	md5ctx.Write([]byte(s))
	cipher :=md5ctx.Sum([]byte("A"))
	return hex.EncodeToString(cipher)
}

func StrToMD5v3(s string) string{
	md5ctx := md5.New()
	md5ctx.Write([]byte(s))
	cipher :=md5ctx.Sum([]byte("Device"))
	return hex.EncodeToString(cipher)
}



func GetDateTime() string{

	return time.Now().String()
}