package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"strconv"
	"time"
)

// MD5哈希
func Md5(data string) string {
	t := md5.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
	//data := []byte(d)
	//has := md5.Sum(data)
	//return fmt.Sprintf("%x", has)
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

// 获取URL中提供的参数
func ParseURLParams(ctx *gin.Context) (account_id, device_id, service_id, id int64, err1, err2, err3, err4 error) {
	account_id, err1 = strconv.ParseInt(ctx.Param("account"), 10, 64)
	device_id, err2 = strconv.ParseInt(ctx.Param("device"), 10, 64)
	service_id, err3 = strconv.ParseInt(ctx.Param("service"), 10, 64)
	id, err4 = strconv.ParseInt(ctx.Param("id"), 10, 64)
	return account_id, device_id, service_id, id, nil, nil, nil, nil
}
