package pool
import(
	"time"
	"sync"
	"strconv"
	"math/rand"
	"container/list"
	"log-lzbagent/conf"
	"log-lzbagent/log"
	"github.com/Shopify/sarama"
)
var KafkaProducer = make(map[string]*list.List)
var InvalidProducer = make(map[string]*list.List)
var lock sync.Mutex
type KafkaObject struct {
	Key string
	Producer *list.List
}
type ProducerObject struct {
	Key string
	Producer sarama.SyncProducer
}
//连接池初始化
func Init() {
	config := sarama.NewConfig()
    config.Producer.Return.Successes = true
	config.Producer.Timeout = 3 * time.Second
	config.Net.DialTimeout = 5 * time.Second
	addrArray := conf.GetConf().Kafka.Addr
	var i int
	for i = 0; i<len(addrArray); i++ {
		var j int
		var KafkaList = list.New()
		for j = 0; j<conf.GetConf().Kafka.PoolSize; j++ {
			p, err := sarama.NewSyncProducer([]string{addrArray[i]}, config)
			if err != nil {
				panic(err)
			}
			log.Info("第" + strconv.Itoa(j) + "个kafkaProducer对象生成")
			KafkaList.PushBack(p)
		}
		KafkaProducer[addrArray[i]] = KafkaList
	}
	//本地文件初始化
	LocalFileInit()
}
//获取集群对象的模式（随机）
func getProducerMode(producerMap map[string]*list.List, size int) KafkaObject{
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := random.Intn(size)
	key := conf.GetConf().Kafka.Addr[index]
	kafkaObject := KafkaObject{}
	kafkaObject.Key = key
	kafkaObject.Producer = producerMap[key]
	return kafkaObject;
}
//获取池子对象
func GetKafkaProducer() ProducerObject{
	//对连接池取值进行加锁
	lock.Lock()
	defer lock.Unlock()
	kafkaObject := getProducerMode(KafkaProducer, len(conf.GetConf().Kafka.Addr))
	kafkaList := kafkaObject.Producer
	if kafkaList.Len() == 0 {
		return ProducerObject{"",nil}
	}
	//获取开头元素
	syncProducer := kafkaList.Front().Value
	//移除该元素
	KafkaProducer[kafkaObject.Key].Remove(kafkaList.Front())
	producer := syncProducer.(sarama.SyncProducer)
	producerObject := ProducerObject{}
	producerObject.Key = kafkaObject.Key
	producerObject.Producer = producer
	return producerObject
}
//对象归还
func PutKafkaProducer(producer ProducerObject) bool{
	//对连接池归还进行加锁
	lock.Lock()
	defer lock.Unlock()
	//添加该元素
	KafkaProducer[producer.Key].PushBack(producer.Producer)
	return true
}
//剔除宕机的kafka服务节点
func RemoveInvalidProducer(kafkaIp string) bool{
	if KafkaProducer[kafkaIp] != nil {
		InvalidProducer[kafkaIp] = KafkaProducer[kafkaIp]
		delete(KafkaProducer, kafkaIp)
	}
	return true
}
//恢复之前宕机的kafka服务节点
func AddInvalidProducer(kafkaIp string) bool{
	if InvalidProducer[kafkaIp] != nil {
		KafkaProducer[kafkaIp] = InvalidProducer[kafkaIp]
		delete(InvalidProducer, kafkaIp)
	}
	return true
}