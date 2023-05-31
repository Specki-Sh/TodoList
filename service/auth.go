package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"
	u "todolist/domain/use_cases"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
}

func NewAuthService(userUseCase u.UserUseCase) *AuthService {
	return &AuthService{userUseCase: userUseCase}
}

type AuthService struct {
	userUseCase u.UserUseCase
}

func (a *AuthService) Authenticate(email, password string) (string, error) {
	password = generatePasswordHash(password)
	fmt.Println(email, password)
	user, err := a.userUseCase.ShowByEmailAndPassword(email, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
		user.Role,
	})

	return token.SignedString([]byte(signingKey))
}

func (a *AuthService) ParseToken(accessToken string) (int, string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, "", errors.New("token claims are not of type *tokenClaims")
	}

	if time.Now().Unix() > claims.ExpiresAt {
		return 0, "", errors.New("token expired")
	}

	return claims.UserId, claims.Role, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
