package main

import (
	"context"
	"fmt"
	"os"

	"github.com/alexpts/edu-go/cmd/api/di"
)

func main() {
	ctx := context.Background()
	exitCode := run(ctx)

	os.Exit(exitCode)
}

func run(ctx context.Context) (exitCode int) {
	a, cleanUp, err := di.InjectA()
	defer cleanUp()

	if err != nil {
		return 1
	}

	fmt.Println(a.Name)

	return 0
}
