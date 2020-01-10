/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package answerer

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
		mockGetErr = nil
		mockGetRecords = map[string]ikvdb.Record{}

		answer := Answer(ctx, "")
		assert.Equal(t, 0, len(answer))
	}

	// PopularQuestions
	{
		mockGetErr = nil
		mockGetRecords = map[string]ikvdb.Record{}

		// Emulate Get results
		for q, a := range PopularQuestions() {
			mockGetRecords[q] = ikvdb.NewRecord(a)
		}

		assert.Equal(t, "Not bad, thanks", Answer(ctx, "How are you"))
		assert.Equal(t, "Not bad, thanks", Answer(ctx, "How ae yu"))
		assert.Equal(t, "I like robots, computers, and answering", Answer(ctx, "What are your hobbies"))
		assert.Equal(t, "I like robots, computers, and answering", Answer(ctx, "And your hobbies?"))
		assert.Equal(t, "I like robots, computers, and answering", Answer(ctx, "Any hobby?"))

	}
}

func Test_PopularQuestions(t *testing.T) {
	ctx, err := setUp(t)
	defer tearDown(ctx, t)
	require.Nil(t, err, err)

	// All answers must match

	mockGetErr = nil
	mockGetRecords = map[string]ikvdb.Record{}

	// Emulate Get results
	for q, a := range PopularQuestions() {
		mockGetRecords[q] = ikvdb.NewRecord(a)
	}

	for q, a := range PopularQuestions() {
		assert.Equal(t, a, Answer(ctx, q))
	}
}
