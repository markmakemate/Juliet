package Components

import (
	"fmt"
	"net/http"
)

/**
所有的生成器的实现
*/

/**
Domain 生成器
*/
func NewDomain(name string, id uint64) *Domain {
	return &Domain{
		Name:        name,
		Id:          id,
		LayerIdPool: []uint64{},
	}

}

/**
Layer 生成器
*/
func NewLayer(layerId uint64, name string) *Layer {
	return &Layer{
		Name:       name,
		Id:         layerId,
		ExptIdPool: []uint64{},
	}
}

/**
Experiment生成器
*/
func NewExperiment(exptId uint64, name string,
	config ExptConfig) *Experiment {
	return &Experiment{
		Id:        exptId,
		Type:      config.Type,
		Param:     config.Param,
		Name:      name,
		Ratio:     config.Ratio,
		Timestamp: config.Timestamp,
	}
}

/**
Traffic生成器
*/
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

/**
container 生成器
根据配置自适应生成对应的 container
若配置出现错误，打印错误提示并返回 nil
*/
func NewContainer() interface{} {
	config = ConfigurationOfContainer
	list := transformListType(config.ComponentList)
	mapper := transformMapperType(config.ComponentMapper)
	if list != nil && mapper != nil {
		switch config.Type {
		case 1:
			return &DomainContainer{
				uuid:         config.UUID,
				DomainList:   list.([]*Domain),
				DomainMapper: mapper.(map[uint64]*Domain),
			}
		case 2:
			return &LayerContainer{
				uuid:        config.UUID,
				LayerList:   list.([]*Layer),
				LayerMapper: mapper.(map[uint64]*Layer),
			}
		case 3:
			return &ExperimentContainer{
				uuid:       config.UUID,
				ExptList:   list.([]*Experiment),
				ExptMapper: mapper.(map[uint64]*Experiment),
			}
		case 4:
			return &ParameterContainer{
				uuid:        config.UUID,
				ParamMapper: mapper.(map[string]*Parameter),
			}
		default:
			return nil
		}
	} else {
		fmt.Println("Configuration occurs unidentified error! Please check your configuration")
		return nil
	}
}
