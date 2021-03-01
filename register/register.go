package register

// Register 服务注册接口
type Register interface {
	// RegisterServer 注册服务
	RegisterServer(ServiceName string,Port int,Address string)
	// UnRegisterServer 注销服务
	UnRegisterServer(ServerName string)
}
