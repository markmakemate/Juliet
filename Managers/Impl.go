package Managers

import (
	"caibeike-abtest/Components"
	"caibeike-abtest/Utils"
	"fmt"
	"net/http"
	"strconv"
)

/**
Manager 实现
*/

/**
Resource Manager all methods implementation
*/
func (rm *ResourceManager) Init(uuid string) {
	rm.uuid = uuid
}

//Insert domain with domainId in DomainContainer
func (rm *ResourceManager) InsertDomain(domain *Components.Domain) {
	Components.ConfigurationOfContainerManager.DomainContainer.Inject(domain)
}

//Delete domain with domainId from DomainContainer
func (rm *ResourceManager) DeleteDomain(domainId uint64) {
	Components.ConfigurationOfContainerManager.DomainContainer.Eject(domainId)
}

//Insert layer with layerId in domain with domainId and LayerContainer
func (rm *ResourceManager) InsertLayerInDomain(domainId uint64, layer *Components.Layer) {
	Components.ConfigurationOfContainerManager.DomainContainer.Get(domainId).InsertLayer(layer)
}

/**
Receive traffic from session manager and do the split-flow progress
*/
func (rm *ResourceManager) ReceiveTraffic(traffic Components.Traffic) {
	if traffic.DomainId != -1 {
		traffic = Components.ConfigurationOfContainerManager.DomainContainer.Get(traffic.DomainId).ReceiveTraffic(&traffic)
		TrafficQueue <- &traffic
	} else if traffic.DomainId != -1 && traffic.LayerId != -1 {
		if Utils.IsExist(Components.ConfigurationOfContainerManager.DomainContainer.Get(traffic.DomainId).LayerIdPool,
			traffic.LayerId) {
			/**
			TODO
			Layer 做分流
			*/
			traffic = Components.ConfigurationOfContainerManager.LayerContainer.Get(traffic.LayerId).ReceiveTraffic(&traffic)
			TrafficQueue <- &traffic
		} //不存在需返回一个默认参数

	} else {
		id := Diversor.Hash(traffic) //当流量不指定 domain 和 layer 时随机分配一个 domain
		/**
		TODO
		Domain 做分流，push 进 TrafficQueue 全局队列中
		*/
		traffic = Components.ConfigurationOfContainerManager.DomainContainer.Get(id).ReceiveTraffic(&traffic)
		TrafficQueue <- &traffic
	}
}

//Push new parameter type in ParameterContainer
func (rm *ResourceManager) PushNewParameter(param *Components.Parameter, Type string) {
	Components.ConfigurationOfContainerManager.ParameterContainer.Inject(param, Type)
}

/**
Session Manager all methods implementation
*/
func (sm *SessionManager) Init(uuid string) {
	sm.uuid = uuid
}

/**
Pack is the method to package the request from client in the form of Traffic
:param w httpResponseWriter
:param r httpRequestPtr
*/
func (sm *SessionManager) Pack(w http.ResponseWriter, r *http.Request) {
	go func() {
		domainId := uint64(-1)
		layerId := uint64(-1)
		vars := r.URL.Query()
		if vars["domainId"][0] != "" {
			dId, errDomainId := strconv.ParseUint(vars["domainId"][0], 10, 64)
			domainId = dId
			Utils.Log(errDomainId)
		}
		if vars["layerId"][0] != "" {
			lId, errLayerId := strconv.ParseUint(vars["layerId"][0], 10, 64)
			layerId = lId
			Utils.Log(errLayerId)
		}
		traffic := new(Components.Traffic)
		traffic.Init(w, r, domainId, layerId)
		if len(SessionPool) <= cap(SessionPool) {
			SessionPool <- traffic
		} else if len(Cache) <= cap(Cache) {
			Cache <- traffic
		} else {
			n, _ := fmt.Fprint(w, 500)
		}
	}()
}

/**
Send() method is to control the elements in TrafficQueue to send the experiment to clients
:param signal, the signal decides when to send experiments. When signal channel receives 1, TrafficQueue will send the args
*/
func (sm *SessionManager) Send(signal chan int) {
	for {
		x, ok := <-TrafficQueue
		if ok {
			go x.Send(signal)
		}
	}
}

/**
Web Manager all methods implementation
*/
