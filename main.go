package main

import (
	"fmt"
	"time"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
)
func main()  {
	startCac()
	fmt.Println(MD5("asdfgvf"))
	fmt.Println(GetRandomString(4))
}
//生成32位小写MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

//生成随机字符串
func GetRandomString(lenth int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0;i < lenth ;i++  {
		result = append(result,bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//计算该方法运行的时间
func startCac()  {
	t1 := time.Now()	//获取系统当前时间
	for i:=0;i<1000 ;i++  {
		fmt.Print("*")
	}
	elapsed := time.Since(t1)
	fmt.Println("")
	fmt.Println("App elapsed:",elapsed)
}

