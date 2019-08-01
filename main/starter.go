package main

import "ABTest/Managers"

/**
服务启动流程：
1. 实例化资源管理器、连接管理器和分流器
2. 启动连接管理器的监听器，监听指定地址和端口
3. 分流
*/
func main() {
	resourceManager := new(Managers.ResourceManager)
	sessionManager := new(Managers.SessionManager)
}
