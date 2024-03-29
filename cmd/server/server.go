package server

import (
	"context"
	"metrics/api"
	"metrics/config"
	"metrics/manager"
	"metrics/storage"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"golang.org/x/sync/errgroup"
)

// StartServerCommand - старт сервера для обработки запросов
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

	strg, storageErr := storage.NewFileStorage(cfg.File)
	if storageErr != nil {
		return storageErr
	}

	metricManager := manager.NewMetricManager(strg, 100, cfg.Period)

	router := gin.New()

	api.NewAPIController(metricManager, router)
	srv := http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	//создание группы горутин с errgroup
	grp, _ := errgroup.WithContext(context.Background())
	grp.Go(srv.ListenAndServe) // запуск HTTP сервера

	grp.Go(metricManager.Run) // запуск обработчика сообщений

	sig := make(chan os.Signal, 1) // добавляем обработку сигналов для graceful shutdown
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	srv.Shutdown(context.Background()) //nolint:errcheck
	metricManager.Close()

	if err := grp.Wait(); err != nil { // ожидаем ошибки от обработчиков
		return err
	}
	return nil
}
