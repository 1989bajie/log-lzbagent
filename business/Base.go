package business
import(
	"log-lzbagent/queue"
	//"log-lzbagent/conf"
)
type Base struct{
	rulePass bool
	ruleMsg string
}
var channel = make(chan bool,100)
//数据规则统一校验
func (base *Base)dataRule(logData map[string]interface{}) Base{
	base.rulePass = false
	if len(logData) == 0 {
		base.ruleMsg = "log data type error"
		return *base
	}
	if logData["target_id"] == nil || logData["target_id"] == "" {
		base.ruleMsg = "log target_id is null"
		return *base
	}
	if logData["action"] == nil || logData["action"] == "" {
		base.ruleMsg = "log action is null"
		return *base
	}
	if logData["source"] == nil || logData["source"] == "" {
		base.ruleMsg = "log source is null"
		return *base
	}
	base.rulePass = true
	return *base
}
//异步执行队列插入
func (base *Base)asyncQueuePut(queueData *map[string]interface{}){
	channel <- true
	go queue.KafkaProducer(queueData, channel)
}
