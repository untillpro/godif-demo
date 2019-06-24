/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package finder

import (
	"github.com/untillpro/godif"
	"github.com/untillpro/godif-demo/ikvdb"
)

// Declare s.e.
func Declare() {
	godif.Require(&ikvdb.Get)
}
