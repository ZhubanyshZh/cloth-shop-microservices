package middlewares

import (
	"context"
	"encoding/json"
	"github.com/ZhubanyshZh/go-project-service/internal/dto"
	"github.com/ZhubanyshZh/go-project-service/internal/utils"
	"net/http"
)

func ProductCreateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			utils.HandleError(w, err, "❌ Failed to parse multipart form", http.StatusBadRequest)
			return
		}

		productJson := r.FormValue("product")
		var product dto.ProductCreate
		if err := json.Unmarshal([]byte(productJson), &product); err != nil {
			utils.HandleError(w, err, "❌ Invalid product JSON", http.StatusBadRequest)
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
