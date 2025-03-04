package jwt

import (
	"errors"
	"time"

	driJWT "github.com/dgrijalva/jwt-go"
)

// Define any service related to JWT
type JWTService interface {
	GenerateToken(userId string, role string, email string) (t string, err error)
	ParseToken(tokenString string) (claims JwtCustomClaim, err error)
}

// Define claim included in the JWT
type JwtCustomClaim struct {
	UserId string
	Role   string
	Email  string
	driJWT.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
	expired   int
}

func NewJWTService(secretKey, issuer string, expired int) JWTService {
	return &jwtService{
		issuer:    issuer,
		secretKey: secretKey,
		expired:   expired,
	}
}

// Function to generate JWT
func (j *jwtService) GenerateToken(userId string, role string, email string) (t string, err error) {
	// Define the claim
	claims := &JwtCustomClaim{
		userId,
		role,
		email,
		driJWT.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(time.Hour * time.Duration(j.expired)).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().UTC().Unix(),
		},
	}

	// Generate token and sign it
	token := driJWT.NewWithClaims(driJWT.SigningMethodHS256, claims)
	t, err = token.SignedString([]byte(j.secretKey))
	return
}

// Function to parse JWT
func (j *jwtService) ParseToken(tokenString string) (claims JwtCustomClaim, err error) {
	if token, err := driJWT.ParseWithClaims(tokenString, &claims, func(token *driJWT.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	}); err != nil || !token.Valid {
		return JwtCustomClaim{}, errors.New("token is not valid")
	}

	return
}
