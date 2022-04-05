package pin

import "context"

type Handler interface {
	Init(ctx context.Context) error
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
