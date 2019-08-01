package Components

/*
Parameter：最小的控制单元
接口调用形式，必须实现toString()方法
Parameter实行插件式管理，每次实现一个Parameter的继承对象，都会自动加入到对应的资源组中
Parameter的具体实现要根据不同的业务进行单独开发
*/
type Parameter interface {
	toString() []byte
}
