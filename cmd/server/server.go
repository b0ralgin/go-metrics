package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"golang.org/x/sync/errgroup"
	"metrics/api"
	"metrics/config"
	"metrics/manager"
	"metrics/storage"
	"net/http"
)

// StartHTTPServerCommand - старт HTTP сервера для обработки запросов по REST API
func StartServerCommand() cli.Command {
	return cli.Command{
		Name:   "server",
		Usage:  "Start Metric server",
		Action: startServer,
	}
}

func startServer(c *cli.Context) error {
	cfg := config.Config{}
	if configErr := config.LoadConfig(c.App.Name, &cfg); configErr != nil {
		return configErr
	}

	storage, storageErr := storage.NewFileStorage(cfg.File)
	if storageErr != nil {
		return storageErr
	}

	metricManager := manager.NewMetricManager(storage, 100, cfg.Period)

	router := gin.New()

	api.NewApiController(metricManager, router)
	srv := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	grp, _ := errgroup.WithContext(context.Background())
	grp.Go(srv.ListenAndServe)
	grp.Go(metricManager.Run)

	if err := grp.Wait(); err != nil {
		return err
	}
	return nil
}
