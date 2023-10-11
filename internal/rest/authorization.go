package rest

import (
	"net/http"
	service "referralservice"
	"referralservice/gen/models"
	"referralservice/gen/restapi/operations"
	"referralservice/internal/utils"
)

func Authorization(rt *service.Runtime, api *operations.ReferralServerAPI) {
	api.KeyAuth = func(token string) (p *models.Principal, err error) {
		claims, err := utils.ParseAndCheckToken(token)
		if err != nil {
			return nil, rt.SetError(http.StatusUnauthorized, "Forbidden: insufficient API key privileges")

		}

		if claims.Role != "" {
			p.Role = claims.Role
		}

		return
	}

	api.HasRoleAuth = func(token string) (p *models.Principal, err error) {
		claims, err := utils.ParseAndCheckToken(token)
		if err != nil || claims.Role == "" {
			return nil, rt.SetError(http.StatusUnauthorized, "Forbidden: insufficient API key privileges")

		}

		if claims.Role != "contributor" {
			return nil, rt.SetError(http.StatusUnauthorized, "Forbidden: insufficient API key privileges")
		}

		return
	}

}
