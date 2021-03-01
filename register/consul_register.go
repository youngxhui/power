package register

import (
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/go-uuid"
	"github.com/youngxhui/power/log"
)

type Consul struct {
	Port    int
	Address string
}

func (c Consul) RegisterServer(ServiceName string, Port int, Address string) {
	log.Info("开始注册服务")

	client, _ := api.NewClient(api.DefaultConfig())
	register := new(api.AgentServiceRegistration)
	uuid, _ := uuid.GenerateUUID()
	register.ID = uuid
	register.Port = Port
	register.Name = ServiceName
	register.Address = Address
	register.Tags = []string{"Power"}
	check := &api.AgentServiceCheck{
		HTTP:     "http://127.0.0.1:1234/health",
		Timeout:  "5s",
		Interval: "5s",
	}
	register.Check = check
	err := client.Agent().ServiceRegister(register)

	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
}

// UnRegisterServer 注销服务
func (c Consul) UnRegisterServer(ServerName string) {
	log.Info("反注册服务")
}
