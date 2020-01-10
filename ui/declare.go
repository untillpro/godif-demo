package ui

import (
	"github.com/untillpro/godif"
	"github.com/untillpro/godif-demo/iconfig"
	"github.com/untillpro/godif-demo/ikvdb"
	"github.com/untillpro/godif/services"
)

// Declare s.e.
func Declare(inMem bool, configName string) {

	godif.Require(&iconfig.GetConfig)
	godif.Require(&iconfig.PutConfig)
	godif.Require(&ikvdb.Put)
	godif.ProvideSliceElement(&services.Services, &service{inMem: inMem, configName: configName})

}
