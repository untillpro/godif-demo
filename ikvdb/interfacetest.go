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
	"time"

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
	t.Run("testModified", testModified)
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

	var records map[string]Record
	records, err = Get(ctx, "")
	require.Nil(t, err)
	assert.Equal(t, 3, len(records))
	assert.Equal(t, "v1", records["k1"].Value)
	assert.Equal(t, "v2", records["k2"].Value)
	assert.Equal(t, "v3", records["k3"].Value)

	// Get first value

	records, err = Get(ctx, "k1")
	require.Nil(t, err)
	assert.Equal(t, 1, len(records))
	assert.Equal(t, "v1", records["k1"].Value)

	// Get for non-existing key

	records, err = Get(ctx, "k-1")
	require.Nil(t, err)
	assert.Equal(t, 0, len(records))

	// Remove second value

	err = Remove(ctx, "k2")

	// Get all values, k2 should be deleted

	records, err = Get(ctx, "")
	require.Nil(t, err)
	assert.Equal(t, 2, len(records))
	assert.Equal(t, "v1", records["k1"].Value)
	assert.Equal(t, "v3", records["k3"].Value)

	// Remove all values

	err = Remove(ctx, "")
	records, err = Get(ctx, "")
	require.Nil(t, err)
	assert.Equal(t, 0, len(records))

}

func testModified(t *testing.T) {
	var err error

	// Remove all records
	Remove(ctx, "")

	// Put a record

	err = Put(ctx, "k1", "v1")
	require.Nil(t, err)

	// Get record mod time

	records, err := Get(ctx, "k1")
	mod1 := records["k1"].Modified

	// Sleep

	time.Sleep(20 * time.Millisecond)

	// Put record again

	err = Put(ctx, "k1", "v2")
	records, err = Get(ctx, "k1")
	mod2 := records["k1"].Modified

	// mod2 must be after mod1

	require.True(t, mod2.After(mod1))

}
