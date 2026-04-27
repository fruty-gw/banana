package grpchandler

import (
	"context"

	bananav1 "github.com/fruty-gw/banana/gen/go/api/banana/v1"
)

type bananaService interface {
	GetBanana()
	UpdateBanana()
	DeleteBanana()
	AddBanana()
}

type BananaHandler struct {
	bananav1.UnimplementedBananaServiceServer
	svc bananaService
}

func NewBananaHandler(svc bananaService) *BananaHandler {
	return &BananaHandler{svc: svc}
}

func (h *BananaHandler) GetBanana(
	ctx context.Context,
	in *bananav1.GetBananaRequest,
) (*bananav1.GetBananaResponse, error) {
	panic("implement me")
}

func (h *BananaHandler) UpdateBanana(
	ctx context.Context,
	in *bananav1.UpdateBananaRequest,
) (*bananav1.UpdateBananaResponse, error) {
	panic("implement me")
}

func (h *BananaHandler) DeleteBanana(
	ctx context.Context,
	in *bananav1.DeleteBananaRequest,
) (*bananav1.DeleteBananaResponse, error) {
	panic("implement me")
}

func (h *BananaHandler) AddBanana(
	ctx context.Context,
	in *bananav1.AddBananaRequest,
) (*bananav1.AddBananaResponse, error) {
	panic("implement me")
}
