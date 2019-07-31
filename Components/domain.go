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
	name string
	LayerMapper map[uint64] *Layer
	LayerList []*Layer
	diversor *Diversor
	id uint64
}

func (domain *Domain)Init(name string, id uint64)  {
	domain.name = name
	domain.id = id
}

func (domain *Domain)GetName() string  {
	return domain.name
}
func (domain *Domain)SetName(name string)  {
	domain.name = name
}
func (domain *Domain)InsertLayer(layer *Layer)  {
	domain.LayerMapper[layer.GetId()] = layer
	domain.LayerList = append(domain.LayerList, layer)
}
func (domain *Domain)GetId() uint64  {
	return domain.id
}

func (domain *Domain)SetId(id uint64)  {
	domain.id = id
}

func (domain *Domain)GetDiversor() *Diversor {
	return domain.diversor
}
func (domain *Domain)SetDiversor(diversor Diversor)  {
	domain.diversor = &diversor
}
func (domain *Domain)HashTraffic(traffic Traffic)  {

}