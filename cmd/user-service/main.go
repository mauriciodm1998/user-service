package main

import (
	"os"
	"os/signal"
	"user-service/internal/channels/grpc"
	"user-service/internal/channels/rest"
	"user-service/internal/config"
	repository "user-service/internal/repositories"
	"user-service/internal/service"

	"github.com/rs/zerolog/log"

	"github.com/sirupsen/logrus"
)

var (
	cfg = &config.Cfg
)

func main() {
	config.ParseFromFlags()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	log.Info().Any("config", config.Get()).Msg("configuration file")

	customerService, userService := startDependencies()

	go func() {
		if err := rest.New(customerService, userService).Start(); err != nil {
			logrus.Panic()
		}
	}()

	go func() {
		if err := grpc.ListenUser(cfg.UserServer.Port, userService); err != nil {
			logrus.Panic()
		}
	}()

	go func() {
		if err := grpc.ListenCustomer(cfg.CustomerServer.Port, customerService); err != nil {
			logrus.Panic()
		}
	}()

	<-stop
}

func startDependencies() (service.CustomerService, service.UserService) {
	customerRepository := repository.NewCustomerRepo(repository.New())
	customerService := service.NewCustomerService(customerRepository)

	userRepository := repository.NewUserRepo(repository.New())
	userService := service.NewUserService(userRepository, customerService)

	return customerService, userService
}
