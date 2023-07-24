package repository

import (
	"context"
	"github.com/raaaaaaaay86/go-otel-example/entity"
	"go.opentelemetry.io/otel/trace"
	"time"
)

type IUserRepository interface {
	FindById(ctx context.Context, id uint) (*entity.User, error)
}

var _ IUserRepository = (*UserRepository)(nil)

type UserRepository struct {
	TracerProvider trace.TracerProvider
}

func NewUserRepository(tracerProvider trace.TracerProvider) *UserRepository {
	return &UserRepository{TracerProvider: tracerProvider}
}

func (u UserRepository) FindById(ctx context.Context, id uint) (*entity.User, error) {
	_, span := u.TracerProvider.Tracer("root.repository").Start(ctx, "FindById")
	defer span.End()

	time.Sleep(3 * time.Second) // Simulate slow database query time

	return entity.NewUser(id, "John Doe", "johndoe@email.com"), nil
}
