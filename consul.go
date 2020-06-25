package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"os"
	"strconv"
	"strings"
	"time"
)

//go:generate kratos tool protoc --grpc consul.proto

func (p *ServiceConsul) ServiceRegister(consulConfig *api.Config) {
	if consulConfig == nil {
		consulConfig = api.DefaultConfig()
	}
	//register consul

	client, err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Printf("NewClient error\n%v", err)
		return
	}
	agent := client.Agent()
	var interval time.Duration
	if p.Interval == 0 {
		interval = time.Duration(10) * time.Second
	} else {
		interval = time.Duration(p.Interval) * time.Second
	}

	deregister := time.Duration(p.Deregister) * time.Minute

	server := new(api.AgentServiceCheck)
	if strings.ToLower(p.ServerType) == "grpc" {
		server = &api.AgentServiceCheck{ // 健康检查
			Interval:                       interval.String(),                             // 健康检查间隔
			GRPC:                           fmt.Sprintf("%v:%v/%v", p.Ip, p.Port, p.CheckName), // grpc 支持，执行健康检查的地址，service 会传到 Health.Check 函数中
			DeregisterCriticalServiceAfter: deregister.String(),                           // 注销时间，相当于过期时间
		}
	} else {
		server = &api.AgentServiceCheck{ // 健康检查
			Interval:                       interval.String(),                             // 健康检查间隔
			HTTP:                           fmt.Sprintf("%v:%v/%v", p.Ip, p.Port, p.CheckName), // http 支持，执行健康检查的地址，service 会传到 Health.Check 函数中
			DeregisterCriticalServiceAfter: deregister.String(),                           // 注销时间，相当于过期时间
		}
	}
	reg := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%v-%v-%v", p.ServerName, p.Ip, p.Port), // 服务节点的名称
		Name:    p.ServerName,                                        // 服务名称
		Tags:    p.Tags,                                        // tag，可以为空
		Port:    int(p.Port),                                   // 服务端口
		Address: p.Ip,                                          // 服务 IP
		Check:   server,
	}
	if err := agent.ServiceRegister(reg); err != nil {
		fmt.Printf("Service Register error\n%v", err)
		return
	}
}

func DockerCreateServiceConsul(ipEnv, portEnv, serverName, checkName, serverType string, tags... string) (res *ServiceConsul, err error) {
	res = &ServiceConsul{}
	ip := os.Getenv(ipEnv)
	portStr := os.Getenv(portEnv)
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return res, err
	}
	res.Port = int64(port)
	res.Ip = ip
	res.ServerName = serverName
	res.CheckName = checkName
	res.ServerType = serverType
	res.Tags = tags
	return res, err
}

func CreateServiceConsul(ip string, port int, serverName, checkName, serverType string, tags... string) (res *ServiceConsul, err error) {
	res = &ServiceConsul{}
	res.Port = int64(port)
	res.Ip = ip
	res.ServerName = serverName
	res.CheckName = checkName
	res.ServerType = serverType
	res.Tags = tags
	return res, err
}