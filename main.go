/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"path"

	"github.com/untillpro/godif"
	"github.com/untillpro/godif-demo/iconfig"
	"github.com/untillpro/godif-demo/iconfigfile"
	"github.com/untillpro/godif/services"
)

func main() {

	// Parse command line

	pInMem := flag.Bool("m", false, "Use in-memory key-value database")
	flag.Parse()

	// Disable logging in services
	services.DisableLogging()

	// Declare others

	iconfigfile.Declare(path.Join(".", ".configs"))
	Declare(*pInMem, "ui")

	// Run
	err := services.Run()

	if err != nil {
		log.Println("Error :", err)
	}

}

// Declare s.e.
func Declare(inMem bool, configName string) {

	godif.Require(&iconfig.GetConfig)
	godif.Require(&iconfig.PutConfig)

	godif.ProvideSliceElement(&services.Services, &service{inMem: inMem, configName: configName})

}

type service struct {
	inMem      bool
	configName string
}

func ui(ctx context.Context, s *service) {

	config := map[string]string{"ResetMe": "1"}
	iconfig.GetConfig(ctx, s.configName, &config)
	if config["ResetMe"] == "1" {
		config = map[string]string{"Prompt": "Your question: ", "ResetMe": "0"}
		iconfig.PutConfig(ctx, s.configName, config)
	}

	fmt.Println("# Intro")
	fmt.Println("")
	fmt.Println("- I'm a simple chatbot and I can demonstrate `godif` library features")
	fmt.Println("- Ask me questions, I'll try to answer")
	fmt.Println("- If I'm wrong type 'wrong' and you will be prompted for the correct answer")

	if s.inMem {
		fmt.Println("")
		fmt.Println("# In-memory Mode")
		fmt.Println("")
		fmt.Println("- NOTE: Your answers won't be saved since in-memory database is used")
	}
	fmt.Println("")
	fmt.Println("--------------")
	fmt.Println("")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(config["Prompt"])
		if !scanner.Scan() {
			break
		}
		text := scanner.Text()
		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}
	}
}

func (s *service) Start(ctx context.Context) (context.Context, error) {
	go ui(ctx, s)
	return ctx, nil
}

func (s *service) Stop(ctx context.Context) {
	fmt.Println("")
	fmt.Println("Bye, see you later")
}
