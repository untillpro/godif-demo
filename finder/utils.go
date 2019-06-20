/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package finder

import (
	"context"
	"regexp"

	"github.com/untillpro/godif-demo/ikvdb"
)

// Find records which match `rex` using ikbdb.Get
func Find(ctx context.Context, rex string) map[string]ikvdb.Record {
	res := map[string]ikvdb.Record{}
	re := regexp.MustCompile(rex)
	recs, err := ikvdb.Get(ctx, "")
	if err != nil {
		return res
	}

	for key, rec := range recs {
		matches := re.FindStringSubmatch(rec.Value)
		if len(matches) > 0 {
			res[key] = rec
		}
	}
	return res
}
