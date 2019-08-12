package Managers

/**
所有的生成器方法
*/
func NewResourceManager(uuid string, cm *ContainerManager) *ResourceManager {
	return &ResourceManager{
		uuid:             uuid,
		ContainerManager: cm,
	}
}

func NewSessionManager(uuid string) *SessionManager {
	return &SessionManager{uuid: uuid}
}

func NewContainerManager() *ContainerManager {
	return &ContainerManager{
		uuid:                Config.uuid,
		DomainContainer:     Config.DomainContainer,
		ExperimentContainer: Config.ExptContainer,
		LayerContainer:      Config.LayerContainer,
	}
}
