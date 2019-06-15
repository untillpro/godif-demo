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
	"github.com/untillpro/godif"
	"github.com/untillpro/godif/iservices"
	"github.com/untillpro/godif/services"
)

func Test_StartStop(t *testing.T) {
	ctx := start(t)
	defer stop(ctx, t)

	log.Println("### Service:", *getService(ctx))

	/*
		Your tests here
	*/
}

func start(t *testing.T) context.Context {

	// Provide iservices interface
	services.DeclareRequire()

	// Declare own service
	Declare()

	errs := godif.ResolveAll()
	require.True(t, len(errs) == 0, "Resolve problem", errs)

	ctx, err := iservices.Start(context.Background())
	require.Nil(t, err)
	return ctx
}

func stop(ctx context.Context, t *testing.T) {
	iservices.Stop(ctx)
	godif.Reset()
}
