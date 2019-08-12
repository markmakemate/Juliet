package Components

import "sync/atomic"

/**
Layer实现
*/
func (layer *Layer) Init(name string, layerId uint64) {
	layer.Id = atomic.LoadUint64(&layerId)
	layer.Name = name
	layer.ExptIdPool = []uint64{}
}

func (layer *Layer) InsertExperiment(expt *Experiment) {
	Config.ContainerManager.ExperimentContainer.Inject(expt)
	layer.ExptIdPool = append(layer.ExptIdPool, expt.Id)
}

func (layer *Layer) DeleteExperiment(exptId uint64) {
	Config.ContainerManager.ExperimentContainer.Eject(exptId)
	layer.ExptIdPool = DeleteFromUint64Array(layer.ExptIdPool, exptId)
}

func (layer *Layer) ReceiveTraffic(traffic *Traffic) {
	/**
	TODO
	对traffic做layer层分流
	*/

}
