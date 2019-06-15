/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package iconfigmem

import (
	"context"
)

// Service s.e.
type Service struct {
	configs map[string][]byte
}

type contextKeyType string

const contextKey = contextKeyType("contextKey")

func getService(ctx context.Context) *Service {
	return ctx.Value(contextKey).(*Service)
}

// Start service
func (s *Service) Start(ctx context.Context) (context.Context, error) {
	return context.WithValue(ctx, contextKey, s), nil
}

// Stop service
func (s *Service) Stop(ctx context.Context) {

}
