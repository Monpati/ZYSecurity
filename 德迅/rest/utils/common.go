package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func HashSalt(str, salt string) (string, error) {
	m := md5.New()
	io.WriteString(m, str)
	m.Sum(nil)
	io.WriteString(m, salt)
	return hex.EncodeToString(m.Sum(nil)), nil
}

// MD5哈希
func Md5(data string) string {
	t := md5.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
	//data := []byte(d)
	//has := md5.Sum(data)
	//return fmt.Sprintf("%x", has)
}

func MD5(v interface{}) string {
	a := v.(string)
	d := []byte(a)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

// 获取随机数 纯文字
func GetRandomString(n int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
