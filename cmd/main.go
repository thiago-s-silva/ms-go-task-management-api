package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thiago-s-silva/ms-go-task-management-api/config"
	"github.com/thiago-s-silva/ms-go-task-management-api/internal/app"
	"github.com/thiago-s-silva/ms-go-task-management-api/pkg"
)

func main() {
	// init Logger
	logger := pkg.NewLogger("main.go")

	// init Gin Gonic
	server := gin.Default()

	// init Config
	cfg, err := config.Load()
	if err != nil {
		logger.Errorf("error when tried to load config package: %v", err)
		panic(err)
	}

	// init App
	app.Load(&app.Dependencies{
		Server:       server,
		ServerConfig: &cfg.ServerConfig,
	})

	// initialize Server
	server.Run(cfg.ServerConfig.SERVER_HOST)
}
