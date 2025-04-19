package middlewares

import (
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/dtos"
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthReqMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userLoginReq dtos.AuthReq
		if !utils.DecodeJSONRequest(c.Writer, c.Request, &userLoginReq) {
			return
		}

		if err := utils.ValidateStruct(userLoginReq); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.Set("validatedUserReq", userLoginReq)
		c.Next()
	}
}
