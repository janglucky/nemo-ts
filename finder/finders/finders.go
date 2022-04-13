package finders

import (
	_ "github.com/janglucky/nemo-ts/finder/finders/bns"
	_ "github.com/janglucky/nemo-ts/finder/finders/consul"
	_ "github.com/janglucky/nemo-ts/finder/finders/zookeeper"
)

var Finders = make(map[string]FindersInter)

type FindersInter interface {}
