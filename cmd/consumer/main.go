package main

import (
	"github.com/joho/godotenv"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	wbtechl0 "wb-tech-l0"
	"wb-tech-l0/pkg/handler"
	"wb-tech-l0/pkg/repository"
	"wb-tech-l0/pkg/service"
)

func main() {
	logrus.Println("Consumer started to work...")

	logrus.Println("Initializing the config...")
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	logrus.Println("Initialization completed.")

	logrus.Println("Loading env variables...")
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	logrus.Println("Loading completed.")

	logrus.Println("Initializing the database")
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialized db: %s", err.Error())
	}
	logrus.Println("Initialization completed.")

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	sNat, err := stan.Connect(
		"test-cluster",
		"order-consumer", stan.NatsURL(viper.GetString("nats.uri")),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer sNat.Close()

	if _, err = sNat.Subscribe("orders", services.CreateOrder, stan.StartWithLastReceived()); err != nil {
		return
	}

	srv := new(wbtechl0.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
