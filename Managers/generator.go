package Managers

/**
所有的生成器方法
*/
func NewResourceManager(uuid string) *ResourceManager {
	return &ResourceManager{
		uuid: uuid,
	}
}

func NewSessionManager(uuid string) *SessionManager {
	return &SessionManager{uuid: uuid}
}
