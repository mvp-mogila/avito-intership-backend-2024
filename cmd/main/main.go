package main

import (
	"log"
	"net/http"

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
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("start listening om port 8080 ...")
	if err = srv.ListenAndServe(); err != nil {
		log.Fatal("Fatal error")
	}

}
