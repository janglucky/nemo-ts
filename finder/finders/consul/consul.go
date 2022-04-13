package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/janglucky/nemo-ts/finder/finders"
)

const (
	CONSUL_ADDR = "127.0.0.1:8500"
)

type Consul struct {
}

func Call(req interface{})  {
}

func (this *Consul) GetInstances( serviceName string) []finders.NodeInfo{
	var lastIndex uint64
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500" //consul server

	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println("api new client is failed, err:", err)
		return []finders.NodeInfo{}
	}
	services, metainfo, err := client.Health().Service("nginx", "v1000", true, &api.QueryOptions{
		WaitIndex: lastIndex, // 同步点，这个调用将一直阻塞，直到有新的更新
	})
	if err != nil {
		fmt.Println(err.Error())
		return []finders.NodeInfo{}
		//logrus.Warn("error retrieving instances from Consul: %v", err)
	}
	lastIndex = metainfo.LastIndex

	var instances []finders.NodeInfo
	for _, service := range services {
		instances = append(instances, finders.NodeInfo{
			NodeId: service.Service.ID,
			ServiceName: serviceName,
			Ip: service.Service.Address,
			Port: service.Service.Port,
		})
	}
	return instances
}

func init()  {
	finders.Finders["consul"] = Consul{}
}