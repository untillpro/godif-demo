/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*

	Test service start/stop here

*/

package iconfigmem

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	intf "github.com/untillpro/godif-demo/iconfig"
	"github.com/untillpro/godif/services"
)

func Test_StartStop(t *testing.T) {
	ctx, err := setUp(t)
	defer tearDown(ctx, t)
	require.Nil(t, err, err)

	log.Println("### Service:", *getService(ctx))
}

func setUp(t *testing.T) (context.Context, error) {
	Declare()
	intf.DeclareForTest()
	return services.ResolveAndStart()
}

func tearDown(ctx context.Context, t *testing.T) {
	services.StopAndReset(ctx)
}
