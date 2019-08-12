package Container

import (
	"caibeike-abtest/Components"
	"caibeike-abtest/Utils"
)

type Config struct {
	uuid            string
	ComponentList   interface{}
	ComponentMapper interface{}
	/**
	Type == 1: Domain configuration
	Type == 2: Layer configuration
	Type == 3: Experiment configuration
	*/
	Type int
}

//Domain Container
type DomainContainer struct {
	uuid         string
	DomainList   []*Components.Domain
	DomainMapper map[uint64]*Components.Domain
}

func (dc *DomainContainer) Inject(layer *Components.Domain) {
	dc.DomainList = append(dc.DomainList, layer)
	dc.DomainMapper[layer.Id] = layer
}
func (dc *DomainContainer) Eject(id uint64) {
	delete(dc.DomainMapper, id)
	Utils.DeleteFromArray(dc.DomainList, dc.DomainList[id])
}
func (dc *DomainContainer) Get(id uint64) *Components.Domain {
	return dc.DomainMapper[id]
}

//Layer Container
type LayerContainer struct {
	uuid        string
	LayerList   []*Components.Layer
	LayerMapper map[uint64]*Components.Layer
}

func (lc *LayerContainer) New() {

}
func (lc *LayerContainer) Inject(layer *Components.Layer) {
	lc.LayerList = append(lc.LayerList, layer)
	lc.LayerMapper[layer.Id] = layer
}
func (lc *LayerContainer) Eject(id uint64) {
	delete(lc.LayerMapper, id)
	Utils.DeleteFromArray(lc.LayerList, lc.LayerList[id])
}
func (lc *LayerContainer) Get(id uint64) *Components.Layer {
	return lc.LayerMapper[id]
}

//Experiment Container
type ExperimentContainer struct {
	uuid       string
	ExptList   []*Components.Experiment
	ExptMapper map[uint64]*Components.Experiment
}

func (ec *ExperimentContainer) Inject(expt *Components.Experiment) {
	ec.ExptList = append(ec.ExptList, expt)
	ec.ExptMapper[expt.Id] = expt
}
func (ec *ExperimentContainer) Eject(id uint64) {
	delete(ec.ExptMapper, id)
	Utils.DeleteFromArray(ec.ExptList, ec.ExptList[id])
}
func (ec *ExperimentContainer) Get(id uint64) *Components.Experiment {
	return ec.ExptMapper[id]
}

func NewContainer(config Config) interface{} {
	list := transformListType(config.ComponentList)
	mapper := transformMapperType(config.ComponentMapper)
	if list != nil && mapper != nil && config.Type == 1 {
		return &DomainContainer{
			uuid:         config.uuid,
			DomainList:   list.([]*Components.Domain),
			DomainMapper: mapper.(map[uint64]*Components.Domain),
		}
	} else if config.Type == 2 {
		return &LayerContainer{
			uuid:        "",
			LayerList:   list.([]*Components.Layer),
			LayerMapper: mapper.(map[uint64]*Components.Layer),
		}
	} else if config.Type == 3 {
		return &ExperimentContainer{
			uuid:       "",
			ExptList:   list.([]*Components.Experiment),
			ExptMapper: mapper.(map[uint64]*Components.Experiment),
		}
	} else {
		return nil
	}
}
