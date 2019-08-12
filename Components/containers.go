package Components

import (
	"caibeike-abtest/Utils"
)

var config ConfigOfContainer

//Domain Container
type DomainContainer struct {
	uuid         string
	DomainList   []*Domain
	DomainMapper map[uint64]*Domain
}

func (dc *DomainContainer) Inject(domain *Domain) {
	dc.DomainList = append(dc.DomainList, domain)
	dc.DomainMapper[domain.Id] = domain
}
func (dc *DomainContainer) Eject(id uint64) {
	delete(dc.DomainMapper, id)
	Utils.DeleteFromArray(dc.DomainList, dc.DomainList[id])
}
func (dc *DomainContainer) Get(id uint64) *Domain {
	return dc.DomainMapper[id]
}

//Layer Container
type LayerContainer struct {
	uuid        string
	LayerList   []*Layer
	LayerMapper map[uint64]*Layer
}

func (lc *LayerContainer) Inject(layer *Layer) {
	lc.LayerList = append(lc.LayerList, layer)
	lc.LayerMapper[layer.Id] = layer
}
func (lc *LayerContainer) Eject(id uint64) {
	delete(lc.LayerMapper, id)
	Utils.DeleteFromArray(lc.LayerList, lc.LayerList[id])
}
func (lc *LayerContainer) Get(id uint64) *Layer {
	return lc.LayerMapper[id]
}

//Experiment Container
type ExperimentContainer struct {
	uuid       string
	ExptList   []*Experiment
	ExptMapper map[uint64]*Experiment
}

func (ec *ExperimentContainer) Inject(expt *Experiment) {
	ec.ExptList = append(ec.ExptList, expt)
	ec.ExptMapper[expt.Id] = expt
}
func (ec *ExperimentContainer) Eject(id uint64) {
	delete(ec.ExptMapper, id)
	Utils.DeleteFromArray(ec.ExptList, ec.ExptList[id])
}
func (ec *ExperimentContainer) Get(id uint64) *Experiment {
	return ec.ExptMapper[id]
}

type ParameterContainer struct {
	uuid        string
	ParamMapper map[string]*Parameter
}

func (pc *ParameterContainer) Inject(param *Parameter, Type string) {
	pc.ParamMapper[Type] = param
}
func (pc *ParameterContainer) Eject(Type string) {
	delete(pc.ParamMapper, Type)
}
func (pc *ParameterContainer) Get(Type string) *Parameter {
	return pc.ParamMapper[Type]
}

/**
根据配置自适应生成对应的container
若配置出现错误，则返回nil
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
		}
	} else {
		return nil
	}
}
