package main

import (
	"github.com/garyburd/redigo/redis"
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

	conn , err := redis.DialTimeout("tcp", "127.0.0.1:6379", 0, 1*time.Second, 1*time.Second)
	if err != nil {
		panic(err)
	}
	size ,err:= conn.Do("DBSIZE")
	fmt.Printf("size is %d \n",size)

	_,err = conn.Do("SET","user:user0",123)
	_,err = conn.Do("SET","user:user1",456)
	_,err = conn.Do("APPEND","user:user0",87)

	user0,err := redis.Int(conn.Do("GET","user:user0"))
	user1,err := redis.Int(conn.Do("GET","user:user1"))

	fmt.Printf("user0 is %d , user1 is %d \n",user0,user1)

	_,err = conn.Do("HSET","user_trace","A","1")
	_,err = conn.Do("HSET","user_trace","B","2")
	hlen ,err:= conn.Do("HLEN","user_trace")
	fmt.Println("user_trace size is %d",hlen)

	conn.Close()
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

