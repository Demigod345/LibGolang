package middleware

import (
	"LibGolang/pkg/models"
	"LibGolang/pkg/types"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/exp/slices"
	"log"
	"net/http"
)

type contextKey string

const (
	UserIdContextKey   = contextKey("UserId")
	IsAdminContextKey  = contextKey("IsAdmin")
	UsernameContextKey = contextKey("Username")
)

func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		goThroughUrls := []string{"/", "/signup", "/login", "/403", "/500", "/loginAdmin", "/loginUser"}
		if slices.Contains(goThroughUrls, request.URL.Path) {
			next.ServeHTTP(writer, request)
			return
		}
		cookie, err := request.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Redirect(writer, request, "/403", http.StatusSeeOther)
				return
			}
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
		ctx := context.WithValue(request.Context(), UserIdContextKey, claims.UserId)
		ctx = context.WithValue(ctx, IsAdminContextKey, claims.IsAdmin)
		ctx = context.WithValue(ctx, UsernameContextKey, claims.Username)
		request = request.WithContext(ctx)

		next.ServeHTTP(writer, request)

	})
}

func RoleMiddleware(isAdminAuth bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			isAdmin := request.Context().Value(IsAdminContextKey).(bool)
			userId := request.Context().Value(UserIdContextKey).(int)
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
