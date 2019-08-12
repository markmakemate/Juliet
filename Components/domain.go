package Components

/**
Domain实现
*/
func (domain *Domain) Init(name string, id uint64) {
	domain.Name = name
	domain.Id = id
	domain.LayerIdPool = []uint64{}
}

func (domain *Domain) ReceiveTraffic(traffic *Traffic) {
	/**
	TODO
	对traffic进行domain层分流
	*/

}
func (domain *Domain) Diverse() {

}

func (domain *Domain) Serialize() {

}
func (domain *Domain) DeleteLayer(layerId uint64) {
	domain.LayerIdPool = DeleteFromUint64Array(domain.LayerIdPool, layerId)
}

func (domain *Domain) InsertLayer(layer *Layer) {
	domain.LayerIdPool = append(domain.LayerIdPool, layer.Id)
}
