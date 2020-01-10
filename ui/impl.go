package ui

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/untillpro/godif-demo/answerer"
	"github.com/untillpro/godif-demo/iconfig"
)

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
			answerer.Correct(ctx, prevQuestion, answer)
		} else {
			answer := answerer.Answer(ctx, question)
			fmt.Println(answer)
			prevQuestion = question
		}
	}
}
