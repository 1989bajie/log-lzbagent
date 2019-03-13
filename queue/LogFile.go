package queue
import(
	"time"
	"io/ioutil"
	"strings"
	"encoding/json"
	"log-lzbagent/pool"
	"log-lzbagent/log"
)
//本地文件存储日志数据
func FileProducer(queueData string){
	f := pool.GetFileProducer()
	//归还对象
    defer pool.PutFileProducer(f)
	if(f == nil){
        for {
            time.Sleep(300 * time.Millisecond)
            f = pool.GetFileProducer()
            if(f != nil){
                break
            }
        }
    }
	result,err := f.WriteString(queueData + "\n")
	if(err != nil){
		log.Error("log write file error:" + err.Error())
	}
	if(result == 0){
		log.Error("log write file size 0,file name:"+f.Name())
	}
}
//消费本地文件日志数据
func FileConsumer() error{
	f := pool.GetFileProducer()
	//归还对象
	defer pool.PutFileProducer(f)
	if(f == nil){
        for {
            time.Sleep(300 * time.Millisecond)
            f = pool.GetFileProducer()
            if(f != nil){
                break
            }
        }
    }
	logData,err := ioutil.ReadFile(f.Name())
	if err != nil {
		log.Error("local file consumer read err:" + err.Error())
		return err
	}
	array := strings.Split(string(logData),"\n")
	for _,logString := range array {
		if logString != "" {
			var logMap *map[string]interface{}
			json.Unmarshal([]byte(logString),&logMap)
			//重新扔进kafka进行消费
			KafkaProducer(logMap,nil)
		}
	}
	//消费完清空文件内容
	f.Truncate(0)
	return nil
}