package main

import (
	jewelry "curs"
	"curs/pkg/handler"
	"curs/pkg/repository"
	"curs/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

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
	if err := srv.Run(viper.GetString("port"), h.InitRoutes()); err != nil {
		logrus.Fatalf("error in start server Config %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
