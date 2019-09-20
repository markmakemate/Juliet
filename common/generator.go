package common

import (
	"net/http"
	"sync"
)

/**
所有的生成器的实现
*/

/**
Domain 生成器
*/
func NewDomain(id uint64) *Domain {
	return &Domain{
		Id:          id,
	}

}

/**
Layer 生成器
*/
func NewLayer(layerId uint64) *Layer {
	return &Layer{
		Id:         layerId,
		Mutex:      sync.Mutex{},
	}
}

/**
Experiment生成器
*/
func NewExp(expId uint64,
	config ExpIdConfig) *Exp {
	return &Exp{
		Id:        expId,
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
		ExpList:     nil,
		Mutex:    sync.Mutex{},
	}
}

func NewDiversor(base uint64) *Diversor {
	return &Diversor{
		mutex:    sync.Mutex{},
		features: &feature{},
		Base:     base,
	}
}