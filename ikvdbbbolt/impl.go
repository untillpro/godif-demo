/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package ikvdbbbolt

import (
	"context"
	"encoding/json"

	bolt "github.com/coreos/bbolt"
	intf "github.com/untillpro/godif-demo/ikvdb"
	//bolt "github.com/coreos/bbolt"
)

func get(ctx context.Context, key string) (values map[string]intf.Record, err error) {
	values = map[string]intf.Record{}
	if len(key) == 0 {
		err = getService(ctx).db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(bucketName))
			err := b.ForEach(func(k, v []byte) error {
				var r intf.Record
				json.Unmarshal(v, &r)
				values[string(k)] = r
				return nil
			})
			return err
		})
		return
	}
	err = getService(ctx).db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		data := b.Get([]byte(key))
		if nil != data {
			var r intf.Record
			json.Unmarshal(data, &r)
			values[key] = r
		}
		return nil
	})
	return
}

func put(ctx context.Context, key, value string) error {
	return getService(ctx).db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		r := intf.NewRecord(value)
		data, err := json.Marshal(r)
		if nil != err {
			return err
		}
		return b.Put([]byte(key), data)
	})
}

func remove(ctx context.Context, key string) error {
	var err error
	if len(key) == 0 {
		err = getService(ctx).db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(bucketName))
			err := b.ForEach(func(k, v []byte) error {
				err := b.Delete(k)
				if nil != err {
					return err
				}
				return nil
			})
			return err
		})
		return err
	}
	err = getService(ctx).db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		return b.Delete([]byte(key))
	})
	return err
}
