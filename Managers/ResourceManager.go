package Managers

import (
	"ABTest/Components"
	"net/http"
)

/*
统一的资源管理器，管理一个业务组所有实验资源的元数据，并完成流量分流
1. 资源管理器的两个作用：异步接受前台传来的流量和流量分流
2. 资源管理器连同组件会被打包在镜像中，由Docker来发布，暴露给业务层一组RESTful API服务
3. 分流可异步执行，不论如何分流，顺序如何，都会将traffic推入一个全局FIFO队列，最后由连接管理器指示Traffic Send系统参数
4. 资源管理器不管理Traffic，其在作为参数传进来时已经被封装好了，这是服务器启动时监听端口，获取连接后由连接管理器打包传来。
资源管理器只会将Traffic推入某个Domain，由Domain来存储（应对较大流量时，需要缓冲队列）并分流。
5. 资源管理器不管理连接实例，连接会被封装进Traffic里，资源管理器只管理动态的实验资源
*/
type ResourceManager struct {
	Name         string
	Id           uint64
	Server       *http.Server
	DomainList   []*Components.Domain
	DomainMapper map[uint64]*Components.Domain
	queue        chan *Components.Traffic
	Diversor     *Components.Diversor
}

func (resourceManager *ResourceManager) Init(Server *http.Server, queue chan *Components.Traffic) {
	resourceManager.Server = Server
	resourceManager.Diversor = new(Components.Diversor)
	resourceManager.queue = queue
}

func (resourceManager *ResourceManager) InsertDomain(domain *Components.Domain) {
	resourceManager.DomainList = append(resourceManager.DomainList, domain)
	resourceManager.DomainMapper[domain.Id] = domain
}
func (resourceManager *ResourceManager) InsertExptInDomain(domainId uint64, expt *Components.Experiment) {
	resourceManager.DomainMapper[domainId].InsertExpt(expt)
}
func (resourceManager *ResourceManager) ReceiveTraffic(traffic Components.Traffic) {
	if traffic.DomainId != -1 {
		resourceManager.DomainMapper[traffic.DomainId].ReceiveTraffic(&traffic)
	} else if traffic.DomainId != -1 && traffic.LayerId != -1 {
		resourceManager.DomainMapper[traffic.DomainId].LayerMapper[traffic.LayerId].ReceiveTraffic(&traffic)
		/**
		TODO
		Layer做分流
		*/
	} else {
		id := resourceManager.Diversor.Hash(traffic)
		resourceManager.DomainList[id].ReceiveTraffic(&traffic)
		/**
		TODO
		Domain做分流，传queue全局队列为参数，push进该队列中
		*/
	}
}
func (resourceManager *ResourceManager) Diverse() Components.Experiment {
	expt := new(Components.Experiment)
	return *expt
}
