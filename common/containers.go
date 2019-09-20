package common

import (
	"Juliet/Utils"
	"fmt"
	"reflect"
	"time"
)


/**
The Container Implementation of Juliet A/B test platform

Version 0.0.1-SNAPSHOT

Data: 2019.8

Author: Song Yi

Contact: yi.song@caibeike.com

Copyright: CaiBeiKe
*/


/**
Domain Container implementation
*/
func (dc *DomainContainer) Init(uuid string) {
	dc.uuid = uuid
}

//Inject a given domain into Container
func (dc *DomainContainer) Inject(domain *Domain) {
	dc.mutex.Lock()
	dc.DomainList = append(dc.DomainList, domain)
	dc.DomainMapper[domain.Id] = domain
	dc.Timestamp = time.Now()
	dc.mutex.Unlock()
}

//Del a domain with id from Container
func (dc *DomainContainer) Eject(id uint64) {
	dc.mutex.Lock()
	delete(dc.DomainMapper, id)
	err := Utils.DelFromArray(dc.DomainList, dc.DomainList[id])
	dc.mutex.Unlock()
	fmt.Println(err)
}

//Find a domain by id from Container
func (dc *DomainContainer) Get(id uint64) *Domain {
	return dc.DomainMapper[id]
}

/**
Layer Container implementation
*/
func (lc *LayerContainer) Init(uuid string) {
	lc.uuid = uuid
}

//Inject a given layer into Container
func (lc *LayerContainer) Inject(layer *Layer) {
	lc.mutex.Lock()
	if _, ok := lc.LayerMapper[layer.Id]; !ok{
		lc.LayerList = append(lc.LayerList, layer)
		lc.LayerMapper[layer.Id] = layer
		lc.Timestamp = time.Now()
	} else {
		old := lc.LayerMapper[layer.Id]
		Utils.Log(Utils.DelFromArray(lc.LayerList, old))
		lc.LayerList = append(lc.LayerList, layer)
		lc.LayerMapper[layer.Id] = layer
	}
	lc.mutex.Unlock()
}

//Del a layer with id from Container
func (lc *LayerContainer) Eject(id uint64) {
	lc.mutex.Lock()
	delete(lc.LayerMapper, id)
	err := Utils.DelFromArray(lc.LayerList, lc.LayerList[id])
	lc.mutex.Unlock()
	fmt.Println(err)
}

//Find a layer by id from Container
func (lc *LayerContainer) Get(id uint64) *Layer {
	return lc.LayerMapper[id]
}


/**
Experiment Container implementation
*/
func (ec *ExpContainer) Init(uuid string) {
	ec.uuid = uuid
}

//Inject a given layer into Container
func (ec *ExpContainer) Inject(Exp *Exp) {
	ec.mutex.Lock()
	ec.ExpList = append(ec.ExpList, Exp)
	ec.ExpMapper[Exp.Id] = Exp
	ec.mutex.Unlock()
}

//Del a layer with id from Container
func (ec *ExpContainer) Eject(id uint64) {
	ec.mutex.Lock()
	delete(ec.ExpMapper, id)
	err := Utils.DelFromArray(ec.ExpList, ec.ExpList[id])
	ec.mutex.Unlock()
	fmt.Println(err)
}

//Find a layer by id from Container
func (ec *ExpContainer) Get(id uint64) *Exp {
	return ec.ExpMapper[id]
}

func (ec *ExpContainer) InjectParamInExp(param Parameter,
	expId uint64) {
	ec.Get(expId).InjectParam(param)
}

func (ec *ExpContainer) EjectParamInExp(Type string,
	expId uint64) {
	ec.Get(expId).EjectParam(Type)
}
/**
Parameter Container Implementation
*/

//Inject a given parameter of type into Container
func (pc *ExpParamContainer) Inject(param reflect.Value) {
	pc.mutex.Lock()
	paramId := getId(param)
	pc.ParamMapper[paramId] = param
	pc.mutex.Unlock()
}

func getId(param reflect.Value) string {
	if param.Kind() == reflect.Map{
		typeValue := reflect.ValueOf("Id")
		mapIndex := param.Elem().MapIndex(typeValue)
		return mapIndex.Interface().(string)
	} else if param.Kind() == reflect.Struct{
		var value []reflect.Value
		methodByName := param.Elem().MethodByName("GetId")
		return methodByName.Call(value)[0].Interface().(string)
	}
	return ""
}

//Del a parameter with type from Container
func (pc *ExpParamContainer) Eject(Type string) {
	pc.mutex.Lock()
	delete(pc.ParamMapper, Type)
	pc.mutex.Unlock()
}

//Find a parameter with type from Container
func (pc *ExpParamContainer) Get(paramId string) Parameter {
	return pc.ParamMapper[paramId].Elem().Interface().(Parameter)
}

func (pc *ExpParamContainer) String() []string {
	var result []string
	for _, v :=range pc.ParamMapper{
		var value []reflect.Value
		dataValue := v.MethodByName("String").Call(value)
		result = append(result,
			dataValue[0].Interface().(string))
	}
	return result
}

func (gpc *LayerParamContainer) Push(parameter Parameter)  {
	gpc.paramMapper[parameter.GetType()] = parameter
}

func (gpc *LayerParamContainer) GetParam(T string) Parameter {
	return gpc.paramMapper[T]
}

func (gpc *LayerParamContainer) SetParam(T string,
	data []byte) *LayerParamContainer {
	gpc.paramMapper[T].Receive(data)
	return gpc
}

func (gpc *LayerParamContainer) Register(rcv interface{}) {
	register(gpc, rcv)
}

func register(gpc *LayerParamContainer, rcv interface{})  {
	reflect.Indirect(reflect.ValueOf(rcv)).Type().Name()
}
