package main

import (
	"todolist-v1/config"
	"todolist-v1/pkg/database"
	"todolist-v1/pkg/server"

	"github.com/sirupsen/logrus"

	activityHandler "todolist-v1/modules/activity/handler"
	activityRepo "todolist-v1/modules/activity/repository"
	activityUsecase "todolist-v1/modules/activity/usecase"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.WithError(err).Fatal("Failed to load configuration file")
	}

	db := database.NewPostgresDatabase()
	if err := db.Connect(cfg.Database.URL); err != nil {
		log.WithError(err).Fatal("Failed to connect to the database")
	}
	defer db.Close()

	srv := server.NewFiberServer(cfg)

	repo := activityRepo.NewActivityRepository(db.GetDB())
	usecase := activityUsecase.NewActivityUsecase(repo)
	handler := activityHandler.NewActivityHttpHandler(srv.GetEngine(), usecase)

	handler.RegisterRoutes()

	log.WithField("port", cfg.Server.Port).Info("Server is running")
	if err := srv.Start(); err != nil {
		log.WithError(err).Fatal("Failed to start server")
	}
}
