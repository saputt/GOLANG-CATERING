package auth

import (
	"catering-api/internal/httpx"
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const userContextKey contextKey = "authUser"

type AuthUser struct {
	Id string
	Email string
	Role UserRole
}

func Middleware(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				httpx.WriteError(w, http.StatusUnauthorized, "Auth header cannot be empty")
				return 
			}

			if !strings.HasPrefix(authHeader, "Bearer "){
				httpx.WriteError(w, http.StatusUnauthorized, "Invalid authorization format")
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			claims := &JWTClaims{}

			token, err := jwt.ParseWithClaims(tokenString, claims, func (token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid
				}	
				return []byte(jwtSecret), nil
			})

			if err != nil || !token.Valid {
				httpx.WriteError(w, http.StatusUnauthorized, "Invalid or expired token")
				return
			}

			authUser := AuthUser{
				Id: claims.Id,
				Email: claims.Email,
				Role: claims.Role,
			}

			ctx := context.WithValue(r.Context(), userContextKey, authUser)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	} 
}

func GetUserFromContext(ctx context.Context) (AuthUser, bool) {
	user, ok := ctx.Value(userContextKey).(AuthUser)
	return user, ok
}

func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
		user, ok := GetUserFromContext(r.Context())

		if !ok {
			httpx.WriteError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		if user.Role != "admin" {
			httpx.WriteError(w, http.StatusForbidden, "This only for admin")
			return
		}

		next.ServeHTTP(w, r)
	})
}