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

// func Refresh(w http.ResponseWriter, r *http.Request) {
// 	// (BEGIN) The code until this point is the same as the first part of the `Welcome` route
// 	c, err := r.Cookie("token")
// 	if err != nil {
// 		if err == http.ErrNoCookie {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	tknStr := c.Value
// 	claims := &Claims{}
// 	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})
// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	if !tkn.Valid {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
// 	// (END) The code until this point is the same as the first part of the `Welcome` route

// 	// We ensure that a new token is not issued until enough time has elapsed
// 	// In this case, a new token will only be issued if the old token is within
// 	// 30 seconds of expiry. Otherwise, return a bad request status
// 	if time.Until(claims.ExpiresAt.Time) > 30*time.Second {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// Now, create a new token for the current use, with a renewed expiration time
// 	expirationTime := time.Now().Add(5 * time.Minute)
// 	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString(jwtKey)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	// Set the new token as the users `token` cookie
// 	http.SetCookie(w, &http.Cookie{
// 		Name:    "token",
// 		Value:   tokenString,
// 		Expires: expirationTime,
// 	})
// }
