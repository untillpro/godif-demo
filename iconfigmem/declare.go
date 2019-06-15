/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package iconfigmem

import (
	"github.com/untillpro/godif"
	"github.com/untillpro/godif/iservices"
)

// Service s.e.
type Service struct {
	storage map[string]interface{}
}

// Declare s.e.
func Declare() {
	var service Service
	godif.ProvideSliceElement(&iservices.Services, &service)

	//godif.Provide(&???.???, implFunc)
}
