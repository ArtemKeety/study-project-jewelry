package main

import (
	"context"
	jewelry "curs"
	_ "curs/docs" // импортируйте сгенерированные документы
	"curs/pkg/handler"
	"curs/pkg/repository"
	"curs/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "github.com/swaggo/http-swagger" // http-swagger для Swagger UI
	"os"
	"os/signal"
)

// @title TODO APP Jewelry
// @version 1.0
// @description API Server
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error in initial Config %s", err.Error())
	}

	db, err := repository.NewMysqldb(repository.ConfigDB{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		Dbname:   viper.GetString("db.dbname"),
	})

	if err != nil {
		logrus.Fatalf("error in initial DB %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	h := handler.NewHandler(services)

	srv := new(jewelry.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), h.InitRoutes()); err != nil {
			logrus.Fatalf("error in start server Config %s", err.Error())
		}
	}()

	logrus.Println("start server")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	if err := srv.ShutDown(context.Background()); err != nil {
		logrus.Fatalf("error in shutdown Server %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
