/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package ikvdb

import "context"

// Get returns all values if key is empty and value for given key otherwise
var Get func(ctx context.Context, key string) (records map[string]Record, err error)

// Put s.e.
var Put func(ctx context.Context, key, value string) error

// Remove all values if key is empty, or value with given key otherwise
var Remove func(ctx context.Context, key string) error
