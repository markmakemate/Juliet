package main

import (
	"encoding/json"
	"fmt"
)

/**
服务启动流程：
1. 实例化资源管理器、连接管理器和分流器
2. 将持久化的实验资源数据加载到内存中
3. 启动连接管理器的监听器，监听指定地址和端口
AB测试平台只维护少量的实验资源的元数据信息，大部分都是通过RPC调用或者HTTP请求RESTful接口完成参数请求。数据拉取的过程都是由业务组自定义的
*/
type Arg interface {
	toString() []byte
}
type Add struct {
	A int
	B int
}

func (add *Add) toString() []byte {
	jsonstr, err := json.Marshal(add)
	if err != nil {
		fmt.Println(err.Error())
	}
	return jsonstr
}
func main() {
	var a []Arg
	a = append(a, &Add{1, 2})
	a = append(a, &Add{3, 4})
	fmt.Println(string(a[0].toString()))
}
