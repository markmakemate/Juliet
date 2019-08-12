package Components

import (
	"caibeike-abtest/Managers"
	"github.com/golang/protobuf/ptypes/timestamp"
	"net/http"
)

type Traffic struct {
	Writer   http.ResponseWriter
	Request  *http.Request
	DomainId uint64
	LayerId  uint64
	Expt     Experiment
}

type Domain struct {
	Name        string
	Id          uint64
	LayerIdPool []uint64
}

type ExptConfig struct {
	Type      string
	Ratio     float32
	Timestamp timestamp.Timestamp
}

type Experiment struct {
	Id        uint64
	Type      string
	Param     Parameter
	Name      string
	Ratio     float32
	Timestamp timestamp.Timestamp
}

type Layer struct {
	Name       string
	Id         uint64
	ExptIdPool []uint64
}

/**
Components pkg is
*/
type BaseComponent interface {
	Init(name string, id uint64)
}

type ConfigOfComponents struct {
	ContainerManager *Managers.ContainerManager
	Diversor         *Diversor
}

//config variable is the global configuration of Component package
var Config ConfigOfComponents

/*
Parameter：最小的控制单元；
toString()方法是将客户端需要的参数集合转换成字节流数组；
Receive()方法是接受一组参数，修改原参数（若存在）并持久化；
Parameter通过RPC实行对不同类型资源的管理
Parameter的具体实现要根据不同的业务进行单独开发。
*/
type Parameter interface {
	toString() []byte
	Receive([]byte)
}
