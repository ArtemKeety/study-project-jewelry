package service

import (
	"curs/jewelrymodel"
	"curs/pkg/repository"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const salt = "ARTEM363IVT"

var jwtSecret = []byte("ARTEM363IVT")

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *AuthService) CreateUser(user jewelrymodel.User) (int, error) {
	if checkUser, _ := s.repo.GetUser(user.Login); checkUser.Id > 0 {
		return -1, errors.New("user already exists")
	}
	user.Password = HashPassword(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(login, password string) (string, error) {
	user, err := s.repo.GetUser(login)

	if err != nil {
		return "", err
	}

	if !checkPasswordHash(password, user.Password) {
		return "", errors.New("invalid password or Email")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserID: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), //time token
			Issuer:    "Artem Medvedev",
		},
	})

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign the token: %v", err)
	}

	return signedToken, nil
}

func (s *AuthService) ParseToken(tokenString string) (int, error) {

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return 0, fmt.Errorf("error parsing token: %v", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}

	return claims.UserID, nil
}

func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}
