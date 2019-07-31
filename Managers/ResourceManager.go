package Managers

import (
	"ABTest/Components"
	"net/http"
)
type ResourceManager struct {
	Server *http.Server
	DomainList []*Components.Domain
	diversor *Components.Diversor
}

func (resourceManager *ResourceManager) Init(Server *http.Server) {
	resourceManager.Server = Server
}

func (resourceManager *ResourceManager) InsertDomain(domain *Components.Domain) {
	resourceManager.DomainList = append(resourceManager.DomainList, domain)
}

func (resourceManager *ResourceManager) SetServer(server *http.Server) {
	resourceManager.Server = server
}

func (resourceManager *ResourceManager) GetServer() http.Server{
	return *resourceManager.Server
}

func (resourceManager *ResourceManager) ListenAndDiverse() {

}