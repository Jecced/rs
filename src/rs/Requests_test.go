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
