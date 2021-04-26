package main

import (
	"github.com/Askaell/homework/pkg/service"
	"github.com/spf13/viper"

	_ "github.com/lib/pq" //postgres driver
)

func main() {
	// if err := initConfig(); err != nil {
	// 	log.Fatalf("error initializating configs: %s", err.Error())
	// }

	// db, err := repository.NewPostgresDB(repository.Config{
	// 	Host:     viper.GetString("db.host"),
	// 	Port:     viper.GetString("db.port"),
	// 	Username: viper.GetString("db.username"),
	// 	Password: viper.GetString("db.password"),
	// 	DBname:   viper.GetString("db.dbname"),
	// 	SSLMode:  viper.GetString("db.sslmode"),
	// })
	// if err != nil {
	// 	log.Fatalf("failed to initialize db: %s", err.Error())
	// }

	// repository := repository.NewItemRepository(db)
	// handler := handler.NewHandler(repository)

	// server := new(server.Server)
	// go func() {
	// 	if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
	// 		log.Fatalf("error occured while running http server, %s", err.Error())
	// 	}
	// }()

	// quit := make(chan os.Signal)
	// signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	// <-quit

	discountService := service.NewDiscountService()
	discountService.Start("https://raw.githubusercontent.com/goarchitecture/lesson-2/feature/lesson-4/assets/discounts.csv", 5)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
