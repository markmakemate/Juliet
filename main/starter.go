package main

import (
	"Juliet/Utils"
	"Juliet/common"
	"Laurence/elem"
	"net/http"
	"net/rpc"
	"reflect"
	"time"
)

/**
The starter of Juliet A/B test platform

Version 0.0.1-SNAPSHOT

Data: 2019.8

Author: Song Yi

Contact: yi.song@caibeike.com

Copyright: CaiBeiKe
*/

type Type2Param map[string]interface{}


//参数注册中心，通过参数的类型管理
var ParamRegisterCenter  = Type2Param{
	"recommend": new(common.Recommend),
}

type Layer struct {
	LayerId uint64
	Exp elem.Exp
}

type WebService struct {
	Data []byte
	Result *string
}

func (web *WebService) UpdateParamInExpOfLayer(arg Layer, result *string) error {
	common.UpdateParamInExpOfLayer(arg.Exp.Param.Values, arg.Exp.Id, arg.LayerId)
	return nil
}

func (web *WebService) DelLayer(layerId uint64, result *string) error {
	common.DelLayer(layerId)
	*result = "success"
	return nil
}

func (web *WebService) DelExpInLayer(arg Layer, result *string) error  {
	common.DelExpInLayer(arg.LayerId, arg.Exp.Id)
	return nil
}

func (web *WebService) DelParamInExpOfLayer(arg Layer, result *string) error  {
	common.DelParamInExpOfLayer(arg.Exp.Param.Id, arg.Exp.Id, arg.LayerId)
	return nil
}

func (web *WebService) InsertExpInLayer(arg Layer, result *string) error  {
	paraContainer := new(common.ExpParamContainer)
	param := ParamRegisterCenter[arg.Exp.Param.Type]
	paramValue := reflect.ValueOf(param)
	value := []reflect.Value{reflect.ValueOf(arg.Exp.Param.Id),
		reflect.ValueOf(arg.Exp.Param.Type)}
	paramValue.MethodByName("Init").Call(value)   //初始化参数
	paramValue.MethodByName("Receive").Call([]reflect.Value{reflect.ValueOf(arg.Exp.Param.Values)}) //注入数据
	paraContainer.Inject(paramValue)  //注入到容器中
	exp := &common.Exp{
		Id:        arg.Exp.Id,
		Params:    paraContainer,
		Timestamp: time.Time{},
		Start:     arg.Exp.Start,
		End:      arg.Exp.End,
	}
	common.InsertExpInLayer(arg.LayerId, exp)
	return nil
}

func (web *WebService) InitParam() {

}

var Web = &WebService{}

func init()  {
	Utils.Log(rpc.Register(Web))
}

func main() {
	/**
	  服务启动流程：
	  1. 实例化连接管理器和分流器，配置必要的全局配置；
	  2. 刷新日志，从实验平台后台请求实验资源数据，对实验资源数据本地处理，记录日志；若实验平台宕机，则从日志中恢复部分数据；
	  3. 启动连接管理器的监听器，监听指定地址；
	  4. 启动 RpcServer 服务，等待实验平台调度；
	  5. 启动 HTTP Restful 服务，等待前端和业务系统调用；

	  A/B 测试平台只维护实验资源的元信息和完成服务端分流。通过 RPC 调用或者 HTTP 请求 RESTful 接口完成参数请求。数据拉取的过程由各个业务组自定义；
	  实验平台与测试平台解耦，通过 Rpc 互相调用。客户端或业务系统通过 RESTful 接口获取分流后的实验数据；
	  作为灰度发布系统，整个系统应与业务系统解耦，系统内部应高度解耦，包与包之间的引用一定要保持单向性，不能出现环状导包或包互引现象；
	  发生宕机后重新需要从实验平台拉取资源，日志的作用仅是在实验平台宕机未恢复，测试平台宕机重新启动后恢复部分数据；
	  除网络和日志相关功能会采用成熟的第三方库外，其他均不采用第三方库；
	  测试平台的数据应与实验平台保持最终一致，打点上报的数据应该由单独的数据平台来解析；
	*/
	rpc.HandleHTTP()
	Utils.Log(http.ListenAndServe(":8000", nil))
	/**
	sm := new(SessionManager)
	sm.Start(":8080")
	*/
}
