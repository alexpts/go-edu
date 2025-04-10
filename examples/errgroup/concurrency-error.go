package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, gCtx := errgroup.WithContext(context.Background())
	g.SetLimit(3)

	m := sync.Mutex{}
	var joinErrors error

	for i := 0; i < 20; i++ {
		localI := i
		g.Go(func() error {
			if gCtx.Err() != nil {
				fmt.Printf("%d пропускаем горутину\n", localI)
				return nil
			}

			fmt.Printf("%d выполняем горутину\n", localI)
			r := rand.Intn(3000)
			time.Sleep(time.Duration(r) * time.Millisecond)

			m.Lock()
			defer m.Unlock()

			if i > 3 {
				joinErrors = errors.Join(joinErrors, errors.New("some error "+strconv.Itoa(i)))
			}

			return joinErrors
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println("111=====")
		fmt.Printf("\n\n%v\n\n", err)
	}

	fmt.Println("222=====")
	fmt.Printf("%v\n", joinErrors)
}
