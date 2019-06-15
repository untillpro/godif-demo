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
	"github.com/untillpro/godif/iservices"
	"github.com/untillpro/godif/services"
	"github.com/untillpro/godif-demo/iconfig"
)

func Test_StartStop(t *testing.T) {
	ctx, err := start(t)
	defer stop(ctx, t)
	require.Nil(t, err, err)

	log.Println("### Service:", *getService(ctx))
}

func start(t *testing.T) (context.Context, error) {
	return iservices.StartInTest(t, services.Declare, iconfig.DeclareTest, Declare)
}

func stop(ctx context.Context, t *testing.T) {
	iservices.StopInTest(ctx, t)
}
