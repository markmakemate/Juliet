package common

/**
common implementation of Juliet A/B test platform

Version 0.0.1-SNAPSHOT

Data: 2019.8

Author: Song Yi

Contact: yi.song@caibeike.com

Copyright: CaiBeiKe
*/

import (
	"Juliet/Utils"
	"encoding/json"
	"math/rand"
	"net/http"
	"reflect"
	"sync"
	"sync/atomic"
)

/**
common 实现
traffic 不论如何分流最终都要返回一个已经配置好 exp list 的 Traffic 对象
Domain, layer 和 experiment Component 都必须实现 Init(uint64) 和 Serialize() 方法
*/


//全局hash
func Hash(key interface{}) uint64 {
	_, ok1 := key.(string)
	_, ok2 := key.(uint64)
	if ok1 && !ok2{
		return Utils.TransformString2Uint64(key.(string))
	}else if !ok1 && ok2{
		return key.(uint64) % Diversion.Base
	}else{
		return 0
	}
}

/**
Domain 实现
*/
func (domain *Domain) Init(id uint64) {
	domain.Id = atomic.LoadUint64(&id)
}

//Serialize domain resource
func (domain *Domain) Serialize() {

}

//Del layer with layerId from domain and LayerContainer
func (domain *Domain) DelLayer(layerId uint64) {
	domain.Lpc.Eject(layerId)
	domain.Mutex.Lock()
	domain.Mutex.Unlock()
}

//Insert a given layer into domain and LayerContainer
func (domain *Domain) InsertLayer(layer *Layer) {
	domain.Mutex.Lock()
	domain.Lpc.Inject(layer)
	domain.Mutex.Unlock()
}

/**
Layer 实现
*/
func (layer *Layer) Init(layerId uint64) {
	layer.Id = atomic.LoadUint64(&layerId)
}

//Insert a given experiment into layer and expIdContainer
func (layer *Layer) InsertExp(exp *Exp) {
	layer.Mutex.Lock()
	layer.Epc.Inject(exp)
	layer.Mutex.Unlock()
}

//Del experiment with expId from layer and expIdContainer
func (layer *Layer) DelExp(expId uint64) {
	layer.Epc.Eject(expId)
}


//Receive a traffic and do the split-flow progress
func (layer *Layer) ReceiveTraffic(traffic *Traffic) Traffic {
	/**
	TODO
	对 traffic 做 layer 层分流
	*/
	traffic.InsertExp(Diversion.GetExp(traffic, layer.Epc))
	return *traffic
}

//Serialize the layer resource
func (layer *Layer) Serialize() {

}

/**
experiment 实现
*/

//Serialize the experiment resources
func (exp *Exp) Serialize() {

}

func (exp *Exp) Init(expId uint64) {
	exp.Id = atomic.LoadUint64(&expId)
}

func (exp *Exp) SetExpIdConfig(config ExpIdConfig) {
	exp.Timestamp = config.Timestamp
}

func (exp *Exp) String() []byte {
	result, _ := json.Marshal(
		map[string]interface{}{
		"id": string(exp.Id),
		"Params": exp.Params.String(),
		"start_time": exp.start.String(),
		"end_time": exp.end.String()},
		)
	return result
}

func (exp *Exp) GetParam(paramId string) Parameter{
	return exp.Params.Get(paramId)
}

func (exp *Exp) EjectParam(paramId string) {
	exp.Params.Eject(paramId)
}

func (exp *Exp) InjectParam(param Parameter) {
	exp.Params.Inject(reflect.ValueOf(param))
}
/*
分流器实现
*/

type feature []string

func (f *feature) extract(traffic *Traffic) *feature {
	return f
}

func (f *feature) get() string {
	return (*f)[rand.Int() % (len(*f) - 1)]
}

type Diversor struct {
	mutex sync.Mutex
	features *feature
	Base uint64
}

func (diver *Diversor) Hash(traffic *Traffic) uint64 {
	if traffic.LayerId != 0 {
		return traffic.LayerId
	}
	return Hash(*diver.features.extract(traffic))
}

func (diver *Diversor) GetExp(traffic *Traffic, epc *ExpContainer) Exp {
	diver.features = diver.features.extract(traffic)
	id := Hash(diver.features.get())
	return *epc.Get(id)
}
/**
Traffic实现
*/
func (traffic *Traffic) Init(writer http.ResponseWriter, r *http.Request,
	domainId uint64, layerId uint64) {
	traffic.Request = r
	traffic.DomainId = domainId
	traffic.Writer = writer
	traffic.LayerId = layerId
}

//Send the experiment args to clients
func (traffic *Traffic) Send() {
	go func() {
		traffic.Mutex.Lock()
		status, err := traffic.Writer.Write(traffic.String())
		traffic.Mutex.Unlock()
		PrintLog(status, err, traffic)

	}()
}

//Insert an experiment in traffic
func (traffic *Traffic) InsertExp(experiment Exp) {
	traffic.Mutex.Lock()
	traffic.ExpList = append(traffic.ExpList, experiment)
	traffic.Mutex.Unlock()
}

func (traffic *Traffic) String() []byte {
	var data []string
	for x := range traffic.ExpList {
		data = append(data, string(traffic.ExpList[x].String()))
	}
	jsonStr, _ := json.Marshal(data)
	return jsonStr
}

