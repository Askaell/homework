package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Askaell/homework/pkg/handler"
	"github.com/Askaell/homework/pkg/repository"
	"github.com/Askaell/homework/pkg/server"
	"github.com/Askaell/homework/pkg/service"
	"github.com/spf13/viper"

	_ "github.com/lib/pq" //postgres driver
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializating configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repository := repository.NewItemRepository(db)

	discountService := service.NewDiscountService(repository)
	discountService.Start(
		viper.GetString("discount_service.url"),
		viper.GetString("discount_service.activationTime"),
		viper.GetString("discount_service.location"))

	handler := handler.NewHandler(repository)

	server := new(server.Server)
	go func() {
		if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server, %s", err.Error())
		}
	}()

	// app shutting down
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := server.Shutdown(context.Background()); err != nil && err != http.ErrServerClosed {
		log.Printf("error occured on server shutting down: %s", err)
	}

	if err := db.Close(); err != nil {
		log.Printf("error occured on db connection close: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
