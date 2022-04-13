package service

func GetService(serviceName string)  (*ServiceConfig, bool){
	v, ok := ServiceMap.Load(serviceName)
	if !ok {
		return nil, false
	}

	conf, ok := v.(*ServiceConfig)
	if !ok {
		return nil, false
	}

	return conf, true
}