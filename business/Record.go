package business
import(
    "errors"
    "encoding/json"
    "log-lzbagent/log"
    "log-lzbagent/po"
    "log-lzbagent/conf"
    "gopkg.in/mgo.v2/bson"
    "log-lzbagent/persistence"
)
type Record struct {
    Base
}
//日志存储(队列)
func (record *Record) InsertLog(args string, reply *map[string]interface{}) error {
    var logData map[string]interface{}
    json.Unmarshal([]byte(args),&logData)
    //数据校验
    base := record.dataRule(logData)
    if(base.rulePass == false){
        log.Error(base.ruleMsg)
        return errors.New(base.ruleMsg)
    }
    record.asyncQueuePut(&logData)
    return nil
}
//日志查询
func (record *Record) QueryLog(args string, reply *string) error {
    var mongoBase persistence.Mongo
    session := mongoBase.GetMongoSession()
    defer session.Close()
	db := session.DB(conf.GetConf().Database.Mongo.Database)
    bsonM := bson.M{}
    json.Unmarshal([]byte(args),&bsonM)
    var poList []po.MongoRecord
    err := db.C("record").Find(bsonM).All(&poList)
    if err != nil {
        log.Error(err.Error())
        return err
    }
    poByte,_ := json.Marshal(poList)
    poString := string(poByte)
    *reply = poString
    return nil
}
