package main

import (
	"fmt"
	"os"

	"response-service/pkg/config"
	"response-service/pkg/http/rest"
	"response-service/pkg/rabbitmq"
	"response-service/pkg/storage/mysql"
	"response-service/pkg/util"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	cfg := config.NewConfig()

	err := cfg.LoadConfig()
	if err != nil {
		return util.WrapErrorf(err, util.ErrorCodeInternal, "Environment configuration failed")
	}

	sql, err := mysql.NewMySQLConnection(cfg.MySQL)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", "Oh wow the service is online")

	server := rest.NewServer(
		cfg.Version,
		cfg.Environment,
		cfg.HTTP,
		sql,
	)

	server.Init()
	go rabbitmq.StartServer(cfg.RabbitMQ, server.ResponseService)

	// Runs the new server instance.
	server.Run(cfg.Name)

	return nil
}
