package middlewares

import (
	"context"
	"github.com/ZhubanyshZh/go-project-service/internal/dto"
	"github.com/ZhubanyshZh/go-project-service/internal/utils"
	"net/http"
)

func ProductUpdateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var product dto.ProductUpdate
		if !utils.DecodeJSONRequest(w, r, &product) {
			return
		}

		if err := utils.ValidateStruct(product); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "validatedProduct", product)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
