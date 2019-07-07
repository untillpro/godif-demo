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
// s2 can be sginifically longer, than s1
// Not commutative, of course
func StringSimilarity(s1, s2 string) int {
	coords1 := CalcStringCoords(s1)
	coords2 := CalcStringCoords(s2)

	sum := 0
	unmatched := 0
	for i, c1 := range coords1 {
		c2 := coords2[i]
		if c1 > 0 && c2 < c1 {
			sum += (c1 - c2) * (c1 - c2)
		}
		if c1 == 0 && c2 > 0 {
			unmatched++
		}
	}

	// Small penalty for unmatched characters
	return sum + unmatched/2
}

// PopularQuestions returns popular questions and their answers
func PopularQuestions() map[string]string {
	res := map[string]string{}
	res["Are you real"] = "Yes I am a real computer program"
	res["What is your name"] = "My name is unknown so far"
	res["How old are you"] = "I was born in 2019"
	res["Where do you live"] = "I'm living in computer"
	res["Can you help me"] = "It depends"
	res["Which languages do you speak"] = "Only English, please"
	res["How are you"] = "Not bad, thanks"
	res["Hi"] = "Hello!"
	res["What time is it"] = "Who notes, in happiness, how time is flying?"
	res["What are your hobbies"] = "I like robots, computers, and answering"
	res["What do you look like"] = "Like an 18 year old student"
	res["Are you human"] = "Not quite sure"
	res["Why"] = "What `why`?"
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
