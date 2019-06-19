/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package iconfigfile

import (
	"github.com/untillpro/godif"
	intf "github.com/untillpro/godif-demo/iconfig"
	"github.com/untillpro/godif/services"
)

// Declare s.e.
func Declare(configFolder string) {

	// Functions

	godif.Provide(&intf.GetConfig, getConfig)
	godif.Provide(&intf.PutConfig, putConfig)

	service.configFolder = configFolder
	godif.ProvideSliceElement(&services.Services, &service)

}
