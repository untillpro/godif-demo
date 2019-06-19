/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package ikvdb

import "time"

// Record for Get/Put functions
type Record struct {
	Value    string
	Modified time.Time
}

// NewRecord constructs new record from value and current time
func NewRecord(value string) Record {
	return Record{value, time.Now()}
}
