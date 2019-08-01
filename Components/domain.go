package Components

/*
type AbstractDomain interface {
	InsertLayer(layer *Layer)
	DiverseTraffic(traffic Traffic)
	GetId() uint64
	SetId(id uint64)
	GetDiversor() *Diversor
	SetDiversor(diversor Diversor)
}
*/
type Domain struct {
	trafficChan chan *Traffic
	Name        string
	LayerMapper map[uint64]*Layer
	LayerList   []*Layer
	Id          uint64
	ExptList    []*Experiment
}

func (domain *Domain) Init(name string, id uint64) {
	domain.Name = name
	domain.Id = id
}

func (domain *Domain) InsertLayer(layer *Layer) {
	domain.LayerMapper[layer.Id] = layer
	domain.LayerList = append(domain.LayerList, layer)
}

func (domain *Domain) ReceiveTraffic(traffic *Traffic) {
	domain.trafficChan <- traffic
}

func (domain *Domain) Diverse() {

}
func (domain *Domain) InsertExpt(expt *Experiment) {
	domain.ExptList = append(domain.ExptList, expt)
}
