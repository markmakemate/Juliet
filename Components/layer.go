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
	DomainList []*Domain
	name string
	id uint64
	ExperimentList []*Experiment
	DomainMapper map[string] *Domain
	ExperimentMapper map[string]*Experiment
}

func (layer *Layer) SetId(id uint64) {
	layer.id = id
}
func (layer *Layer) GetId() uint64 {
	return layer.id
}
func (layer *Layer) GetName()string {
	return layer.name
}
func (layer *Layer)SetName(name string)  {
	layer.name = name
}
func (layer *Layer) InsertDomain(domain *Domain) {
	layer.DomainList = append(layer.DomainList, domain)
	layer.DomainMapper[domain.GetName()] = domain
}

func (layer *Layer) InsertExperiment(expt *Experiment) {
	layer.ExperimentList = append(layer.ExperimentList, expt)
	layer.ExperimentMapper[expt.GetName()] = expt
}