package concurrent

import (
	"fmt"
	"runtime"
	"testing"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func TestCoutine(t *testing.T) {
	go say("world") //开一个新的Goroutines执行
	say("hello") //当前Goroutines执行
}
