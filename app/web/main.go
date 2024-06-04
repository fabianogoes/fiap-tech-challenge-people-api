package main

import (
	"context"
	"fmt"
	"github.com/fabianogoes/fiap-people/domain/usecases"
	"github.com/fabianogoes/fiap-people/frameworks/repository"
	"log/slog"
	"os"

	"github.com/fabianogoes/fiap-people/domain/entities"

	"github.com/fabianogoes/fiap-people/frameworks/rest"
)

func init() {
	fmt.Println("Initializing...")

	var logHandler *slog.JSONHandler

	config, _ := entities.NewConfig()
	if config.Environment == "production" {
		logHandler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		logHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}

	logger := slog.New(logHandler)
	slog.SetDefault(logger)
}

func main() {
	fmt.Println("Starting web server...")

	ctx := context.Background()
	var err error

	config, err := entities.NewConfig()
	if err != nil {
		panic(err)
	}
	db, err := repository.InitDB(ctx, config)
	if err != nil {
		panic(err)
	}

	attendantRepository := repository.NewAttendantRepository(db)
	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	attendantHandler := rest.NewAttendantHandler(attendantUseCase)

	customerRepository := repository.NewCustomerRepository(db)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	customerHandler := rest.NewCustomerHandler(customerUseCase, config)

	router, err := rest.NewRouter(
		customerHandler,
		attendantHandler,
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("DB connected")
	fmt.Println(db)

	err = router.Run(config.AppPort)
	if err != nil {
		panic(err)
	}
}
