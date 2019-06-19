/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package ikvdbmem

import (
	"testing"

	"github.com/stretchr/testify/require"
	intf "github.com/untillpro/godif-demo/ikvdb"
)

func Test_Impl(t *testing.T) {
	ctx, err := start(t)
	defer stop(ctx, t)
	require.Nil(t, err, err)

	intf.TestImpl(ctx, t)
}
