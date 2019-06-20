/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package finder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/untillpro/godif-demo/ikvdb"
)

func Test_BasicUsage(t *testing.T) {
	ctx, err := setUp(t)
	defer tearDown(ctx, t)
	require.Nil(t, err, err)

	// Empty db
	{
		retGetErr = nil
		retGetRecords = map[string]ikvdb.Record{}

		recs := Find(ctx, ".*")
		assert.Equal(t, 0, len(recs))
	}

	// Three records
	{
		retGetErr = nil
		retGetRecords = map[string]ikvdb.Record{}
		retGetRecords["k1"] = ikvdb.NewRecord("2018-01-01")
		retGetRecords["k2"] = ikvdb.NewRecord("2018-02-02")
		retGetRecords["k3"] = ikvdb.NewRecord("2019-02-01")

		recs := Find(ctx, "01-01*")
		assert.Equal(t, 1, len(recs))
		assert.Equal(t, "2018-01-01", recs["k1"].Value)

		recs = Find(ctx, "-02")
		assert.Equal(t, 2, len(recs))
		assert.Equal(t, "2018-02-02", recs["k2"].Value)
		assert.Equal(t, "2019-02-01", recs["k3"].Value)

		recs = Find(ctx, "02$")
		assert.Equal(t, 1, len(recs))
		assert.Equal(t, "2018-02-02", recs["k2"].Value)

	}

}
