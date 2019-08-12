package Components

import "net/http"

/**
所有的生成器的实现
*/
func NewDomain(name string, id uint64) *Domain {
	return &Domain{
		Name:        name,
		Id:          id,
		LayerIdPool: []uint64{},
	}

}

func NewLayer(layerId uint64, name string) *Layer {
	return &Layer{
		Name:       name,
		Id:         layerId,
		ExptIdPool: []uint64{},
	}
}

func NewExperiment(exptId uint64, name string, config ExptConfig) *Experiment {
	return &Experiment{
		Id:        exptId,
		Type:      config.Type,
		Param:     config.Param,
		Name:      name,
		Ratio:     config.Ratio,
		Timestamp: config.Timestamp,
	}
}

func NewTraffic(w http.ResponseWriter, r *http.Request,
	domainId uint64, layerId uint64) *Traffic {
	return &Traffic{
		Writer:   w,
		Request:  r,
		DomainId: domainId,
		LayerId:  layerId,
		Expt:     nil,
	}
}
