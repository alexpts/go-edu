package main

import (
	"context"
	"github.com/rs/zerolog"
	"os"
	"runtime"

	"github.com/alexpts/edu-go/cmd/api/di"
	"github.com/valyala/fasthttp"
)

func IndexHandler(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Welcome!")
}

func main() {
	ctx := context.Background()
	exitCode := run(ctx)

	os.Exit(exitCode)
}

func run(ctx context.Context) (exitCode int) {
	logger := di.InjectApiLogger()
	logOnStart(&logger)

	server := di.InjectHttpServer(IndexHandler)
	_ = server.ListenAndServe(":3000")

	return 0
}

func logOnStart(logger *zerolog.Logger) {
	logger.Info().Dict("process", zerolog.Dict().
		Int("pid", os.Getpid()).
		Int("count_cpu", runtime.NumCPU()).
		Str("go_ver", runtime.Version()[2:]),
	).Msg("start service...")
}
