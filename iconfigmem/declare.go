/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package iconfigmem

import (
	"github.com/untillpro/godif"
	"github.com/untillpro/godif-demo/iconfig"
	"github.com/untillpro/godif/services"
)

// Declare s.e.
func Declare() {

	// Functions

	godif.Provide(&iconfig.GetConfig, getConfig)
	godif.Provide(&iconfig.PutConfig, putConfig)

	// Service
	var service Service
	service.configs = make(map[string][]byte)
	godif.ProvideSliceElement(&services.Services, &service)

}
