package pin

import (
	"context"
	"fmt"

	"github.com/make-go-great/ioe-go"
)

type Handler interface {
	Init(ctx context.Context) error
	SearchPin(ctx context.Context) error
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Init(ctx context.Context) error {
	for length := 4; length <= 6; length++ {
		if err := h.service.InitPin(ctx, length); err != nil {
			return err
		}
	}

	return nil
}

func (h *handler) SearchPin(ctx context.Context) error {
	fmt.Printf("Input sha256: ")
	sha256 := ioe.ReadInput()

	pin, err := h.service.GetPinBySHA256(ctx, sha256)
	if err != nil {
		return err
	}

	fmt.Printf("Pin: %s\n", pin.Pin)
	return nil
}
