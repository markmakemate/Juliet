package Managers

import (
	"caibeike-abtest/Components"
	"caibeike-abtest/Container"
)

var (
	TrafficQueue chan *Components.Traffic
	Diversor     *Components.Diversor
	SessionPool  chan *Components.Traffic
	Cache        chan *Components.Traffic
	Config       ConfigOfContainerManager
)

//The definition of Resource Manager
type ResourceManager struct {
	uuid             string
	ContainerManager *ContainerManager
}

//The definition of Session Manager
type SessionManager struct {
	uuid string
}

//The definition of Container Manager's configuration
type ConfigOfContainerManager struct {
	uuid            string
	DomainContainer *Container.DomainContainer
	ExptContainer   *Container.ExperimentContainer
	LayerContainer  *Container.LayerContainer
}

//The definition of container manager
type ContainerManager struct {
	uuid                string
	DomainContainer     *Container.DomainContainer
	ExperimentContainer *Container.ExperimentContainer
	LayerContainer      *Container.LayerContainer
}

//The definition of web manager
type WebManager struct {
	uuid string
}
