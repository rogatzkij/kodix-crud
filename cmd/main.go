package main

import (
	"context"
	"fmt"
	"github.com/rogatzkij/kodix-crud/config"
	"github.com/rogatzkij/kodix-crud/internal/core"
	"github.com/rogatzkij/kodix-crud/internal/handel"
	"github.com/rogatzkij/kodix-crud/internal/mongo"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Считываем настройки из переменных окружения
	mainConfig, err := config.GetConfigFromEnv()
	if err != nil {
		log.Error().Err(err).Msg("не можем прочитать настройки")
		return
	}

	// Создаем коннектор для БД
	mainConnector := mongo.NewConnector(mainConfig)
	defer mainConnector.Close()

	// Создаем ядро
	mainCore := &core.Core{
		Auto:  mainConnector,
		Brand: mainConnector,
	}

	// Создаем роутер
	mainRouter := handel.NewRouter(mainCore)

	// Создаем и запускаем сервер
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", mainConfig.Port),
		Handler: mainRouter,
	}

	log.Info().Msg("Сервис запускается")

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error().Err(err).Msg("не можем запустить сервер")
			os.Exit(1)
		}
	}()

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	log.Info().Msg("Сервис закрывается...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error().Err(err)

	}
	log.Info().Msg("Сервис закрыт")
}
