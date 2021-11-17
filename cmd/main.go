package main

import (
	"chat"
	"chat/pkg/handler"
	"chat/pkg/repository"
	"chat/pkg/service"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	db, err := repository.NewPostgresDB(repository.PostgresConfig{
		"",
		"",
		"",
		"",
		"",
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	port := os.Getenv("PORT")
	if port == "" {
		port = "1234"
	}

	srv := new(chat.Server)
	go func() {
		if err = srv.Run(port, handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Print("App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("App Shutting Down")

	if err = srv.Shutdown(context.Background()); err != nil {
		log.Printf("error occured on server shutting down: %s", err.Error())
	}

	if err = db.Close(); err != nil {
		log.Printf("error occured on db connection close: %s", err.Error())
	}
}
