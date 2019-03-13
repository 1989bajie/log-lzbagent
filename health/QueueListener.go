package health
import(
	"net"
	"time"
	"log-lzbagent/conf"
	"log-lzbagent/pool"
	"log-lzbagent/queue"
)
var queueList = []string{"kafka","localfile"}
func HealthRun(){
	for _,queue := range queueList {
		switch queue {
		case "kafka":
			go KafkaListener()
			break
		case "localfile":
			go LocalFileListener()
			break
		}
	}
}
//kafka健康检查
func KafkaListener(){
	timer := time.NewTimer(30 * time.Second)
	addrArray := conf.GetConf().Kafka.Addr
    for {
        select {
		case <-timer.C:
			statusFlag := 0
			for i:=0; i<len(addrArray); i++ {
				p, err := net.DialTimeout("tcp", addrArray[i], 3 * time.Second)
				if err != nil {
					statusFlag++
					pool.RemoveInvalidProducer(addrArray[i])
				}else {
					p.Close()
					pool.AddInvalidProducer(addrArray[i])
				}
			}
			//集群全部挂掉，写本地文件
			if statusFlag == len(addrArray) {
				queue.KafkaStatus = false
			}else {
				queue.KafkaStatus = true
			}
            timer.Reset(time.Second * 30)
        }
	}
}