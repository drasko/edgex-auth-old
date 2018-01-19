package auth

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const Issuer string = "mainflux"

var secretKey string = "mainflux-api-key"

// SetSecretKey sets the secret key that will be used for decoding and encoding
// tokens. If not invoked, a default key will be used instead.
func SetSecretKey(key string) {
	secretKey = key
}

// DecodeJWT decodes jwt token
func DecodeJwt(key string) (*jwt.StandardClaims, error) {
	claims := jwt.StandardClaims{}

	token, err := jwt.ParseWithClaims(
		key,
		&claims,
		func(val *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)

	// Validate the token and return the custom claims
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// CreateKey creates a JSON Web Token with a given subject.
func CreateKey(subject string) (string, error) {
	claims := jwt.StandardClaims{
		Issuer:   Issuer,
		IssuedAt: time.Now().UTC().Unix(),
		Subject:  subject,
	}

	key := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	raw, err := key.SignedString([]byte(secretKey))
	if err != nil {
		return "", &AuthError{http.StatusInternalServerError, err.Error()}
	}

	return raw, nil
}
