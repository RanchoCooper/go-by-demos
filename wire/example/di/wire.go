package di

import (
	"context"
	"io"

	"github.com/google/wire"

	"github.com/RanchoCooper/go-by-demos/wire/example/db"
	"github.com/RanchoCooper/go-by-demos/wire/example/service"
)

func NewService(c *db.Config, m *service.MailConfig) (*service.Service, error) {
	wire.Build(service.NewService, service.MailSet, service.UserSet, db.NewDB)
	return &service.Service{}, nil
}

func InitGreeter(ctx context.Context, s []string, w io.Writer) (*service.Greeter, error) {
	wire.Build(service.GreeterSet)
	return nil, nil
}
