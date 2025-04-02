package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/ZhubanyshZh/go-project-service/internal/models"
	"github.com/ZhubanyshZh/go-project-service/internal/utils"
)

func ValidateProductMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var product models.ProductEdit

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "❌ Ошибка чтения тела запроса", http.StatusBadRequest)
			log.Fatal("❌ Ошибка чтения тела запроса", err.Error())
			return
		}

		if err := json.Unmarshal(body, &product); err != nil {
			http.Error(w, "❌ Неверный формат JSON", http.StatusBadRequest)
			log.Fatal("❌ Неверный формат JSON: ", err.Error())
			return
		}

		if err := utils.ValidateStruct(&product); err != nil {
			http.Error(w, "❌ Ошибка валидации: "+err.Error(), http.StatusBadRequest)
			log.Fatal("❌ Ошибка валидации:", err.Error())
			return
		}

		r.Body = io.NopCloser(bytes.NewReader(body))

		next.ServeHTTP(w, r)
	})
}
