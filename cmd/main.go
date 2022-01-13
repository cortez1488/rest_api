package main

import (
	todoServer "github.com/cortez1488/rest_todo"
	handler "github.com/cortez1488/rest_todo/pkg/handler"
	"github.com/cortez1488/rest_todo/pkg/repository"
	service "github.com/cortez1488/rest_todo/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error init configs %s", err.Error())
	}

	db, err := repository.NewSqliteDb()
	if err != nil {
		logrus.Fatalf("error with get new database %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todoServer.Server)
	err = srv.Run(viper.GetString("port"), handlers.InitRoutes())
	if err != nil {
		logrus.Fatal(err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
