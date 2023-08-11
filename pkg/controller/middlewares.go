package controller

import (
	"LibGolang/pkg/models"
	"LibGolang/pkg/types"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

type contextKey string

const (
	userIdContextKey   = contextKey("UserId")
	isAdminContextKey  = contextKey("IsAdmin")
	usernameContextKey = contextKey("Username")
)

var jwtKey = []byte(models.GetJWTSecretKey())

func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" || r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				// w.WriteHeader(http.StatusUnauthorized)
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
			// For any other type of error, return a bad request status
			// w.WriteHeader(http.StatusBadRequest)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		tknStr := c.Value
		claims := &types.Claims{}

		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				// w.WriteHeader(http.StatusUnauthorized)
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
			// w.WriteHeader(http.StatusBadRequest)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		if !tkn.Valid {
			// w.WriteHeader(http.StatusUnauthorized)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		ctx := context.WithValue(r.Context(), userIdContextKey, claims.UserId)
		ctx = context.WithValue(ctx, isAdminContextKey, claims.IsAdmin)
		ctx = context.WithValue(ctx, usernameContextKey, claims.Username)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

	})
}

func RoleMiddleware(isAdminAuth bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Retrieve the role from the context (assuming you stored it during authentication)
			isAdmin := r.Context().Value(isAdminContextKey).(bool)
			userId := r.Context().Value(userIdContextKey).(int)
			isAdminDb, err := models.CheckAdmin(userId)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(isAdmin, isAdminAuth, isAdminDb)

			if isAdmin == isAdminAuth && isAdmin == isAdminDb {
				// If the user's role matches, allow access to the next handler
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "Unauthorized access", http.StatusForbidden)
				return
			}
		})
	}
}