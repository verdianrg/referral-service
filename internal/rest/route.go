package rest

import (
	service "referralservice"
	"referralservice/gen/models"
	"referralservice/gen/restapi/operations"
	"referralservice/gen/restapi/operations/dummy"
	"referralservice/gen/restapi/operations/login"
	"referralservice/gen/restapi/operations/unique"
	"referralservice/gen/restapi/operations/user"

	login_domain "referralservice/internal/login"
	referral_domain "referralservice/internal/referral"
	user_domain "referralservice/internal/user"

	"github.com/go-openapi/runtime/middleware"
)

func Route(rt *service.Runtime, api *operations.ReferralServerAPI) {
	rt.Info().Println("Initialize Route")

	api.LoginPostV1LoginHandler = login.PostV1LoginHandlerFunc(func(pvlp login.PostV1LoginParams) middleware.Responder {
		return login_domain.Login(rt)(pvlp)
	})

	api.DummyGetV1DummyHandler = dummy.GetV1DummyHandlerFunc(func(gvdp dummy.GetV1DummyParams, p *models.Principal) middleware.Responder {
		return login_domain.Dummy(rt)(gvdp, p)
	})

	api.UserPostV1UserHandler = user.PostV1UserHandlerFunc(func(pvup user.PostV1UserParams) middleware.Responder {
		return user_domain.CreateUserHandler(rt)(pvup)
	})

	api.UserGetV1UserIDHandler = user.GetV1UserIDHandlerFunc(func(gvui user.GetV1UserIDParams, p *models.Principal) middleware.Responder {
		return user_domain.GetUserHandler(rt)(gvui, p)
	})

	api.UniquePostUniqueLinkLinkHandler = unique.PostUniqueLinkLinkHandlerFunc(func(pullp unique.PostUniqueLinkLinkParams, p *models.Principal) middleware.Responder {
		return referral_domain.UniqueLinkHandler(rt)(pullp, p)
	})
}
