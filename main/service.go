package main

import (
	"Juliet/Utils"
	"Juliet/common"
	"log"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
)


/**
Managers definition of Juliet A/B test platform

Version 0.0.1-SNAPSHOT

Data: 2019.8

@author: Song Yi

Contact: yi.song@caibeike.com

Copyright: CaiBeiKe
*/

/**
The definition of Session Manager
1. 连接管理器，用来管理连接实例。将流量打包成 Traffic；
2. 连接管理器不参与实验资源的管理，只管理流量资源；
3. 连接管理器同时管理一个连接 channel 和一个缓冲区，连接打包推入 sessionPool，当 sessionPool 满时阻塞，则推入缓冲区。缓冲区定时查看 sessionPool；
4. 连接管理器管理一个动态的全局 Traffic 队列，并通过信号指示 Traffic 发送数据；
*/
type SessionManager struct {
	addr string
}


func NewSessionManager(addr string) *SessionManager {
	return &SessionManager{
		addr: addr,
	}
}

var (
	trafficQueue = make(chan *common.Traffic, 1000)
)

func diverse(traffic common.Traffic) {
	if traffic.LayerId != 0 {
		/**
		TODO
		Layer 做分流
		*/
		traffic = common.DomainEntity.Lpc.Get(traffic.LayerId).ReceiveTraffic(&traffic)
		push(&traffic)
	} //不存在需返回一个默认参数
}

/**
Receive traffic from session manager and do the split-flow progress
*/


//Push new parameter type in ExpParamContainer
func push(traffic *common.Traffic)  {
	trafficQueue <- traffic
}

/**
Session Manager all methods implementation
*/

/**
Pack is the method to package the request from client in the form of Traffic
:param w httpResponseWriter
:param r httpRequestPtr
*/


func (sm *SessionManager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	traffic := pack(w, r)
	go diverse(*traffic)
	go send()
}

/**
Send() method is to control the elements in TrafficQueue to send the experiment to clients
:param signal, the signal decides when to send experiments. When signal channel receives 1, TrafficQueue will send the ar
*/
func send() {
	if x, ok := <-trafficQueue; ok {
		x.Send()
	}
	//优化点
}

func pack(w http.ResponseWriter, r *http.Request) *common.Traffic {
	traffic := &common.Traffic{
		Writer:   w,
		Request:  r,
		DomainId: 0,
		LayerId:  0,
		ExpList: []common.Exp{},
		Mutex:    sync.Mutex{},
	}
	vars := r.URL.Query()
	if vars["domainId"][0] != "" {
		dId, errDomainId := strconv.ParseUint(vars["domainId"][0], 10, 64)
		Utils.Log(errDomainId)
		traffic.DomainId = atomic.LoadUint64(&dId)
	}
	if vars["layerId"][0] != "" {
		lId, errLayerId := strconv.ParseUint(vars["layerId"][0], 10, 64)
		Utils.Log(errLayerId)
		traffic.LayerId = atomic.LoadUint64(&lId)
	}
	return traffic
}

func (sm *SessionManager) Start(addr string)  {
	for{
		err := http.ListenAndServe(addr, sm)
		if  err != nil{
			log.Print(err)
		}
	}
}




