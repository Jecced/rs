package rs

import (
	"fmt"
	"testing"
)

func TestGetReadText(t *testing.T) {
	resp := Get("http://www.baidu.com/").Send().ReadText()
	fmt.Println(resp)
}

func TestGetWriteToFile(t *testing.T) {
	Get("http://www.baidu.com/img/sug_bd.png?v=09816787.png").
		Send().
		WriteToFile("/Users/ankang/develop/test/test/1.png")
}

func TestGetProxy(t *testing.T) {
	resp := Get("http://www.google.com").Proxy("127.0.0.1:1081").Send().ReadText()
	fmt.Println(resp)
}
