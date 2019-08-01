package Managers

import (
	"ABTest/Components"
	"net/http"
)

/**
连接管理器，用来管理连接实例
*/

type SessionManager struct {
	queue chan *Components.Traffic
}

func (session *SessionManager) Listen(addr string) {

}
func (session *SessionManager) HandlerFunc(addr string,
	handler func(w http.ResponseWriter, r *http.Request)) {

}
func (session *SessionManager) Start() {

}
