package Managers

import (
	"ABTest/Components"
	"net/http"
)

/**
连接管理器，用来管理连接实例。将流量打包成Traffic随机分配给这些子进程。
*/

type SessionManager struct {
	NumCPU           int
	resourceManagers chan *ResourceManager
}

func (session *SessionManager) Init(NumCPU int) {
	session.NumCPU = NumCPU
	session.resourceManagers = make(chan *ResourceManager, NumCPU)
}
func (session *SessionManager) Start(w http.ResponseWriter, r *http.Request) {
	traffic := new(Components.Traffic)
	traffic.Request = *r
	traffic.Writer = w
	x := <-session.resourceManagers
	go x.ReceiveTraffic(*traffic)
}
