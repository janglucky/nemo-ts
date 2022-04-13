package service

import "sync"

var ServiceMap sync.Map

type ServiceConfig struct {
	Name         string           `toml:"Name"`
	ConnTimeOut  int              `toml:"ConnTimeOut"`
	WriteTimeOut int              `toml:"WriteTimeOut"`
	ReadTimeOut  int              `toml:"ReadTimeOut"`
	Retry        int              `toml:"Retry"`
	Protocol     string           `toml:"Protocol"`
	Converter    string           `toml:"Converter"`
	Strategy     string           `toml:"Strategy"`
	Finder       string           `toml:"Finder"`
	Resources    []ResourceConfig `toml:"Resources"`
}

type ResourceConfig struct {
	Host string `toml:"Host"`
	IP   string `toml:"Ip"`
	Port int    `toml:"Port"`
}
