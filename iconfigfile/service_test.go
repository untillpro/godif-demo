/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*

	Test service start/stop here

*/

package iconfigfile

import (
	"context"
	"os"
	"log"
	"testing"
	"io/ioutil"

	"github.com/untillpro/godif-demo/iconfig"

	"github.com/stretchr/testify/require"
	"github.com/untillpro/godif/services"
)

func Test_StartStop(t *testing.T) {
	ctx, err := start(t)
	defer stop(ctx, t)
	require.Nil(t, err, err)

	log.Println("### Service:", service)
}

var tempFolder string

func start(t *testing.T) (context.Context, error) {
	var err error
	tempFolder, err = ioutil.TempDir("", "iconfigmem")
	if nil != err{
		return nil, err
	}
	Declare(tempFolder)
	iconfig.DeclareForTest()
	return services.ResolveAndStart()
}

func stop(ctx context.Context, t *testing.T) {
	services.StopAndReset(ctx)
	log.Println("os.RemoveAll", tempFolder)
	err := os.RemoveAll(tempFolder)
	if nil != err {
		log.Println("os.RemoveAll error:", err)
	}
}
