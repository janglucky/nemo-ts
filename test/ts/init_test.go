package ts

import (
	"fmt"
	"github.com/janglucky/nemo-ts/ts"
	"github.com/janglucky/nemo-ts/ts/service"
	"testing"
)

func TestInit(t *testing.T)  {
	ts.InitRal("dev")
	service.ServiceMap.Range(func(key, value interface{}) bool {
		fmt.Printf("%+v,%+v", key, value)
		return true
	})

}