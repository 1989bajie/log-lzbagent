package persistence
import(
	"strconv"
	"log-lzbagent/log"
	"log-lzbagent/conf"
	"log-lzbagent/po"
	"gopkg.in/mgo.v2"
)
var mgoSession *mgo.Session = nil
type Mongo struct {
	session *mgo.Session
}
func getSession() *mgo.Session {
	mongoAddr := "mongodb://" + conf.GetConf().Database.Mongo.Addr + ":" + strconv.Itoa(conf.GetConf().Database.Mongo.Port)
	if mgoSession == nil {
		session, err := mgo.Dial(mongoAddr)
		if err != nil {
			panic(err)
		}
		//设置连接池大小(默认4096)
		session.SetPoolLimit(100)
		mgoSession = session
	}
	return mgoSession.Clone()
}
func (*Mongo)GetMongoSession() *mgo.Session {
	return getSession()
}
func (*Mongo)Insert(po po.PoInterface) bool{
	session := getSession()
	db := session.DB(conf.GetConf().Database.Mongo.Database)
	err := db.C(po.GetTable()).Insert(po)
	session.Close()
	var result bool = true
	if(err != nil){
		log.Error("insert err:" + err.Error())
		result = false
	}
	return result
}