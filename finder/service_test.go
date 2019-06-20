/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*

	Test service start/stop here

*/

package finder

import (
	"context"
	"testing"

	"github.com/untillpro/godif"
	"github.com/untillpro/godif-demo/ikvdb"
	"github.com/stretchr/testify/require"
	"github.com/untillpro/godif/services"
)

func Test_Service(t *testing.T) {
	ctx, err := setUp(t)
	defer tearDown(ctx, t)
	require.Nil(t, err, err)
}

var retGetRecords map[string]ikvdb.Record
var retGetErr error

func get(ctx context.Context, key string) (records map[string]ikvdb.Record, err error){
	return retGetRecords, retGetErr
}


func setUp(t *testing.T) (context.Context, error) {
	Declare()
	godif.Provide(&ikvdb.Get, get)
	
	return services.ResolveAndStart()
}

func tearDown(ctx context.Context, t *testing.T) {
	services.StopAndReset(ctx)
}
