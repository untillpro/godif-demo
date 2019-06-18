/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package ikvdb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/untillpro/godif"
)

// DeclareForTest s.e.
func DeclareForTest() {
	godif.Require(&Get)
	godif.Require(&Put)
}

var ctx context.Context

// TestImpl tests iconfig implementation
func TestImpl(actx context.Context, t *testing.T) {
	ctx = actx
	t.Run("testBasicUsage", testBasicUsage)
}

func testBasicUsage(t *testing.T) {

	var err error

	// Remove all records
	Remove(ctx, "")

	// Put three values

	err = Put(ctx, "k1", "v1")
	require.Nil(t, err)
	err = Put(ctx, "k2", "v2")
	require.Nil(t, err)
	err = Put(ctx, "k3", "v3")
	require.Nil(t, err)	


	// Get all values

	var values map[string]string
	values, err = Get(ctx, "")
	require.Nil(t, err)
	assert.Equal(t, 3, len(values))
	assert.Equal(t, "v1", values["k1"])
	assert.Equal(t, "v2", values["k2"])
	assert.Equal(t, "v3", values["k3"])

	// Get first value

	values, err = Get(ctx, "k1")
	require.Nil(t, err)
	assert.Equal(t, 1, len(values))
	assert.Equal(t, "v1", values["k1"])

	// Get for non-existing key

	values, err = Get(ctx, "k-1")
	require.Nil(t, err)
	assert.Equal(t, 0, len(values))

	// Remove second value

	err =  Remove(ctx, "k2")

	// Get all values, k2 should be deleted

	values, err = Get(ctx, "")
	require.Nil(t, err)
	assert.Equal(t, 2, len(values))
	assert.Equal(t, "v1", values["k1"])
	assert.Equal(t, "v3", values["k3"])

	// Remove all values

	err =  Remove(ctx, "")
	values, err = Get(ctx, "")
	require.Nil(t, err)
	assert.Equal(t, 0, len(values))

}
