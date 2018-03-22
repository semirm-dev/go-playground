package token

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Model for JWT
type Model struct {
	SecretKey []byte
	Token     string
}

// New will create new token Model
func New() *Model {
	m := new(Model)

	m.SecretKey = secretKey()

	return m
}

// Generate token string
func (m *Model) Generate(c map[string]string) error {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	for k, v := range c {
		claims[k] = v
	}

	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()

	tokenStr, err := token.SignedString(m.SecretKey)

	if err != nil {
		return err
	}

	m.Token = tokenStr

	return nil
}

// Valid will check if token from request is valid or not
func (m *Model) Valid(tokenStr string) bool {
	if token, _ := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Something went wrong while parsing")
		}

		return m.SecretKey, nil
	}); token != nil {
		if _, claimsOk := token.Claims.(jwt.MapClaims); claimsOk && token.Valid {
			return true
		}
	}

	return false
}

func secretKey() []byte {
	return []byte("secret123")
}
