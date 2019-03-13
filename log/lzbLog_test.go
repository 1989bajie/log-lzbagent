package log

import (
	"testing"
)

func TestError(t *testing.T) {
	//conf.InitConf()
	//t.Error(conf.GetConf())
	//var c Conf
	Error("ceshi", 1)
}

/*func BenchmarkGetConf(b *testing.B) {
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
*/
