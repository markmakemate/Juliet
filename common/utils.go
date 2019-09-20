package common

import (
	"Juliet/Utils"
	"encoding/json"
	"fmt"
	"reflect"
)

//Print log on shell
func PrintLog(status int, err error, traffic *Traffic) {
	if err != nil {
		fmt.Printf("Client's url path:" + traffic.Request.URL.Path)
		fmt.Printf("The status is: " + string(status))
		fmt.Printf("The error info: " + err.Error())
	}
}


func UpdateParamInExpOfLayer(param[]byte, expId uint64, layerId uint64) {
	arg := new(Arg)
	Utils.Log(json.Unmarshal(param, arg))
	args := []reflect.Value{reflect.ValueOf(arg.data)}
	paramType := DomainEntity.Lpc.Get(layerId).Epc.Get(expId).GetParam(arg.paramId)
	methodType := reflect.ValueOf(paramType).MethodByName("Receive")
	parameter := methodType.Call(args)
	DomainEntity.Lpc.Get(layerId).Epc.InjectParamInExp(parameter[0].Interface().(Parameter), expId)
}

func  DelLayer(layerId uint64) {
	DomainEntity.Lpc.Eject(layerId)
}

func  DelExpInLayer(layerId uint64, expId uint64) {
	DomainEntity.Lpc.Get(layerId).Epc.Eject(expId)
}

func  DelParamInExpOfLayer(paramId string, expId uint64, layerId uint64) {
	DomainEntity.Lpc.Get(layerId).Epc.Get(expId).EjectParam(paramId)
}

//Insert layer with layerId in domain with domainId and LayerContainer
func  InsertLayer(layer *Layer) {
	DomainEntity.Lpc.Inject(layer)
}

func  InsertExpInLayer(layerId uint64, exp *Exp) {
	DomainEntity.Lpc.Get(layerId).InsertExp(exp)
}


func  UpdateLayer(layer *Layer) {
	DomainEntity.Lpc.Inject(layer)
}

func  UpdateExpOfLayer(exp *Exp, layerId uint64) {
	DomainEntity.Lpc.Get(layerId).DelExp(exp.Id)
	DomainEntity.Lpc.Get(layerId).InsertExp(exp)
}
