package Components

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
