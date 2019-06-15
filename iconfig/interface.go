/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package iconfig

import "context"

// GetConfig fills `config` parameter from persistent config with given name
// If persistent config does not exist `config` parameter is left intact
var GetConfig func(ctx context.Context, configName string, config interface{}) error

// PutConfig saves config with given name
var PutConfig func(ctx context.Context, configName string, config interface{}) error
