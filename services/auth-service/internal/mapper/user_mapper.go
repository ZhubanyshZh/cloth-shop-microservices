package mapper

import (
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/dtos"
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/models"
	"github.com/jinzhu/copier"
)

func ToMeResponse(user *models.User) *dtos.MeResponse {
	var res dtos.MeResponse
	_ = copier.Copy(&res, &user)
	res.Id = user.ID.String()
	res.Role = string(user.Role)
	return &res
}
