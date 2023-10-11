package login

import (
	"context"
	"net/http"
	service "referralservice"
	"referralservice/gen/models"
	"referralservice/gen/restapi/operations/dummy"
	"referralservice/gen/restapi/operations/login"
	"referralservice/internal/resource"
	"referralservice/internal/utils"

	"github.com/go-openapi/runtime/middleware"
)

func Login(rt *service.Runtime) func(r login.PostV1LoginParams) middleware.Responder {
	return func(r login.PostV1LoginParams) middleware.Responder {
		user, err := resource.GetUserLogin(rt, context.Background(), *r.Data.Email, *r.Data.Password)
		if err != nil {
			rt.Error().Println(err)
			return login.NewPostV1LoginDefault(http.StatusNotFound).WithPayload(&models.BaseResponse{
				Code:    http.StatusNotFound,
				Message: "user not found!",
			})
		}

		tokenString, err := utils.CreateToken(user.Email, user.Role)
		if err != nil {
			rt.Error().Println(err)
			return login.NewPostV1LoginDefault(http.StatusInternalServerError).WithPayload(&models.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		return login.NewPostV1LoginCreated().WithPayload(&models.LoginResponse{
			Email: user.Email,
			Role:  user.Role,
			Token: tokenString,
		})
	}
}

func Dummy(rt *service.Runtime) func(r dummy.GetV1DummyParams, p *models.Principal) middleware.Responder {
	return func(r dummy.GetV1DummyParams, p *models.Principal) middleware.Responder {
		return dummy.NewGetV1DummyOK().WithPayload(&models.BaseResponse{
			Code:    http.StatusOK,
			Message: "dummy success",
		})
	}
}
