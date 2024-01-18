package main

import (
	"demo/application"
	"demo/application/utils"
	"demo/config"
)

func main() {
	//configPath := cconfig.GetConfigPath(os.Getenv("ENVIRONMENT"))
	config, err := config.GetConfig("./config/config-dev")
	if err != nil {
		panic(err)
	}
	////init logger
	//validateConfig(config)
	//logger, _ := logger.Init(&config.Logger)
	//defer func() {
	//	_ = logger.Sync()
	//}()
	//undo := zap.ReplaceGlobals(logger)
	//defer undo()

	s := application.NewServer(config)

	go func() {
		s.Start()
	}()
	//defer s.Stop()

	utils.WaitShutdown()
}
