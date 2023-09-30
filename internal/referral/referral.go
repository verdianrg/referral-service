package referral

import (
	"context"
	"net/http"
	"os"
	service "referralservice"
	"referralservice/gen/models"
	"referralservice/gen/restapi/operations/unique"
	"referralservice/internal/resource"
	"strings"

	"github.com/go-openapi/runtime/middleware"
)

func UniqueLinkHandler(rt *service.Runtime) func(r unique.PostUniqueLinkLinkParams, p *models.Principal) middleware.Responder {
	return func(r unique.PostUniqueLinkLinkParams, p *models.Principal) middleware.Responder {
		ctx := context.Background()

		user, err := resource.GetUserByEmail(rt, ctx, r.Data.Email)
		if err != nil {
			rt.Error().Println(err)
			return unique.NewGetUniqueLinkEmailDefault(http.StatusNotFound).WithPayload(&models.BaseResponse{
				Code:    http.StatusNotFound,
				Message: "referral is not exist!",
			})
		}

		prefix := os.Getenv("BASE_URL") + "unique-link/"
		if r.Link != strings.TrimPrefix(user.UniqueLink, prefix) {
			err = rt.SetError(http.StatusNotFound, "referral is not exist!")
			rt.Error().Println(err)
			return unique.NewGetUniqueLinkEmailDefault(http.StatusNotFound).WithPayload(&models.BaseResponse{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			})
		}

		user.Contribution++
		err = resource.UpdateUserByEmail(rt, ctx, r.Data.Email, user)
		if err != nil {
			rt.Error().Println(err)
			return unique.NewGetUniqueLinkEmailDefault(http.StatusInternalServerError).WithPayload(&models.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		return unique.NewGetUniqueLinkEmailOK().WithPayload(&models.BaseResponse{
			Code:    http.StatusOK,
			Message: "success claim unique link",
		})
	}
}
