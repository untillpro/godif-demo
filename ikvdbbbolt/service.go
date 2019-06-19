/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package ikvdbbbolt

import (
	"context"

	bolt "github.com/coreos/bbolt"
)

// Service s.e.
type Service struct {
	dbPath string
	db     *bolt.DB
}

type contextKeyType string

const (
	contextKey = contextKeyType("contextKey")
	bucketName = "ikvdbbbolt"
)

func getService(ctx context.Context) *Service {
	return ctx.Value(contextKey).(*Service)
}

// Start service
func (s *Service) Start(ctx context.Context) (context.Context, error) {

	db, err := bolt.Open(s.dbPath, 0600, bolt.DefaultOptions)
	if err != nil {
		return ctx, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return ctx, err
	}

	s.db = db

	return context.WithValue(ctx, contextKey, s), nil
}

// Stop service
func (s *Service) Stop(ctx context.Context) {
	s.db.Close()
	s.db = nil
}
