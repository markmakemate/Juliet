package Managers

import (
	"caibeike-abtest/Components"
)

var (
	TrafficQueue chan *Components.Traffic
	Diversor     *Components.Diversor
	SessionPool  chan *Components.Traffic
	Cache        chan *Components.Traffic
	Config       Components.ConfigOfContainerManager
)

//The definition of Resource Manager
type ResourceManager struct {
	uuid string
}

//The definition of Session Manager
type SessionManager struct {
	uuid string
}

//The definition of web manager
type WebManager struct {
	uuid string
}
