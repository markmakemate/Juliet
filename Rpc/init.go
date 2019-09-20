package Rpc

type Domain struct {
	Name string
	DomainId uint64
}

type Data struct {
	Obj string
	Id uint64
}

type Layer struct {
	LayerId uint64
	Name string
}

type Type struct {
	Type string
}

type Result struct {
	Result string
}

type Parameter struct {
	Param string
}

var PlatformServer *Server

func init()  {
	PlatformServer = NewRpcServer()
}
