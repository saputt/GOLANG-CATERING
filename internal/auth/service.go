package auth

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailAlreadyExist = errors.New("Email already exist")
	ErrInvalidCredential = errors.New("Invalid credential")
	ErrInvalidInput = errors.New("Invalid input")
)

type JWTClaims struct {
	Id string `json:"id"`
	Email string `json:"email"`
	Role UserRole `json:"role"`

	jwt.RegisteredClaims
}

type Service struct {
	repo *Repository
	jwtSecret string
	jwtExpiresHour int
}

func NewService(repo *Repository, jwtSecret string, jwtExpiresHour int) *Service{
	return &Service{
		repo: repo,
		jwtSecret: jwtSecret,
		jwtExpiresHour: jwtExpiresHour,
	}
}

func (s *Service) generateToken(user User) (string, error) {
	claims := JWTClaims{
		Id: user.Id,
		Email: user.Email,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: user.Id,
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(s.jwtExpiresHour) * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.jwtSecret))
}

func (s *Service) Login(ctx context.Context, req LoginRequest) (LoginResponse, error) {
	req.Email = strings.ToLower(strings.TrimSpace(req.Email))
	if req.Email == "" || req.Password == ""  {
		return LoginResponse{}, ErrInvalidInput
	}

	user, err := s.repo.FindUserByEmail(ctx, req.Email)
	if(err != nil) {
		return LoginResponse{}, ErrInvalidCredential
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return LoginResponse{}, ErrInvalidCredential
	}

	
	token, err := s.generateToken(user)
	fmt.Println(err)
	if err != nil {
		return LoginResponse{}, err
	}


	return LoginResponse{
		Token: token,
		User: UserResponse{
			Id: user.Id,
			Username: user.Username,
			Email: user.Email,
			Role: user.Role,
		},
	}, nil
}

func (s *Service) Register(ctx context.Context, req RegisterRequest) (RegisterResponse, error) {
	req.Email = strings.ToLower(strings.TrimSpace(req.Email))
	req.Username = strings.TrimSpace(req.Username)
	
	if req.Email == "" || req.Password == "" || req.Username == "" {
		return RegisterResponse{}, ErrInvalidInput
	}

	_, err := s.repo.FindUserByEmail(ctx, req.Email)
	if err == nil {
		return RegisterResponse{}, ErrEmailAlreadyExist
	} 

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return RegisterResponse{}, err
	}

	user := User{
		Id: uuid.NewString(),
		Username: req.Username,
		Email: req.Email,
		Password: string(hashedPassword),
		Role: "user",
	}

	createdUser, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return RegisterResponse{}, err
	}

	return RegisterResponse{
		User: UserResponse{
			Id: createdUser.Id,
			Username: createdUser.Username,
			Email: createdUser.Email,
			Role: createdUser.Role,
		},
	}, nil
}