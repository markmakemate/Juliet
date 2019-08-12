package Managers

import (
	"caibeike-abtest/Components"
	"caibeike-abtest/Utils"
	"fmt"
	"net/http"
	"strconv"
)

/**
统一的资源管理器，管理一个业务组所有实验资源的元数据，并完成流量分流
1. 资源管理器的两个作用：异步接受连接管理器传来的流量和对流量分流
2. 资源管理器连同其他组件会被打包在镜像中，由Docker来发布，暴露给业务层一组RESTful API服务
3. 分流可异步执行，不论如何分流，顺序如何，都会将traffic推入一个全局FIFO队列，最后由连接管理器指示Traffic 调用Send()接口发送系统参数, 这两个过程并发执行
4. 资源管理器不管理Traffic，其在作为参数传进来时已经被封装好了，Traffic在服务器启动时监听端口，获取连接后由连接管理器管理。
资源管理器只会将Traffic推入某个Domain，由Domain分流。
5. 资源管理器不管理连接实例，连接会被封装进Traffic里，资源管理器只管理动态的实验资源
6. 资源管理器在实验平台启动时，将持久化的实验资源加载到内存中。用户在实验前台做的所有改动操作都会先在内存中做相应的改动，之后持久化到数据库中。这样
的优点在于分流和实验共用一个平台，实时更新、加快查询速度且保证并发量，添加新的实验资源不需要前端更新代码。
*/

func (rm *ResourceManager) Init(uuid string) {
	rm.uuid = uuid
}

func (rm *ResourceManager) InsertDomain(domain *Components.Domain) {
	Components.ConfigurationOfContainerManager.DomainContainer.Inject(domain)
}

func (rm *ResourceManager) DeleteDomain(domainId uint64) {
	Components.ConfigurationOfContainerManager.DomainContainer.Eject(domainId)
}

func (rm *ResourceManager) InsertLayerInDomain(domainId uint64, layer *Components.Layer) {
	Components.ConfigurationOfContainerManager.DomainContainer.Get(domainId).InsertLayer(layer)
}
func (rm *ResourceManager) ReceiveTraffic(traffic Components.Traffic) {
	if traffic.DomainId != -1 {
		Components.ConfigurationOfContainerManager.DomainContainer.Get(traffic.DomainId).ReceiveTraffic(&traffic)
	} else if traffic.DomainId != -1 && traffic.LayerId != -1 {
		if Utils.IsExist(Components.ConfigurationOfContainerManager.DomainContainer.Get(traffic.DomainId).LayerIdPool,
			traffic.LayerId) {
			/**
			TODO
			Layer做分流
			*/
			Components.ConfigurationOfContainerManager.LayerContainer.Get(traffic.LayerId).ReceiveTraffic(&traffic)
		}

	} else {
		id := Diversor.Hash(traffic)
		/**
		TODO
		Domain做分流，传TrafficQueue全局队列为参数，push进该队列中
		*/
	}
}

/**
1. 连接管理器，用来管理连接实例。将流量打包成Traffic。
2. 连接管理器不参与实验资源的管理，只管理流量资源
3. 连接管理器同时管理一个连接channel和一个缓冲区，连接打包推入smPool，当smPool满时阻塞，则推入缓冲区。缓冲区定时查看smPool
4. 连接管理器管理一个动态的全局Traffic队列，并通过信号指示Traffic发送数据
*/

func (sm *SessionManager) Pack(w http.ResponseWriter, r *http.Request) {
	go func() {
		domainId := uint64(-1)
		layerId := uint64(-1)
		vars := r.URL.Query()
		if vars["domainId"][0] != "" {
			dId, errDomainId := strconv.ParseUint(vars["domainId"][0], 10, 64)
			domainId = dId
			Utils.Log(errDomainId)
		}
		if vars["layerId"][0] != "" {
			lId, errLayerId := strconv.ParseUint(vars["layerId"][0], 10, 64)
			layerId = lId
			Utils.Log(errLayerId)
		}
		traffic := new(Components.Traffic)
		traffic.Init(w, r, domainId, layerId)
		if len(SessionPool) <= cap(SessionPool) {
			SessionPool <- traffic
		} else if len(Cache) <= cap(Cache) {
			Cache <- traffic
		} else {
			n, _ := fmt.Fprint(w, 500)
		}
	}()
}

func (sm *SessionManager) Send(signal chan int) {
	for {
		x, ok := <-TrafficQueue
		if ok {
			go x.Send(signal)
		}
	}
}

/**
web管理器实现
*/
