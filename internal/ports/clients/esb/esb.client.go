package esb_ports

import (
	"context"
	Errors "gin-framework-boilerplate/pkg/errors"
)

type ESBClient interface {
	Sample(ctx context.Context) (GeneralResponseDTO, Errors.CustomError)
}
