package service

import (
	"context"
	"github.com/raaaaaaaay86/go-otel-example/entity"
	"github.com/raaaaaaaay86/go-otel-example/repository"
	"go.opentelemetry.io/otel/trace"
)

type IUserService interface {
	FindUserById(ctx context.Context, id uint) (*entity.User, error)
}

var _ IUserService = (*UserService)(nil)

type UserService struct {
	TracerProvider trace.TracerProvider
	UserRepository repository.IUserRepository
}

func NewUserService(tracerProvider trace.TracerProvider, userRepository repository.IUserRepository) *UserService {
	return &UserService{
		TracerProvider: tracerProvider,
		UserRepository: userRepository,
	}
}

func (u UserService) FindUserById(ctx context.Context, id uint) (*entity.User, error) {
	newCtx, span := u.TracerProvider.Tracer("root.service").Start(ctx, "FindUserById")
	defer span.End()

	user, err := u.UserRepository.FindById(newCtx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
