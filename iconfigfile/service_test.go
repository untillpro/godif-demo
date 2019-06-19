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
	"io/ioutil"
	"log"
	"os"
	"testing"

	intf "github.com/untillpro/godif-demo/iconfig"

	"github.com/stretchr/testify/require"
	"github.com/untillpro/godif/services"
)

func Test_Service(t *testing.T) {
	ctx, err := setUp(t)
	defer tearDown(ctx, t)
	require.Nil(t, err, err)

	log.Println("### Service:", service)
}

var tempFolder string

func setUp(t *testing.T) (context.Context, error) {
	var err error
	tempFolder, err = ioutil.TempDir("", "iconfigmem")
	if nil != err {
		return nil, err
	}
	Declare(tempFolder)
	intf.DeclareForTest()
	return services.ResolveAndStart()
}

func tearDown(ctx context.Context, t *testing.T) {
	services.StopAndReset(ctx)
	log.Println("os.RemoveAll", tempFolder)
	err := os.RemoveAll(tempFolder)
	if nil != err {
		log.Println("os.RemoveAll error:", err)
	}
}
