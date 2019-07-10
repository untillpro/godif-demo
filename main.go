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
	"path"
	"sync"

	"github.com/untillpro/godif"
	"github.com/untillpro/godif-demo/answerer"
	"github.com/untillpro/godif-demo/iconfig"
	"github.com/untillpro/godif-demo/iconfigfile"
	"github.com/untillpro/godif-demo/iconfigmem"
	"github.com/untillpro/godif-demo/ikvdb"
	"github.com/untillpro/godif-demo/ikvdbbbolt"
	"github.com/untillpro/godif-demo/ikvdbmem"
	"github.com/untillpro/godif/services"
)

func main() {

	// Parse command line

	pInMem := flag.Bool("m", false, "Use in-memory key-value database")
	pVerbose := flag.Bool("v", false, "Use verbose output")
	flag.Parse()

	// Declare others
	{

		// Declare ikvdb implementation depending on `-m` option
		if *pInMem {
			ikvdbmem.Declare()
			iconfigmem.Declare()
		} else {
			iconfigfile.Declare(path.Join(".", ".data"))
			ikvdbbbolt.Declare(path.Join(".data", "answers.db"))
		}

		answerer.Declare()

		Declare(*pInMem, "ui")
	}

	// Run

	services.SetVerbose(*pVerbose)
	err := services.Run()

	if err != nil {
		log.Println("Error :", err)
	}

}

// Declare s.e.
func Declare(inMem bool, configName string) {

	godif.Require(&iconfig.GetConfig)
	godif.Require(&iconfig.PutConfig)

	godif.Require(&ikvdb.Put)

	godif.ProvideSliceElement(&services.Services, &service{inMem: inMem, configName: configName})

}

type service struct {
	inMem      bool
	configName string
	wg         sync.WaitGroup
}

func ui(ctx context.Context, s *service) {
	defer s.wg.Done()

	// Reset config, if needed

	config := map[string]string{"ResetMe": "1"}
	iconfig.GetConfig(ctx, s.configName, &config)
	if config["ResetMe"] == "1" {
		config = map[string]string{"Prompt": "Your question: ", "Correct": "Enter correct answer: ", "ResetMe": "0"}
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

	var prevQuestion string
	var wrong bool
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(config["Prompt"])
		if !scanner.Scan() {
			break
		}
		question := scanner.Text()
		wrong = len(prevQuestion) > 0 && question == "wrong"
		if wrong {
			fmt.Print(config["Correct"])
			if !scanner.Scan() {
				break
			}
			answer := scanner.Text()
			if len(answer) == 0 {
				continue
			}
			ikvdb.Put(ctx, prevQuestion, answer)
		} else {
			answer := answerer.Answer(ctx, question)
			fmt.Println(answer)
			prevQuestion = question
		}
	}
}

func (s *service) Start(ctx context.Context) (context.Context, error) {

	initDb(ctx)

	s.wg.Add(1)
	go ui(ctx, s)
	return ctx, nil
}

func (s *service) Stop(ctx context.Context) {
	os.Stdin.Close()
	s.wg.Wait()
	fmt.Println("")
	fmt.Println("Bye, see you later!")
}

func initDb(ctx context.Context) {
	for key, value := range answerer.PopularQuestions() {
		ikvdb.Put(ctx, key, value)
	}
}
