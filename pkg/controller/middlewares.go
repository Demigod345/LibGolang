package controller

import (
	"LibGolang/pkg/models"
	"LibGolang/pkg/types"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
)

type contextKey string

const (
	userIdContextKey   = contextKey("UserId")
	isAdminContextKey  = contextKey("IsAdmin")
	usernameContextKey = contextKey("Username")
)

func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path == "/" || request.URL.Path == "/login" || request.URL.Path == "/500" || request.URL.Path == "/403" || request.URL.Path == "/signup" || request.URL.Path == "/loginAdmin" || request.URL.Path == "/loginUser" {
			next.ServeHTTP(writer, request)
			return
		}
		cookie, err := request.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				// w.WriteHeader(http.StatusUnauthorized)
				http.Redirect(writer, request, "/403", http.StatusSeeOther)
				return
			}
			// For any other type of error, return a bad request status
			// w.WriteHeader(http.StatusBadRequest)
			http.Redirect(writer, request, "/403", http.StatusSeeOther)
			return
		}
		tokenString := cookie.Value
		claims := &types.Claims{}
		key, err := models.GetJWTSecretKey()
		jwtKey := []byte(key)

		if err != nil {
			log.Println(err)
			http.Redirect(writer, request, "/500", http.StatusSeeOther)
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Redirect(writer, request, "/403", http.StatusSeeOther)
				return
			}
			log.Println(err)
			http.Redirect(writer, request, "/500", http.StatusSeeOther)
			return
		}
		if !token.Valid {
			http.Redirect(writer, request, "/403", http.StatusSeeOther)
			return
		}
		ctx := context.WithValue(request.Context(), userIdContextKey, claims.UserId)
		ctx = context.WithValue(ctx, isAdminContextKey, claims.IsAdmin)
		ctx = context.WithValue(ctx, usernameContextKey, claims.Username)
		request = request.WithContext(ctx)

		next.ServeHTTP(writer, request)

	})
}

func RoleMiddleware(isAdminAuth bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			// Retrieve the role from the context (assuming you stored it during authentication)
			isAdmin := request.Context().Value(isAdminContextKey).(bool)
			userId := request.Context().Value(userIdContextKey).(int)
			isAdminDb, err := models.CheckAdmin(userId)

			if err != nil {
				log.Println(err)
				http.Redirect(writer, request, "/500", http.StatusSeeOther)
				return
			}

			if isAdmin == isAdminAuth && isAdmin == isAdminDb {
				next.ServeHTTP(writer, request)
			} else {
				http.Redirect(writer, request, "/403", http.StatusSeeOther)
				return
			}
		})
	}
}
