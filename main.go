package main

import (
	"fmt"

	"github.com/PrakharSrivastav/artist-service-grpc/internal"
	"github.com/PrakharSrivastav/artist-service-grpc/internal/rpc"
)

func main() {
	app := internal.NewApplication()

	artistService := rpc.New(nil)

	app.GrpcServer.Add(artistService)
	err := app.GrpcServer.Start()
	if err != nil {
		fmt.Println(err.Error())
	}
}
