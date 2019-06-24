/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package iconfigfile

import (
	"context"
	"os"
)

// Service s.e.
type Service struct {
	configFolder string
}

var service Service

// Start service
func (s *Service) Start(ctx context.Context) (context.Context, error) {
	err := os.MkdirAll(s.configFolder, 0700)
	return ctx, err
}

// Stop service
func (s *Service) Stop(ctx context.Context) {

}
