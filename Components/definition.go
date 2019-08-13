package Components

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"net/http"
)

type Traffic struct {
	Writer   http.ResponseWriter
	Request  *http.Request
	DomainId uint64
	LayerId  uint64
	ExptList []Experiment
}

type Domain struct {
	Name        string
	Id          uint64
	LayerIdPool []uint64
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
	Serialize()
}

/*
Parameter：最小的控制单元；
toString() 方法是将客户端需要的参数集合转换成字节流数组；
Receive() 方法是接受一组参数，修改原参数（若存在）并持久化；
Parameter 通过 RPC 实行对不同类型资源的管理
Parameter 的具体实现要根据不同的业务进行单独开发。
*/
type Parameter interface {
	toString() []byte
	Receive([]byte)
	GetType() string
}

/**
Domain 容器，用于 Domain 的依赖注入和 Domain 的根据 Id 管理
*/
type DomainContainer struct {
	uuid         string
	DomainList   []*Domain
	DomainMapper map[uint64]*Domain
}

/**
Layer 容器，用于 Layer 的依赖注入和 Layer 的根据 Id 管理
*/
type LayerContainer struct {
	uuid        string
	LayerList   []*Layer
	LayerMapper map[uint64]*Layer
}

/**
Experiment 容器，用于 Experiment 的依赖注入和 Experiment 的根据 Id 管理
*/
type ExperimentContainer struct {
	uuid       string
	ExptList   []*Experiment
	ExptMapper map[uint64]*Experiment
}

/**
Parameter 容器，用于 Parameter 的依赖注入和 Parameter 的根据 Type 管理
*/
type ParameterContainer struct {
	uuid        string
	ParamMapper map[string]*Parameter
}
