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
)

func Test_StringSimilarity(t *testing.T) {

	checkFirstDistanceLess(t, "Ay hobby", "What are your hobbies", "How are you")
	checkFirstDistanceLess(t, "How ae yu", "How are you", "How old are you")
	checkFirstDistanceLess(t, "nme?", "What is your name", "Where do you live")
	checkFirstDistanceLess(t, "how much TV", "How much does that TV cost", "What is your name")
	checkFirstDistanceLess(t, "Any ideas TV prce", "What you think about TV price", "What you think about car price")
}

func checkFirstDistanceLess(t *testing.T, q, a1, a2 string) {
	d1 := StringSimilarity(q, a1)
	d2 := StringSimilarity(q, a2)
	assert.True(t, d1 < d2, "%v, %v, %v: %d %d", q, a1, a2, d1, d2)
}
