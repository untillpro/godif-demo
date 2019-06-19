/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package ikvdbbbolt

import (
	"github.com/untillpro/godif"
	"github.com/untillpro/godif-demo/ikvdb"
	"github.com/untillpro/godif/services"
)

// Declare s.e.
func Declare(dbPath string) {

	// Functions

	godif.Provide(&ikvdb.Get, get)
	godif.Provide(&ikvdb.Put, put)
	godif.Provide(&ikvdb.Remove, remove)

	// Service
	var service = Service{dbPath: dbPath}
	godif.ProvideSliceElement(&services.Services, &service)

}
