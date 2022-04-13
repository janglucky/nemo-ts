package finders

type NodeInfo struct {
	NodeId string `json:"node_id"`
	ServiceName string `json:"service_name"`
	Ip string `json:"ip"`
	Port int `json:"port"`
}