package main

import (
	"context"
	todoServer "github.com/cortez1488/rest_todo"
	handler "github.com/cortez1488/rest_todo/pkg/handler"
	"github.com/cortez1488/rest_todo/pkg/repository"
	service "github.com/cortez1488/rest_todo/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error init configs %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading .env configs %s", err.Error())
	}

	db, err := repository.NewPostgresDb(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("error with get new database %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todoServer.Server)

	go func() {
		err = srv.Run(viper.GetString("port"), handlers.InitRoutes())
		if err != nil {
			logrus.Fatal(err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Print("TodoApp started")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occure on server shutting down %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error occure on db connection close %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
