package user

import (
	"context"
	"net/http"
	service "referralservice"
	"referralservice/gen/models"
	"referralservice/gen/restapi/operations/user"
	"referralservice/internal/resource"
	"referralservice/internal/utils"

	"github.com/go-openapi/runtime/middleware"
)

func CreateUserHandler(rt *service.Runtime) func(r user.PostV1UserParams) middleware.Responder {
	return func(r user.PostV1UserParams) middleware.Responder {
		ctx := context.Background()

		data := &models.User{
			UserData: r.Data.UserData,
		}

		uniqueLink := utils.GenerateUniqueLink(r.Data.Email)
		data.UniqueLink = uniqueLink

		err := CreateUser(rt, ctx, data)
		if err != nil {
			rt.Error().Println(err)
			errData := rt.GetError(err)
			return user.NewPostV1UserDefault(int(errData.Code())).WithPayload(&models.BaseResponse{
				Code:    errData.Code(),
				Message: err.Error(),
			})
		}

		return user.NewPostV1UserCreated().WithPayload(&models.User{
			UserData: data.UserData,
		})
	}
}

func CreateUser(rt *service.Runtime, ctx context.Context, data *models.User) error {
	err := ValidateUser(rt, ctx, data.Email)
	if err != nil {
		return err
	}

	err = resource.CreateUser(rt, ctx, data)
	return err
}

func ValidateUser(rt *service.Runtime, ctx context.Context, email string) error {
	_, err := resource.GetUserByEmail(rt, ctx, email)
	if err == nil {
		return rt.SetError(http.StatusBadRequest, "email already exist!")
	}

	return nil
}
