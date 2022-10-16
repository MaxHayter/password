package main

import (
	"flag"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"github.com/password/configs"
	"github.com/password/internal/controller"
	"github.com/password/internal/service"
	"github.com/password/logger"
	api "github.com/password/password"
)

const (
	configPath = "config.yaml"
)

func main() {
	configFile := flag.String("c", configPath, "specify path to a config.yaml")
	flag.Parse()

	log := logger.DefaultLogger

	config, err := configs.Configure(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	listen, err := net.Listen("tcp", config.Server.GrpcAddress)
	if err != nil {
		log.Fatal("failed to listen grpc port")
	}

	passwordController := controller.NewController(service.NewService())

	var serverOptions []grpc.ServerOption

	grpcServer := grpc.NewServer(serverOptions...)

	api.RegisterPasswordServiceServer(grpcServer, passwordController)

	signalListener := make(chan os.Signal, 1)
	defer close(signalListener)

	signal.Notify(signalListener,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		if err = grpcServer.Serve(listen); err != nil {
			log.Println("failed to listen grpc port")
		}
	}()

	defer func() {
		grpcServer.GracefulStop()
	}()

	stop := <-signalListener
	log.Println("Received", stop)
	log.Println("Waiting for all jobs to stop")
}
