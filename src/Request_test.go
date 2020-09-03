package src

import (
	"fmt"
	"github.com/Jecced/rs/src/rs"
	"testing"
)

func TestGetReadText(t *testing.T) {
	resp := rs.Get("http://www.baidu.com/").Send().ReadText()
	fmt.Println(resp)
}

func TestGetWriteToFile(t *testing.T) {
	rs.Get("http://www.baidu.com/img/sug_bd.png?v=09816787.png").
		Send().
		WriteToFile("/Users/ankang/develop/test/test/1.png")
}

func TestGetProxy(t *testing.T) {
	resp := rs.Get("http://www.google.com").Proxy("127.0.0.1:1081").Send().ReadText()
	fmt.Println(resp)
}
