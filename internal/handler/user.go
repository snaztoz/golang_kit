package handler

import (
	"fmt"
	"net/http"
	"template/internal/appctx"
	"template/internal/params"
	"template/internal/usecase"
	"template/utils/json"
	"template/utils/validator"
	"time"

	"github.com/sirupsen/logrus"
)

type user struct {
	handler Handler
	usecase usecase.UserUsecase
	name    string
}

type UserHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler() UserHandler {
	return &user{
		usecase: usecase.NewUserUsecase(),
		name:    "USER HANDLER",
	}
}

func (u *user) Create(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][Create] is executed", u.name))
	// d := appctx.Data{
	// 	Request: r,
	// }

	var param params.UserCreateParam
	ctx := appctx.NewResponse()

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error("Cannot decode json")
		ctx = ctx.WithErrors(err.Error())
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
	}

	fmt.Printf("Debug: %v", param)

	if len(ctx.Errors) > 0 {
		u.handler.Response(w, *ctx, time.Now())
		return
	}

	resp := u.usecase.Create(param)
	u.handler.Response(w, resp, time.Now())
}
