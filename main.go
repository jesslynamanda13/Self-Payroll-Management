package main

import (
	"self-payrol/config"
	"self-payrol/config/postgres"
	"sync"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Infof(".env is not loaded properly")
	}
	log.Infof("read .env from file")

	db := postgres.InitGorm()
	if db == nil {
		log.Errorf("Failed to initialize the database")
		return
	}
	// defer db.Close()

	config := config.NewConfig()
	server := InitServer(config)

	log.Infof("testing")

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		server.Run()
	}()

	wg.Wait()
}
