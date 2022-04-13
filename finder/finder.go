package finder

import (
	"github.com/janglucky/nemo-ts/finder/finders"
	_ "github.com/janglucky/nemo-ts/finder/finders"
)

func GetFinder(finderType string)  finders.FindersInter{
	return finders.Finders[finderType]
}
