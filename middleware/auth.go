package middleware

import (
	"context"
	"gored/commons/exception"
	"gored/commons/helper"
	"gored/commons/response"
	"gored/pkg/auth"
	"net/http"
	"strings"
)

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if len(authHeader) <= 0 {
			response.JsonEncoder[string](w, http.StatusUnauthorized, exception.ErrUnauthorized.Error(), false, "")
			return
		}

		fields := strings.Fields(authHeader)
		if fields[0] != "Bearer" {
			response.JsonEncoder[string](w, http.StatusUnauthorized, exception.ErrUnauthorized.Error(), false, "")
			return
		}

		token := helper.SplitToken(r)
		_, err := auth.User(context.Background(), token)
		if err != nil {
			response.JsonEncoder[string](w, http.StatusUnauthorized, exception.ErrUnauthorized.Error(), false, "")
			return
		}
		next.ServeHTTP(w, r)
	})

}
