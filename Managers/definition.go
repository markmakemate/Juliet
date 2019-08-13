package Managers

import (
	"caibeike-abtest/Components"
)

//必要的全局变量，在服务启动前需要完成配置
var (
	TrafficQueue chan *Components.Traffic
	Diversor     *Components.Diversor
	SessionPool  chan *Components.Traffic
	Cache        chan *Components.Traffic
	Config       Components.ConfigOfContainerManager
)

/**
The definition of Resource Manager
统一的资源管理器，管理一个业务组所有实验资源的元数据，并完成流量分流；
1. 资源管理器的两个作用：(1) 异步接受连接管理器传来的流量和对流量分流；(2) 管理动态的实验资源；
2. 资源管理器连同其他组件会被打包在镜像中，由 Docker 来发布，暴露给业务层一组 RESTful API 服务；
3. 分流可异步执行，不论如何分流，顺序如何，都会将 traffic 推入一个全局 FIFO 队列，最后由连接管理器指示 Traffic 调用 Send() 接口发送系统参数, 这两个过程并发执行；
4. 资源管理器不管理 Traffic，其在作为参数传进来时已经被封装好了，Traffic 在服务器启动时监听端口，获取连接后由连接管理器管理；
资源管理器只会将 Traffic 推入某个Domain或者 Layer，由 Domain 和 Layer 分流；
5. 资源管理器不管理连接实例，连接会被封装进 Traffic 里，资源管理器只管理动态的实验资源；
6. 资源管理器在实验平台启动时，将本地持久化的实验资源加载到内存中并从通过 web 管理器从实验平台后台根据日志文件的时间戳拉取最新的实验数据；
用户对实验资源做的所有改动，实验平台后台通过 dubbo-rpc 调用 web 管理器方法，先存储在内存中，然后调用 Serialize() 方法持久化到本地(惰性操作)，
优点在于实时更新、加快请求速度、保证并发量，添加新的实验资源类型仅需要前端实现添加该类型的代码；
7. 在实验平台宕机后，测试平台依旧不受影响，当测试平台宕机后，业务系统可以根据默认业务规则服务用户，实验平台可以添加新实验等待测试平台重新启动；
*/
type ResourceManager struct {
	uuid string
}

/**
The definition of Session Manager
1. 连接管理器，用来管理连接实例。将流量打包成 Traffic；
2. 连接管理器不参与实验资源的管理，只管理流量资源；
3. 连接管理器同时管理一个连接 channel 和一个缓冲区，连接打包推入 sessionPool，当 sessionPool 满时阻塞，则推入缓冲区。缓冲区定时查看 sessionPool；
4. 连接管理器管理一个动态的全局 Traffic 队列，并通过信号指示 Traffic 发送数据；
*/
type SessionManager struct {
	uuid string
}

/**
The definition of web manager
Web管理器，管理与实验平台后台的交互，实现实验资源动态调度，所有方法都采用 dubbo-rpc 远程让后台调用
*/
type WebManager struct {
	uuid string
}
