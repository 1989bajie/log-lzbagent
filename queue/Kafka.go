package queue
import(
    "time"
    "strconv"
    "encoding/json"
    "log-lzbagent/conf"
    "log-lzbagent/pool"
    "log-lzbagent/log"
    "github.com/Shopify/sarama"
)
var KafkaStatus bool = false 
//kafka生产
func KafkaProducer(queueData *map[string]interface{}, channel chan bool){
    //无论报错与否,清除channel缓存
    if channel != nil {
        defer func(){<- channel}()
    }
    data,err := json.Marshal(queueData)
    if !KafkaStatus {
        //写入本地文件
        FileProducer(string(data))
        return
    }
    //获取对象
    p := pool.GetKafkaProducer()
    if p.Producer == nil {
        for {
            time.Sleep(300 * time.Millisecond)
            p = pool.GetKafkaProducer()
            if(p.Producer != nil){
                break
            }
        }
    }
    msg := &sarama.ProducerMessage{
        Topic:conf.GetConf().Kafka.Topic,
        Value:sarama.ByteEncoder(data),
    }
    part, offset, err := p.Producer.SendMessage(msg)
    //归还对象
    pool.PutKafkaProducer(p)
    if err != nil {
        log.Error("send kafka message err:" + err.Error())
        log.Error("part:" + strconv.Itoa(int(part)))
        log.Error("offset:" + strconv.FormatInt(offset,10))
        //发送kafka失败，写入本地文件
        FileProducer(string(data))
    }
}