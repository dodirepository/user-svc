package middleware

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

var (
	secretKey = []byte("your_secret_key") // Replace with your own secret key
)

type Claims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email,omitempty"`
	Phone  string `json:"phone,omitempty"`
	jwt.StandardClaims
}

// GenerateToken generates a JWT token with user information
func GenerateToken(userID int64, email, phone string) (string, error) {
	claims := Claims{
		UserID: userID,
		Email:  email,
		Phone:  phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, "Invalid token signature", http.StatusUnauthorized)
				return
			}
			logrus.WithError(err).Error()
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Set the username in the request context
		userID := strconv.FormatInt(claims.UserID, 10)
		r.Header.Set("x-user-id", userID)
		ctx := context.WithValue(r.Context(), "x-user-id", claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
