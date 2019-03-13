package health
import(
	"time"
	"log-lzbagent/queue"
	"log-lzbagent/conf"
)
func LocalFileListener(){
	timer := time.NewTimer(30 * time.Second)
    for {
        select {
		case <-timer.C:
			if(queue.KafkaStatus){
				//kafka正常
				var i int
				for i = 0; i< conf.GetConf().Kafka.PoolSize; i++ {
					queue.FileConsumer()
				}
			}
            timer.Reset(30 * time.Second)
        }
    }
}