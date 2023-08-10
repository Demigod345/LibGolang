package controller

import (
	"LibGolang/pkg/models"
	// "LibGolang/pkg/types"
	"golang.org/x/exp/slices"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"strings"
	"time"
)

func getJWTCookie(writer http.ResponseWriter, request *http.Request, userId int, username string, isAdmin bool) http.Cookie {

	token, err := models.GenerateJWT(userId, username, isAdmin)
	if err != nil {
		log.Fatal(err)
	}

	cookie := http.Cookie{}
	cookie.Name = "Token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(365 * 24 * time.Hour)
	cookie.Secure = false
	cookie.HttpOnly = true
	cookie.Path = "/"

	return cookie
}

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path == "/" || request.URL.Path == "/login" {
			next.ServeHTTP(writer, request)
			return
		}
		cookie, err := request.Cookie("Token")
		if err != nil {
			if request.URL.Path == "/" || request.URL.Path == "/home" || request.URL.Path == "/login" || request.URL.Path == "/signup" {
				next.ServeHTTP(writer, request)
				return
			} else {
				http.Redirect(writer, request, "/login", http.StatusSeeOther)
				return
			}
		}
		tokenString := strings.TrimSpace(cookie.Value)
		fmt.Println(tokenString)
		fmt.Println(models.GetJWTSecretKey())
		// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//     _, ok := token.Method.(*jwt.SigningMethodECDSA)
		//     if !ok {
		//        writer.WriteHeader(http.StatusUnauthorized)
		//        _, err := writer.Write([]byte("You're Unauthorized!"))
		//        if err != nil {
		//           return nil, err

		//        }
		//     }
		//     return "", nil

		//  })
		VerifyJWT(tokenString)

		next.ServeHTTP(writer, request)
	})
}


func VerifyJWT(JWTtoken string) {

	tokenArr := strings.Split(JWTtoken, ".")
	fmt.Println(tokenArr[1])
	fmt.Println(tokenArr[2])

	decodedJSON, err := base64.StdEncoding.DecodeString(tokenArr[1])
	fmt.Println(decodedJSON)

	// claims := jwt.MapClaims{}

	token, err := jwt.Parse(JWTtoken, func(token *jwt.Token) (interface{}, error) {
		return []byte(models.GetJWTSecretKey()), nil
	})

	if err != nil {
		fmt.Println(err)
	}

	if !token.Valid {
		fmt.Println("invalid token")
	}

	if token.Valid {
		fmt.Println("VAlid Token")
	}
}

// makkar



func DecodeJWT(tokenStr string) (jwt.MapClaims, error) {
	secretKey := models.GetJWTSecretKey()

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil

}

func ValidateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		goThroughUrls := []string{"/", "/signup", "/login"}

		if slices.Contains(goThroughUrls, r.URL.Path) || strings.Split(r.URL.Path, "/")[1] == "static" {
			next.ServeHTTP(w, r)
			return
		}


		cookie, err := r.Cookie("Token")
		token := cookie.Value
		fmt.Println("domya")
		fmt.Println(token)
		fmt.Println(models.GetJWTSecretKey())

		if err == nil {
			token := cookie.Value
			claims, err := DecodeJWT(token)
			if err != nil {
				fmt.Println(claims)
				fmt.Println("Here")
				fmt.Println(err)
				// http.Redirect(w, r, "/login", http.StatusSeeOther)
			}
			next.ServeHTTP(w, r)
			// return
		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	})
}
