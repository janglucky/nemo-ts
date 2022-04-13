package ts

import (
	"github.com/janglucky/nemo-ts/library/file"
	"github.com/janglucky/nemo-ts/ts/service"
	"path"
)

var inited = false

// InitRal ... initial ral
func InitRal(idc string) error {
	return initRalWithPath(idc, "/Users/jiangyi07/Github/src/nemo-ts/conf/ts/services", []string{"toml", "conf"})
}

func initRalWithPath(idc string, folder string, suffixs []string) error {
	if inited {
		return nil
	}
	inited = true

	return loadServicesFromFolder(folder, idc, suffixs)
}

func loadServicesFromFolder(folder, idc string, suffixs []string) error {
	files := file.ListFiles(folder)

	for _, fi := range files {
		fileWithSuffix := path.Base(fi)
		fileSuffix := path.Ext(fileWithSuffix)
		for _, suffix := range suffixs {
			if "."+suffix == fileSuffix {
				serc := &service.ServiceConfig{}
				err := file.DecodeTomlFile(fi, serc)
				if err != nil {
					continue
				}
				loadService(serc, idc)
			}
		}
	}
	return nil
}

func loadService(serc *service.ServiceConfig, idc string)  {
	service.ServiceMap.Store(serc.Name, serc)
}

// IsRalInited ... ral initial finished
func IsRalInited() bool {
	return inited
}
