/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package ikvdb

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Marshall(t *testing.T) {
	r := NewRecord("quick brown fox")
	data, err := json.MarshalIndent(r, "", "  ")
	assert.Nil(t, err, err)
	var r2 Record
	err = json.Unmarshal(data, &r2)
	assert.Nil(t, err, err)
	fmt.Println(r, r2)
}
