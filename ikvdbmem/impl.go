/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package ikvdbmem

import (
	"context"
	"time"

	intf "github.com/untillpro/godif-demo/ikvdb"
)

func get(ctx context.Context, key string) (values map[string]intf.Record, err error) {
	data := getService(ctx).data
	values = map[string]intf.Record{}
	if len(key) > 0 {
		value, ok := data[key]
		if ok {
			values[key] = value
		}
		return
	}
	for key, value := range data {
		values[key] = value
	}
	return
}

func put(ctx context.Context, key, value string) error {
	data := getService(ctx).data
	rec := data[key]
	rec.Value = value
	rec.Modified = time.Now()
	data[key] = rec
	return nil
}

func remove(ctx context.Context, key string) error {
	data := getService(ctx).data
	if len(key) > 0 {
		delete(data, key)
		return nil
	}
	getService(ctx).data = make(map[string]intf.Record)
	return nil
}
