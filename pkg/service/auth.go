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

var jwtSecret = []byte("ARTEM363IVT")
var jwtRefreshSecret = []byte("ARTEM_363_IVT")

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	Login  string `json:"login"`
	UserID int    `json:"user_id"`
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

func (s *AuthService) GenerateToken(login, password string) (map[string]string, error) {
	user, err := s.repo.GetUser(login)

	if err != nil {
		return nil, err
	}

	if !checkPasswordHash(password, user.Password) {
		return nil, errors.New("invalid password or Email")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserID: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(120 * time.Minute)), //time token
			Issuer:    "Artem Medvedev",
		},
	})

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to sign the token: %v", err)
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &RefreshClaims{
		Login:  login,
		UserID: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), //time token
			Issuer:    "Artem Medvedev",
		},
	})

	signedRefreshToken, err := refreshToken.SignedString(jwtRefreshSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to sign the token: %v", err)
	}

	if err = s.repo.UpdateRefreshToken(signedRefreshToken, user.Id); err != nil {
		return nil, err
	}

	data := make(map[string]string)
	data["access_token"] = signedToken
	data["refresh_token"] = signedRefreshToken

	return data, nil
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

func (s *AuthService) ParseRefreshToken(tokenString string) (jewelrymodel.User, error) {

	token, err := jwt.ParseWithClaims(tokenString, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtRefreshSecret, nil
	})

	var user jewelrymodel.User

	if err != nil {
		return user, fmt.Errorf("error parsing token: %v", err)
	}

	claims, ok := token.Claims.(*RefreshClaims)
	if !ok || !token.Valid {
		return user, errors.New("invalid token")
	}

	user, err = s.repo.GetUser(claims.Login)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *AuthService) ReGenerateToken(user jewelrymodel.User) (map[string]string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserID: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(120 * time.Minute)), //time token
			Issuer:    "Artem Medvedev",
		},
	})

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to sign the token: %v", err)
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &RefreshClaims{
		Login:  user.Login,
		UserID: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), //time token
			Issuer:    "Artem Medvedev",
		},
	})

	signedRefreshToken, err := refreshToken.SignedString(jwtRefreshSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to sign the token: %v", err)
	}

	if err = s.repo.UpdateRefreshToken(signedRefreshToken, user.Id); err != nil {
		return nil, err
	}

	data := make(map[string]string)
	data["access_token"] = signedToken
	data["refresh_token"] = signedRefreshToken

	return data, nil
}

func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}
