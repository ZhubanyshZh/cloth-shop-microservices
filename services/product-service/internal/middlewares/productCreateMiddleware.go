package middlewares

import (
	"context"
	"github.com/ZhubanyshZh/go-project-service/internal/dto"
	"github.com/ZhubanyshZh/go-project-service/internal/utils"
	"net/http"
	"strconv"
)

func ProductCreateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(10 << 20) // 10MB
		if err != nil {
			http.Error(w, "Failed to parse multipart form", http.StatusBadRequest)
			return
		}

		priceStr := r.FormValue("price")
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			http.Error(w, "Invalid price value", http.StatusBadRequest)
			return
		}

		product := dto.ProductCreate{
			Name:        r.FormValue("product_name"),
			Description: r.FormValue("description"),
			Price:       price,
			Images:      r.MultipartForm.File["images"],
		}

		if err := utils.ValidateStruct(product); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "validatedProduct", product)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
