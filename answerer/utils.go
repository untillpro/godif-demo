/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package answerer

import "strings"

// Dimensions in string coord
const Dimensions = 255

// StringSimilarity calculates if s1 can be made of s2
// s2 can be significally longer, than s1
// Not commutative, of course
func StringSimilarity(s1, s2 string) int {
	coords1 := CalcStringCoords(s1)
	coords2 := CalcStringCoords(s2)

	sum := 0
	unmatchedFromS2 := 0
	unmatchedFromS1 := 0
	for i, c1 := range coords1 {
		c2 := coords2[i]
		if c1 > 0 && c2 < c1 {
			sum += (c1 - c2) * (c1 - c2)
		}
		if c1 == 0 && c2 > 0 {
			unmatchedFromS2++
		}
		if c1 > 0 && c2 == 0 {
			unmatchedFromS1++
		}
	}

	// Add small penalty for unmatched characters
	return sum + unmatchedFromS2/2 + unmatchedFromS1
}

// PopularQuestions returns popular questions and their answers
func PopularQuestions() map[string]string {
	res := map[string]string{
		"human":                        "Not quite sure, seems not",
		"Are you real":                 "Yes I am a real computer program",
		"Can you help me":              "May be I can help, but it depends",
		"Hi":                           "Hello!",
		"How are you":                  "Not bad, thanks",
		"How old are you":              "I was born in 2019",
		"What is your name":            "My name is unknown so far",
		"Where do you live":            "I'm living in computer",
		"Which languages do you speak": "Only English, please",
		"What time is it":              "Who notes, in happiness, how time is flying?",
		"What are your hobbies":        "I like robots, computers, and answering",
		"What do you look like":        "Like an 18 year old student",
		"Why":                          "What `why`?",
		" ":                            "Please enter something",
	}
	return res
}

// CalcStringCoords s.e.
func CalcStringCoords(s string) (res []int) {
	s = strings.ToLower(s)
	for i := 0; i < Dimensions; i++ {
		res = append(res, 0)
	}

	for _, char := range s {
		dim := int(char) % Dimensions
		res[dim] = res[dim] + 1
	}

	return
}
