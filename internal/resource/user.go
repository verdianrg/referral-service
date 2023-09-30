package resource

import (
	"context"
	service "referralservice"
	"referralservice/gen/models"
)

func CreateUser(rt *service.Runtime, ctx context.Context, user *models.User) (err error) {
	// var user models.User
	model := rt.DB().Model(&user)

	err = model.Create(&user).Error

	return
}

func GetUserLogin(rt *service.Runtime, ctx context.Context, email, password string) (user *models.User, err error) {
	model := rt.DB().Model(&user)
	model = model.First(&user, "email = ? AND password = ?", email, password)
	if model.Error != nil {
		return
	}

	return
}

func GetUserByID(rt *service.Runtime, ctx context.Context, id int64) (user *models.User, err error) {
	model := rt.DB().Model(&user)
	model = model.First(&user, id)

	if err = model.Error; err != nil {
		return
	}

	return
}

func GetUserByEmail(rt *service.Runtime, ctx context.Context, email string) (user *models.User, err error) {
	err = rt.DB().Where("email = ?", email).First(&user).Error
	if err != nil {
		return
	}

	return
}

func UpdateUserByEmail(rt *service.Runtime, ctx context.Context, email string, data *models.User) (err error) {
	model := rt.DB().Model(&models.User{})

	err = model.Where("email = ?", email).Update("contribution", data.Contribution).Error
	if err != nil {
		return
	}

	return
}
