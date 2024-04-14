package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/mvp-mogila/avito-intership-backend-2024/internal/config"
	"github.com/mvp-mogila/avito-intership-backend-2024/internal/delivery"
	"github.com/mvp-mogila/avito-intership-backend-2024/internal/pkg/middleware"
	"github.com/mvp-mogila/avito-intership-backend-2024/internal/repository"
	"github.com/mvp-mogila/avito-intership-backend-2024/internal/usecase"
	"github.com/mvp-mogila/avito-intership-backend-2024/pkg/postgres"
	"github.com/mvp-mogila/avito-intership-backend-2024/pkg/redis"
)

func main() {
	cfg := config.GetConfig()

	router := mux.NewRouter()

	userUsecase := usecase.NewUserUsecase(cfg.Auth)

	router.Use(middleware.Authentication(userUsecase))

	pgx, err := postgres.NewPgxDatabase(cfg.Postgres)
	if err != nil {
		log.Fatal("postgres connection error")
	}
	defer pgx.Close()

	redis, err := redis.NewRedisCache(cfg.Redis)
	if err != nil {
		log.Println("redis connection error")
	}
	defer redis.Close()

	bannerRepo := repository.NewBannerRepo(pgx)
	bannerCache := repository.NewBannerCache(redis)
	bannerUsecase := usecase.NewBannerUsecase(bannerRepo, bannerCache)
	bannerHandler := delivery.NewBannerHandler(bannerUsecase)
	bannerHandler.SetupRouting(router)

	srv := &http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("Fatal error")
		}
		log.Printf("server is listening on %s...", srv.Addr)
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	_ = <-c

	waitTime, err := time.ParseDuration(cfg.CloseTime)
	if err != nil {
		waitTime = 5 * time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), waitTime)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Shutdown failed")
	}

	log.Println("Server stopped")
}
