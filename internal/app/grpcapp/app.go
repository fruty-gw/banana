package grpcapp

import (
	"fmt"
	"log"
	"net"

	bananav1 "github.com/fruty-gw/banana/gen/go/api/banana/v1"

	grpchandler "github.com/fruty-gw/banana/internal/transport/grpc"
	"google.golang.org/grpc"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
}

func New(
	handler *grpchandler.BananaHandler,
	port int,
) *App {

	gRPCServer := grpc.NewServer()
	bananav1.RegisterBananaServiceServer(gRPCServer, handler)

	return &App{
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	log.Println("Running on port", a.port)
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return err
	}

	return a.gRPCServer.Serve(l)
}

func (a *App) Stop() {
	a.gRPCServer.GracefulStop()
}
