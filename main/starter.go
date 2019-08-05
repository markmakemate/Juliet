package main

import "caibeike-abtest/Managers"

/**
服务启动流程：
1. 实例化资源管理器、连接管理器和分流器
2. 将持久化的实验资源数据加载到内存中
3. 启动连接管理器的监听器，监听指定地址和端口
*/
func main() {
	resourceManager := new(Managers.ResourceManager)
	sessionManager := new(Managers.SessionManager)
}
