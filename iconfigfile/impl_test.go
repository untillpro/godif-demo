/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package iconfigfile

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/untillpro/godif-demo/iconfig"
)

func Test_Impl(t *testing.T) {
	ctx, err := setUp(t)
	defer tearDown(ctx, t)
	require.Nil(t, err, err)

	iconfig.TestImpl(ctx, t)
}
