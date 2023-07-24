package controller

import (
	"github.com/raaaaaaaay86/go-otel-example/service"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"strconv"
)

type IUserController interface {
	FindUserById(w http.ResponseWriter, r *http.Request)
}

var _ IUserController = (*UserController)(nil)

type UserController struct {
	TracerProvider trace.TracerProvider
	UserService    service.IUserService
}

func NewUserController(tracerProvider trace.TracerProvider, userService service.IUserService) *UserController {
	return &UserController{
		TracerProvider: tracerProvider,
		UserService:    userService,
	}
}

func (u UserController) FindUserById(w http.ResponseWriter, r *http.Request) {
	newCtx, span := u.TracerProvider.Tracer("root.controller").Start(r.Context(), "FindUserById")
	defer span.End()

	w.Header().Add("Content-Type", "application/json")

	id, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := u.UserService.FindUserById(newCtx, uint(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userJson, err := user.ToJSON()

	w.Write([]byte(userJson))
}
