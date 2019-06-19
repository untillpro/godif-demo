/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package ikvdbbbolt

import (
	"context"

	intf "github.com/untillpro/godif-demo/ikvdb"
)

func get(ctx context.Context, key string) (values map[string]intf.Record, err error) {
	values = map[string]intf.Record{}
	return values, nil
}

func put(ctx context.Context, key, value string) error {
	return nil
}

func remove(ctx context.Context, key string) error {

	return nil
}
