package main

import (
	"log-lzbagent/conf"
	"log-lzbagent/health"
	"log-lzbagent/log"
	"log-lzbagent/pool"
	"log-lzbagent/server"
)

func main() {
	//配置文件初始化
	conf.InitConf()
	//日志组件初始化
	log.LogInit()
	//kafka连接池初始化
	pool.Init()
	//健康监测组件启动
	health.HealthRun()
	//服务启动
	server.TcpRpcStart()
}
