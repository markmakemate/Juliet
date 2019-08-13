package Components

import (
	"encoding/json"
	"net/http"
	"sync/atomic"
)

/**
Components 实现
traffic 不论如何分流最终都要返回一个已经配置好 experiment 的 Traffic 对象
Domain, layer 和 experiment Component 都必须实现 Init(uint64) 和 Serialize() 方法
*/

/**
Domain 实现
*/
func (domain *Domain) Init(name string, id uint64) {
	domain.Id = atomic.LoadUint64(&id)
	domain.Name = name
}

//Receive a traffic from session manager and do the split-flow progress
func (domain *Domain) ReceiveTraffic(traffic *Traffic) Traffic {
	/**
	TODO
	对 traffic 进行 domain 层分流
	*/

}

//Serialize domain resource
func (domain *Domain) Serialize() {

}

//Delete layer with layerId from domain and LayerContainer
func (domain *Domain) DeleteLayer(layerId uint64) {
	domain.LayerIdPool = DeleteFromUint64Array(domain.LayerIdPool, layerId)
	ConfigurationOfContainerManager.LayerContainer.Eject(layerId)
}

//Insert a given layer into domain and LayerContainer
func (domain *Domain) InsertLayer(layer *Layer) {
	domain.LayerIdPool = append(domain.LayerIdPool, layer.Id)
}

/**
Layer 实现
*/
func (layer *Layer) Init(name string, layerId uint64) {
	layer.Id = atomic.LoadUint64(&layerId)
	layer.Name = name
}

//Insert a given experiment into layer and ExptContainer
func (layer *Layer) InsertExperiment(expt *Experiment) {
	ConfigurationOfContainerManager.ExptContainer.Inject(expt)
	layer.ExptIdPool = append(layer.ExptIdPool, expt.Id)
}

//Delete experiment with exptId from layer and ExptContainer
func (layer *Layer) DeleteExperiment(exptId uint64) {
	ConfigurationOfContainerManager.ExptContainer.Eject(exptId)
	layer.ExptIdPool = DeleteFromUint64Array(layer.ExptIdPool, exptId)
}

//Receive a traffic and do the split-flow progress
func (layer *Layer) ReceiveTraffic(traffic *Traffic) Traffic {
	/**
	TODO
	对 traffic 做 layer 层分流
	*/

}

//Serialize the layer resource
func (layer *Layer) Serialize() {

}

/**
Experiment 实现
*/

//Serialize the experiment resources
func (experiment *Experiment) Serialize() {

}

func (experiment *Experiment) Init(name string, exptId uint64) {
	experiment.Id = atomic.LoadUint64(&exptId)
	experiment.Name = name
}

func (experiment *Experiment) SetExptConfig(config ExptConfig) {
	experiment.Type = config.Type
	experiment.Ratio = config.Ratio
	experiment.Timestamp = config.Timestamp
}

func (experiment *Experiment) toString() []byte {
	param := string(experiment.Param.toString())
	result := "{exptId: " + string(experiment.Id) + ", type: " + experiment.Type + ", param: " + param + "}"
	return []byte(result)
}

/*
分流器实现
*/
type Diversor struct {
}

func (d *Diversor) Hash(traffic Traffic) uint64 {
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
func (traffic *Traffic) Send(signal chan int) {
	go func() {
		if <-signal == 1 {
			var data []string
			for x := range traffic.ExptList {
				data = append(data, string(traffic.ExptList[x].toString()))
			}
			jsonStr, _ := json.Marshal(data)
			status, err := traffic.Writer.Write(jsonStr)
			PrintLog(status, err, traffic)
		}
	}()
}

//Insert an experiment in traffic
func (traffic *Traffic) InsertExpt(experiment Experiment) {
	traffic.ExptList = append(traffic.ExptList, experiment)
}
