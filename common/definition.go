package common


/**
common definition of Juliet A/B test platform

Version 0.0.1-SNAPSHOT

Data: 2019.8

Author: Song Yi

Contact: yi.song@caibeike.com

Copyright: CaiBeiKe
 */

import (
	"net/http"
	"reflect"
	"sync"
	"time"
)


/**
Traffic encapsulates a HTTP request from one client, some of its Request parameters and a list of experiments. The main
job of Traffic is to package a http Request and processed by Resource Manager which will distribute an experiments list
for it then send the parameters to client with the signal form Session Manager. Such a progress is asynchronously executed.
 */
type Traffic struct {
	Mutex    sync.Mutex
	Writer   http.ResponseWriter
	Request  *http.Request
	DomainId uint64
	LayerId  uint64
	ExpList []Exp
}

/**
Domain is a bigger notion in Juliet. It not represents a specific group but shows the app a user is in. A domain has several
Layer which can get them from a Layer Pool by Id.
 */
type Domain struct {
	Mutex       sync.Mutex
	Id          uint64
	Lpc         *LayerContainer
}

/**
Experiment is one of the basic components in Juliet. It has only one objective to optimize and structure
of parameter. What's more, except an unique Id and a name, an experiment also has a timestamp which shows the
create time of itself.
 */
type Exp struct {
	Id        uint64
	Params    *ExpParamContainer
	Timestamp time.Time
	Start    time.Time
	End      time.Time
}

/**
Layer is the component which can control different application layers. A type will definitely map a parameter structure.
Each parameter structure can control a function as it defines. Every Layer has several Experiments and it can get them from
an Experiment Id pool.
 */
type Layer struct {
	Mutex      sync.Mutex
	Id         uint64
	Epc        *ExpContainer
}

/**
Basic Component Abstract Interface
*/
type BaseComponent interface {
	Init(id uint64)
	Serialize()
}

/**
Parameter：The tiniest controller which masters a given function, interface form；
String() is the method to transform args to JSON array；
Receive() is the method which will receive a byte array，modify (if exists) or create a parameter then serialize it；
GetType() is to get the type of parameter;
Init() is to initialize a Parameter by rpc;
Parameter's implementation is based on biz;
Parameter manage the components as plugins by dependency injection (DI);
*/
type Parameter interface {
	String() []byte
	Receive([]byte)
	GetType() string
	GetId()  string
	Init([]byte, *string) error
	Reset()
}

/**
Domain 容器，用于 Domain 的依赖注入和 Domain 的根据 Id 管理
*/
type DomainContainer struct {
	uuid         string
	DomainList   []*Domain
	DomainMapper map[uint64]*Domain
	mutex sync.Mutex
	Timestamp time.Time
}

/**
Layer 容器，用于 Layer 的依赖注入和 Layer 的根据 Id 管理
*/
type LayerContainer struct {
	uuid        string
	LayerList   []*Layer
	LayerMapper map[uint64]*Layer
	mutex sync.Mutex
	Timestamp time.Time
	Gpc *LayerParamContainer

}

/**
Experiment 容器，用于 Experiment 的依赖注入和 Experiment 的根据 Id 管理
*/
type ExpContainer struct {
	uuid       string
	ExpList   []*Exp
	ExpMapper map[uint64]*Exp
	mutex sync.Mutex
	Timestamp time.Time
}

/**
实验的Parameter 容器，用于管理每个实验的 Parameter 注入和 Parameter 的根据 Type 管理
*/
type ExpParamContainer struct {
	ParamMapper map[string]reflect.Value
	mutex sync.Mutex
	Timestamp time.Time
}

/**
全局 Param 容器，用于管理全局 Param 的依赖注入
 */

type ParamMap map[string]Parameter

// 每个业务层 Param 的元数据容器
type LayerParamContainer struct {
	paramMapper ParamMap
	Timestamp time.Time
}



