package main

import (
	"fmt"

	"github.com/PrakharSrivastav/track-service-grpc/internal"
	"github.com/PrakharSrivastav/track-service-grpc/internal/rpc"
	"github.com/PrakharSrivastav/track-service-grpc/internal/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
)

func main() {
	collector, err := zipkin.NewHTTPCollector(zipkinHTTPEndpoint)
	checkErr(err)
	recorder := zipkin.NewRecorder(collector, debug, hostPort, serviceName)

	tracer, err := zipkin.NewTracer(
		recorder, zipkin.ClientServerSameSpan(sameSpan),
	)
	app := internal.NewApplication(tracer)

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

const (
	// Our service name.
	serviceName = "track-service"

	// Host + port of our service.
	hostPort = "127.0.0.1:6560"

	// Endpoint to send Zipkin spans to.
	zipkinHTTPEndpoint = "http://zipkin:9411/api/v1/spans"

	// Debug mode.
	debug = false

	// same span can be set to true for RPC style spans (Zipkin V1) vs Node style (OpenTracing)
	sameSpan = false
)
