package Components

var config ConfigOfContainer

/**
Container 容器实现
*/

/**
Domain Container implementation
*/
func (dc *DomainContainer) Init(uuid string) {
	dc.uuid = uuid
}

//Inject a given domain into Container
func (dc *DomainContainer) Inject(domain *Domain) {
	dc.DomainList = append(dc.DomainList, domain)
	dc.DomainMapper[domain.Id] = domain
}

//Delete a domain with id from Container
func (dc *DomainContainer) Eject(id uint64) {
	delete(dc.DomainMapper, id)
	DeleteFromArray(dc.DomainList, dc.DomainList[id])
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
	lc.LayerList = append(lc.LayerList, layer)
	lc.LayerMapper[layer.Id] = layer
}

//Delete a layer with id from Container
func (lc *LayerContainer) Eject(id uint64) {
	delete(lc.LayerMapper, id)
	DeleteFromArray(lc.LayerList, lc.LayerList[id])
}

//Find a layer by id from Container
func (lc *LayerContainer) Get(id uint64) *Layer {
	return lc.LayerMapper[id]
}

/**
Experiment Container implementation
*/
func (ec *ExperimentContainer) Init(uuid string) {
	ec.uuid = uuid
}

//Inject a given layer into Container
func (ec *ExperimentContainer) Inject(expt *Experiment) {
	ec.ExptList = append(ec.ExptList, expt)
	ec.ExptMapper[expt.Id] = expt
}

//Delete a layer with id from Container
func (ec *ExperimentContainer) Eject(id uint64) {
	delete(ec.ExptMapper, id)
	DeleteFromArray(ec.ExptList, ec.ExptList[id])
}

//Find a layer by id from Container
func (ec *ExperimentContainer) Get(id uint64) *Experiment {
	return ec.ExptMapper[id]
}

/**
Parameter Container
*/
func (pc *ParameterContainer) Init(uuid string) {
	pc.uuid = uuid
}

//Inject a given parameter of type into Container
func (pc *ParameterContainer) Inject(param *Parameter, Type string) {
	if (*param).GetType() == Type {
		pc.ParamMapper[Type] = param
	}
}

//Delete a parameter with type from Container
func (pc *ParameterContainer) Eject(Type string) {
	delete(pc.ParamMapper, Type)
}

//Find a parameter with type from Container
func (pc *ParameterContainer) Get(Type string) *Parameter {
	return pc.ParamMapper[Type]
}
