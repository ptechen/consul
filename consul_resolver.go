package consul

import (
	"errors"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc/resolver"
	"regexp"
	"sync"
)

const (
	defaultPort = "8500"
)

var (
	errMissingAddr = errors.New("consul resolver: missing address")

	errAddrMisMatch = errors.New("consul resolver: invalied uri")

	errEndsWithColon = errors.New("consul resolver: missing port after port-separator colon")

	regexConsul, _ = regexp.Compile("^([A-z0-9.]+)(:[0-9]{1,5})?/([A-z_]+)$")
)

func Init() {
	resolver.Register(NewBuilder())
}

type consulBuilder struct {
}

type consulResolver struct {
	address              string
	wg                   sync.WaitGroup
	cc                   resolver.ClientConn
	name                 string
	disableServiceConfig bool
	lastIndex            uint64
}

func NewBuilder() resolver.Builder {
	return &consulBuilder{}
}

func (cb *consulBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	host, port, name, err := parseTarget(fmt.Sprintf("%s/%s", target.Authority, target.Endpoint))
	if err != nil {
		return nil, err
	}

	cr := &consulResolver{
		address:              fmt.Sprintf("%s%s", host, port),
		name:                 name,
		cc:                   cc,
		disableServiceConfig: opts.DisableServiceConfig,
		lastIndex:            0,
	}

	cr.wg.Add(1)
	go cr.watcher()
	return cr, nil

}

func (cr *consulResolver) watcher() {
	config := api.DefaultConfig()
	config.Address = cr.address
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Printf("error create consul client: %v\n", err)
		return
	}

	for {
		services, metainfo, err := client.Health().Service(cr.name, cr.name, true, &api.QueryOptions{WaitIndex: cr.lastIndex})
		if err != nil {
			fmt.Printf("error retrieving instances from Consul: %v", err)
		}

		cr.lastIndex = metainfo.LastIndex

		state := resolver.State{
			Addresses: []resolver.Address{},
		}
		for _, service := range services {
			addr := fmt.Sprintf("%v:%v", service.Service.Address, service.Service.Port)
			state.Addresses = append(state.Addresses, resolver.Address{Addr: addr})
		}
		fmt.Printf("newAddrs: %v\n", state.Addresses)
		cr.cc.UpdateState(state)
	}

}

func (cb *consulBuilder) Scheme() string {
	return "consul"
}

func (cr *consulResolver) ResolveNow(opt resolver.ResolveNowOptions) {
}

func (cr *consulResolver) Close() {
}

func parseTarget(target string) (host, port, name string, err error) {

	if target == "" {
		return "", "", "", errMissingAddr
	}

	if !regexConsul.MatchString(target) {
		return "", "", "", errAddrMisMatch
	}

	groups := regexConsul.FindStringSubmatch(target)
	host = groups[1]
	port = groups[2]
	name = groups[3]
	if port == "" {
		port = defaultPort
	}
	return host, port, name, nil
}
