package Managers

import (
	"caibeike-abtest/Components"
	"caibeike-abtest/Utils"
	"fmt"
	"net/http"
	"strconv"
)

/**
1. 连接管理器，用来管理连接实例。将流量打包成Traffic。
2. 连接管理器不参与实验资源的管理，只管理流量资源
3. 连接管理器同时管理一个连接channel和一个缓冲区，连接打包推入smPool，当smPool满时阻塞，则推入缓冲区。缓冲区定时查看smPool
4. 连接管理器管理一个动态的全局Traffic队列，并通过信号指示Traffic发送数据
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

func (sm *SessionManager) Send(signal chan int) {
	for {
		x, ok := <-TrafficQueue
		if ok {
			go x.Send(signal)
		}
	}
}
