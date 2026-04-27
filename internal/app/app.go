package app

import (
	"github.com/fruty-gw/banana/internal/app/grpcapp"
	"github.com/fruty-gw/banana/internal/service"
	grpchandler "github.com/fruty-gw/banana/internal/transport/grpc"
)

type App struct {
	GRPCapp *grpcapp.App
}

func New() *App {
	a := &App{}

	bananaService := service.NewBananaService()
	bananaHandlers := grpchandler.NewBananaHandler(bananaService)

	a.GRPCapp = grpcapp.New(
		bananaHandlers,
		8981,
	)

	return a
}
