package pool
import(
	"os"
	"sync"
	"strconv"
	"container/list"
	"log-lzbagent/conf"
)
var LocalFileProducer = list.New()
var fileLock sync.Mutex
//连接池初始化
func LocalFileInit() {
	var i int
	for i = 1; i<=conf.GetConf().Localfile.PoolSize; i++ {
		file := conf.GetConf().Localfile.Path + "lzb_file_queue_" + strconv.Itoa(i)
		_,err := os.Stat(file)
		var f *os.File
		if err != nil {
			//创建文件
			f,err = os.Create(file)
			if err != nil {
				panic(err)
			}
		}else{
			//打开文件
			f,err = os.OpenFile(file,os.O_RDWR,0666)
			if err != nil {
				panic(err)
			}
		}
		LocalFileProducer.PushBack(f)
	}
}
//获取池子对象
func GetFileProducer() *os.File{
	//对连接池取值进行加锁
	fileLock.Lock()
	defer fileLock.Unlock()
	if(LocalFileProducer.Len() == 0){
		return nil
	}
	//获取开头元素
	fileProducer := LocalFileProducer.Front().Value
	producer := fileProducer.(*os.File)
	//移除该元素
	LocalFileProducer.Remove(LocalFileProducer.Front())
	return producer
}
//对象归还
func PutFileProducer(producer *os.File) bool{
	//对连接池归还进行加锁
	fileLock.Lock()
	defer fileLock.Unlock()
	//添加该元素
	LocalFileProducer.PushBack(producer)
	return true
}