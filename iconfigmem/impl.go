/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package iconfigmem

import (
	"context"
	"encoding/json"
	"errors"
)

func getConfig(ctx context.Context, configName string, config interface{}) error {

	if len(configName) == 0 {
		return errors.New("Empty configName")
	}

	service := getService(ctx)
	res, ok := service.configs[configName]
	if !ok {
		return nil
	}
	err := json.Unmarshal(res, &config)
	if err != nil {
		return err
	}
	return nil
}

func putConfig(ctx context.Context, configName string, config interface{}) error {

	if len(configName) == 0 {
		return errors.New("Empty configName")
	}

	service := getService(ctx)

	data, err := json.Marshal(config)
	if err != nil {
		return err
	}

	service.configs[configName] = data
	return nil
}
