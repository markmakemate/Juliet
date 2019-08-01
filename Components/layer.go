package Components

/*
type AbstractLayer interface {
	InsertDomain(domain Domain)
	GetId() uint64
	SetId(id uint64)
	InsertExperiment(expt Experiment)
}
*/
type Layer struct {
	DomainList       []*Domain
	Name             string
	Id               uint64
	ExperimentList   []*Experiment
	DomainMapper     map[string]*Domain
	ExperimentMapper map[string]*Experiment
	ParameterList    []*Parameter
	trafficChan      chan *Traffic
}

func (layer *Layer) InsertDomain(domain *Domain) {
	layer.DomainList = append(layer.DomainList, domain)
	layer.DomainMapper[domain.Name] = domain
}

func (layer *Layer) InsertExperiment(expt *Experiment) {
	layer.ExperimentList = append(layer.ExperimentList, expt)
	layer.ExperimentMapper[expt.Name] = expt
}

func (layer *Layer) ReceiveTraffic(traffic *Traffic) {
	layer.trafficChan <- traffic
}
