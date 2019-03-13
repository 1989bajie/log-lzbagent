package conf

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Conf struct {
	Local     Local     `yaml:"local"`
	Kafka     Kafka     `yaml:"kafka"`
	Localfile Localfile `yaml:"localfile"`
	Database  Database  `yaml:"database"`
}

type Local struct {
	FilePath   string `yaml:"file_path"`
	Addr       string `yaml:"addr"`
	Port       int    `yaml:"port"`
	ThreadMax  string `yaml:"thread_max"`
	LogMaxSize int    `yaml:"log_max_size"`
}

type Kafka struct {
	Addr     []string `yaml:"addr"`
	Port     int      `yaml:"port"`
	Topic    string   `yaml:"topic"`
	PoolSize int      `yaml:"pool_size"`
}

type Localfile struct {
	Path     string `yaml:"path"`
	PoolSize int `yaml:"pool_size"`
}

type Database struct {
	Mongo Mongo `yaml:"mongo"`
}

type Mongo struct {
	Addr     string `yaml:"addr"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

var conf *Conf

//InitConf
func InitConf() {
	yFile, err := ioutil.ReadFile("conf/conf.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yFile, &conf)
	if err != nil {
		panic(err)
	}
}

func GetConf() *Conf {
	return conf
}
