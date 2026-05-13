package auth

import (
	"catering-api/internal/httpx"
	"context"
	"encoding/json"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
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

func RateLimiter(rdb *redis.Client, limit int, duration time.Duration) func(http.Handler) http.Handler {
	return func (next http.Handler) http.Handler  {
		return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
			ip, _, _ := net.SplitHostPort(r.RemoteAddr)
			key := "rate_limit:" + ip

			count, err := rdb.Incr(r.Context(), key).Result()
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			if count == 1 {
				rdb.Expire(r.Context(), key, duration)
			}

			if count > int64(limit) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusTooManyRequests)

				json.NewEncoder(w).Encode(map[string]string{
					"error" : "Too many requests dude",
				})
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}