package server
import (
    "strconv"
	"log-lzbagent/conf"
    "net"
    "log-lzbagent/business"
    "net/rpc"
	"net/rpc/jsonrpc"
)

func TcpRpcStart() {
    //业务处理类注册
    record := new(business.Record)
    rpc.Register(record)
    tcpAddr := conf.GetConf().Local.Addr + ":" + strconv.Itoa(conf.GetConf().Local.Port)
    addr, _ := net.ResolveTCPAddr("tcp", tcpAddr)
    ln, e := net.ListenTCP("tcp", addr)
    if e != nil {
        panic(e)
    }   
    for {
        conn, e := ln.Accept()
        if e != nil {
            continue
        }
        go jsonrpc.ServeConn(conn)
    }   
}