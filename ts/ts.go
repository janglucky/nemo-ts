package ts

import (
	"errors"
	"github.com/janglucky/nemo-ts/ts/converter/converters"
	"github.com/janglucky/nemo-ts/ts/service"
)

func Call(serviceName string, request interface{}, response interface{}, converterType string)  error{
	s, ok := service.GetService(serviceName)
	if !ok {
		err := errors.New("can't find service name " + serviceName + " or load services failed")
		return err
	}

	options := NewCustomRequestOption()
	do(s, request, response, options)

	return nil
}

func do(conf *service.ServiceConfig, req interface{}, rsp interface{}, options *CustomRequestOption)  {
	
}

func NewCustomRequestOption() *CustomRequestOption{
	return &CustomRequestOption{
		Retry: -1,
		Converters: []converters.ConverterType{converters.EMPTY, converters.EMPTY},
	}
}