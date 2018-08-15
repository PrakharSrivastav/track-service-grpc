package main

import (
	"fmt"

	"github.com/PrakharSrivastav/track-service-grpc/internal"
	"github.com/PrakharSrivastav/track-service-grpc/internal/rpc"
	"github.com/PrakharSrivastav/track-service-grpc/internal/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app := internal.NewApplication()

	// Create a new db connection
	connection, err := sqlx.Open("sqlite3", "./track.db")
	checkErr(err)

	// inject db connection to internal service
	internal := service.New(connection)

	// Initial data migration
	checkErr(internal.CleanupAndInit())

	// inject inter service to trackService
	trackService := rpc.New(internal)

	// Inject track service to grpc
	app.GrpcServer.Add(trackService)

	// Start grpc server
	checkErr(app.GrpcServer.Start())
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error is ", err.Error())
		panic(err)
	}
}
