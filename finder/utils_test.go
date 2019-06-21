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
)

func checkFirstDistanceLess(t *testing.T, q, a1, a2 string) {
	d1 := MatchAnswer(q, a1)
	d2 := MatchAnswer(q, a2)
	assert.True(t, d1 < d2, "%v, %v, %v: %d %d", q, a1, a2, d1, d2)
}

func Test_CalcDistanceBetween(t *testing.T) {

	checkFirstDistanceLess(t, "nme?", "What is your name", "Where do you live")
	checkFirstDistanceLess(t, "Now about TV", "How much does that TV cost", "What is your name")
	checkFirstDistanceLess(t, "Any ideas TV prce", "What you think about TV price", "What you think about car price")

}
