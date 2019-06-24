/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	pInMem := flag.Bool("m", false, "Use in-memory key-value database")
	flag.Parse()

	fmt.Println("# Intro")
	fmt.Println("")
	fmt.Println("- I'm a simple chatbot to demonstrate godif library features")
	fmt.Println("- Ask me questions, I'll try to answer")
	fmt.Println("- If I'm wrong type 'wrong' and you will be prompted for the correct answer")

	if *pInMem {
		fmt.Println("")
		fmt.Println("# In-memory Mode")
		fmt.Println("")
		fmt.Println("- NOTE: Your answers won't be saved since in-memory database is used")
	}

	fmt.Println("")
	fmt.Println("---------------------")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Your question: ")
		if !scanner.Scan() {
			break
		}
		text := scanner.Text()
		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}
	}

}
