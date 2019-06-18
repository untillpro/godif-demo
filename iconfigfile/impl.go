/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package iconfigfile

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
)

func getConfig(ctx context.Context, configName string, config interface{}) error {

	if len(configName) == 0 {
		return errors.New("Empty configName")
	}

	configPath := path.Join(service.configFolder, configName+".json")
	data, err := ioutil.ReadFile(configPath)
	if os.IsNotExist(err) {
		return nil
	}
	if nil != err {
		return err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return err
	}
	return nil
}

func putConfig(ctx context.Context, configName string, config interface{}) error {

	if len(configName) == 0 {
		return errors.New("Empty configName")
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	configPath := path.Join(service.configFolder, configName+".json")
	err = ioutil.WriteFile(configPath, data, 0700)
	return err
}
