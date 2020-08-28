package rs

import (
	"fmt"
	"testing"
)

func Test01(t *testing.T) {
	resp := Get("http://www.baidu.com/img/sug_bd.png?v=09816787.png").Send().ReadText()
	fmt.Println(resp)
}

func Test02(t *testing.T) {
	Get("http://www.baidu.com/img/sug_bd.png?v=09816787.png").
		Send().
		WriteToFile("/Users/ankang/develop/test/test/1.png")
}

func Test03(t *testing.T) {
	resp := Get("http://www.google.com").Proxy("127.0.0.1:1081").Send().ReadText()
	fmt.Println(resp)
}
