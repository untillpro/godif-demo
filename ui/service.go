package ui

import (
	"context"
	"fmt"
	"os"
	"sync"
)

type service struct {
	inMem      bool
	configName string
	wg         sync.WaitGroup
}

func (s *service) Start(ctx context.Context) (context.Context, error) {

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
