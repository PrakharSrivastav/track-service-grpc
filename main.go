package main

import (
	"fmt"

	"github.com/PrakharSrivastav/track-service-grpc/internal"
	"github.com/PrakharSrivastav/track-service-grpc/internal/rpc"
)

func main() {
	app := internal.NewApplication()

	trackService := rpc.New(nil)

	app.GrpcServer.Add(trackService)
	err := app.GrpcServer.Start()
	if err != nil {
		fmt.Println(err.Error())
	}
}
