package consul

import (
	"errors"
	"fmt"
	"github.com/hashicorp/consul/api"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

//go:generate kratos tool protoc --grpc consul.proto

func (p *ServiceConsul) ServiceRegister(consulConfig *api.Config) (err error) {
	if consulConfig == nil {
		consulConfig = api.DefaultConfig()
	}

	//register consul

	client, err := api.NewClient(consulConfig)
	if err != nil {
		return errors.New(fmt.Sprintf("NewClient error\n%v", err))
	}
	agent := client.Agent()
	var interval time.Duration
	if p.Interval == 0 {
		interval = time.Duration(10) * time.Second
	} else {
		interval = time.Duration(p.Interval) * time.Second
	}
	var deregister time.Duration
	if p.Deregister == 0 {
		deregister = time.Duration(365 * 24 * 20) * time.Hour
	} else {
		deregister = time.Duration(p.Deregister) * time.Minute
	}


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
			HTTP:                           fmt.Sprintf("%s://%v:%v/%v", p.ServerType, p.Ip, p.Port, p.CheckName), // http 支持，执行健康检查的地址，service 会传到 Health.Check 函数中
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
		return errors.New(fmt.Sprintf("Service Register error\n%v", err))
	}
	return err
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

func LocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	panic("local ip is empty")
}

func ExtranetIP() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}