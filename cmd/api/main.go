package main

import (
	"github.com/alexpts/edu-go/cmd/api/di"
	"github.com/spf13/viper"
)

func main() {
	logger := di.InjectApiLogger()

	viper.AutomaticEnv()

	app, err := di.InjectApp()
	if err != nil {
		logger.Fatal().Err(err).Msg("Can`t inject app")
	}

	server := di.InjectHttpServer(app.FastHttpHandler)
	err = server.ListenAndServe(":3000")

	if err != nil {
		logger.Err(err).Msg("Can`t start http server")
	}
}

//func logOnStart(logger *zerolog.Logger) {
//	logger.Info().Dict("process", zerolog.Dict().
//		Int("pid", os.Getpid()).
//		Int("count_cpu", runtime.NumCPU()).
//		Str("go_ver", runtime.Version()[2:]),
//	).Msg("start service...")
//}
