package Components

import (
	"net/http"
	"sync/atomic"
)

/**
Domain实现
*/
func (domain *Domain) Init(name string, id uint64) {
	domain.Name = name
	domain.Id = id
}

func (domain *Domain) ReceiveTraffic(traffic *Traffic) {
	/**
	TODO
	对traffic进行domain层分流
	*/

}
func (domain *Domain) Diverse() {

}

func (domain *Domain) Serialize() {

}
func (domain *Domain) DeleteLayer(layerId uint64) {
	domain.LayerIdPool = DeleteFromUint64Array(domain.LayerIdPool, layerId)
	ConfigurationOfContainerManager.LayerContainer.Eject(layerId)
}

func (domain *Domain) InsertLayer(layer *Layer) {
	domain.LayerIdPool = append(domain.LayerIdPool, layer.Id)
}

/**
Layer实现
*/
func (layer *Layer) Init(name string, layerId uint64) {
	layer.Id = atomic.LoadUint64(&layerId)
	layer.Name = name
}

func (layer *Layer) InsertExperiment(expt *Experiment) {
	ConfigurationOfContainerManager.ExptContainer.Inject(expt)
	layer.ExptIdPool = append(layer.ExptIdPool, expt.Id)
}

func (layer *Layer) DeleteExperiment(exptId uint64) {
	ConfigurationOfContainerManager.ExptContainer.Eject(exptId)
	layer.ExptIdPool = DeleteFromUint64Array(layer.ExptIdPool, exptId)
}

func (layer *Layer) ReceiveTraffic(traffic *Traffic) {
	/**
	TODO
	对traffic做layer层分流
	*/

}

/**
Experiment实现
*/
func (experiment *Experiment) Serialize() {

}

func (experiment *Experiment) Init(name string, exptId uint64) {
	experiment.Id = exptId
	experiment.Name = name
}

func (experiment *Experiment) SetExptConfig(config ExptConfig) {
	experiment.Type = config.Type
	experiment.Ratio = config.Ratio
	experiment.Timestamp = config.Timestamp
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

func (traffic *Traffic) Send(signal chan int) {
	go func() {
		x := <-signal
		if x == 1 {
			status, err := traffic.Writer.Write(traffic.Expt.Param.toString())
			if err != nil {
				PrintLog(status, err, traffic)
			}
		}
	}()
}

func (traffic *Traffic) SetExpt(experiment Experiment) {
	traffic.Expt = experiment
}
