package ts

import "github.com/janglucky/nemo-ts/ts/converter/converters"

var JSONConverter = converters.JSON
var FORMConverter = converters.FORM
var MCPACK1Converter = converters.MCPACK1
var MCPACK2Converter = converters.MCPACK2
var STRINGConverter = converters.STRING
var RAWConverter = converters.RAW
var REDISConverter = converters.REDIS
var MYSQLConverter = converters.MYSQL
var TLVConverter = converters.TLV

// CustomRequestOption ... custom request option
type CustomRequestOption struct {
	Converters     []converters.ConverterType
	Retry          int
	RetryCheckData interface{}
}
