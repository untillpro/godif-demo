/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package finder

// Dimensions in string coord
const Dimensions = 255

// CalcStringCoords s.e.
func CalcStringCoords(s string) (res []int) {
	for i := 0; i < Dimensions; i++ {
		res = append(res, 0)
	}

	for _, char := range s {
		dim := int(char) % Dimensions
		if res[dim] == 0 {
			res[dim] = 1
		}
	}

	return
}

// MatchAnswer to question
func MatchAnswer(question, answer string) int {
	coords1 := CalcStringCoords(question)
	coords2 := CalcStringCoords(answer)

	sum := 0
	for i, c1 := range coords1 {
		c2 := coords2[i]
		if c1 > 0 {
			sum += c1 - c2
		}
	}
	return sum
}
