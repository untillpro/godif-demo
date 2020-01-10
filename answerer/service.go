package answerer

import (
	"context"
	"sync"

	"github.com/untillpro/godif-demo/ikvdb"
)

type service struct {
	inMem      bool
	configName string
	wg         sync.WaitGroup
}

func (s *service) Start(ctx context.Context) (context.Context, error) {
	initDb(ctx)
	return ctx, nil
}

func (s *service) Stop(ctx context.Context) {

}

func initDb(ctx context.Context) {
	for key, value := range PopularQuestions() {
		ikvdb.Put(ctx, key, value)
	}
}
