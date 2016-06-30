package main
import (
	"time"
	"fmt"
	"runtime"
)
func main() {
	go say("world")
	say("hello")/*
	for i:=1; i<=10000; i++ {
		fmt.Println(i)
		time.Sleep(time.Second / 1000)
	}
	fmt.Println("time out")*/
}
func say(s string) {
for i := 0; i < 5; i++ {
runtime.Gosched()
fmt.Println(s)
	time.Sleep(time.Second)
}
}