package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"golang/api"
	"golang/internal/service"

	// "golang/internal/model/domain"
	"golang/utils"
	"log"
	"os"
	"time"
)

var ErrConnectToDatabase = errors.New("cannot Connect to Database")

func main() {
	if err := checkConnectionToDatabase(); err != nil {
		log.Fatal(err)
	}
}

func checkConnectionToDatabase() error {

	config, err := utils.LoadConfig(".")
	if err != nil {
		fmt.Println("Cannot not load config ", err)
	}
	// println(os.Getenv("DB_SOURCE"))
	conn, err := sql.Open(config.DBDriver, os.Getenv("DB_SOURCE_QUANG"))
	if err != nil {
		fmt.Println("Error when config to db", err)
	}
	conn.SetMaxOpenConns(10000)
	newContext, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := conn.PingContext(newContext); err != nil {
		return ErrConnectToDatabase
	}
	status := "Connect to Database Success"

	log.Println(status)
	store := service.NewServiceStore(conn)
	server := api.NewServer(store)

	if err != nil {
		fmt.Println("cannot create server:", err)
	}

	err = server.Start(config.ServerAddess)
	if err != nil {
		fmt.Println("cannot start server:", err)
	}
	return nil
	////
	////
}

///
// func slowStartDatabase(ctx context.Context) context.Context  {
// 	newContext , cancel := context.WithTimeout(context.Background(), time.Second*10)
// 	// time.Sleep(time.Second*1)
// 	defer cancel()
// 	select {
// 	case <-time.After(1 * time.Second):
// 		fmt.Println("overslept")
// 	case <-ctx.Done():
// 		fmt.Println(ErrConnectToDatabase) // prints "context deadline exceeded"
// 	}

// 	return newContext
// }
