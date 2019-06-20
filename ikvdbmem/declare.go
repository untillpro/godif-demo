/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package ikvdbmem

import (
	"github.com/untillpro/godif"
	intf "github.com/untillpro/godif-demo/ikvdb"
	"github.com/untillpro/godif/services"
)

// Declare s.e.
func Declare() {

	// Functions

	godif.Provide(&intf.Get, get)
	godif.Provide(&intf.Put, put)
	godif.Provide(&intf.Remove, remove)

	// Service
	var service Service
	service.data = make(map[string]intf.Record)
	godif.ProvideSliceElement(&services.Services, &service)

}
