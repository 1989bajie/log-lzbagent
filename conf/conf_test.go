package conf

import (
	"testing"
)

func TestGetConf(t *testing.T) {
	//var c Conf
	InitConf()
	conf := GetConf()
	t.Error(conf)
	if conf.Local.Port != 65065 {
		t.Error("get conf error")
	}
}

func BenchmarkGetConf(b *testing.B) {
	InitConf()
	for i := 0; i < b.N; i++ {
		//a := Test
		//print(a)
		conf := GetConf()
		if conf.Local.Port != 65065 {
			print(111)
		}
	}
}
