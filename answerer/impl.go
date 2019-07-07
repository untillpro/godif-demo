/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package answerer

import (
	"context"

	"github.com/untillpro/godif-demo/ikvdb"
)

// Answer the question
func Answer(ctx context.Context, question string) string {
	recs, err := ikvdb.Get(ctx, "")
	if err != nil {
		return ""
	}

	best := ""
	bestSimilarity := 10000

	for q, rec := range recs {
		similarity := StringSimilarity(question, q)
		if similarity < bestSimilarity {
			bestSimilarity = similarity
			best = rec.Value
		}
	}
	return best
}
