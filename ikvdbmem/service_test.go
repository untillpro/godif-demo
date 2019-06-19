/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*

	Test service start/stop here

*/

package ikvdbmem

import (
	"context"
	"log"
	"testing"

	intf "github.com/untillpro/godif-demo/ikvdb"

	"github.com/stretchr/testify/require"
	"github.com/untillpro/godif/services"
)

func Test_StartStop(t *testing.T) {
	ctx, err := start(t)
	defer stop(ctx, t)
	require.Nil(t, err, err)

	log.Println("### Service:", *getService(ctx))
}

func start(t *testing.T) (context.Context, error) {
	Declare()
	intf.DeclareForTest()
	return services.ResolveAndStart()
}

func stop(ctx context.Context, t *testing.T) {
	services.StopAndReset(ctx)
}
