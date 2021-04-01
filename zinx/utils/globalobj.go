package utils

import (
	"encoding/json"
	"go_code/zinx/ziface"
	"io/ioutil"
)

var GlobalObject *GlobalObj

type GlobalObj struct {
	TcpServer         ziface.IServer
	Host              string
	TcpPort           int
	Name              string
	Version           string
	MaxConn           int
	MaxPackageSize    uint32
	WorkerPoolSize    uint32
	MaxWorkerTaskSize uint32
}

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("G:\\Golang\\src\\go_code\\zinxV1.0Demo\\conf\\zinx.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

func init() {
	GlobalObject = &GlobalObj{
		Name:              "ZinxServerApp",
		Version:           "V0.9",
		Host:              "0.0.0.0",
		TcpPort:           8888,
		MaxConn:           1000,
		MaxPackageSize:    4096,
		WorkerPoolSize:    10,
		MaxWorkerTaskSize: 1024,
	}
	GlobalObject.Reload()
}
