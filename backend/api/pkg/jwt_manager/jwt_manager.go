package jwt_manager

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var JwtAlgorithm = jwt.SigningMethodHS512

// define errors.
var (
	ErrInvalidToken            = errors.New("invalid token")
	ErrInvalidTokenType        = errors.New("invalid token type")
	ErrUnexpectedSigningMethod = errors.New("unexpected token signing method")
)

type JwtClaims struct {
	TokenType string
	UserID    string
	CreatedAt time.Time
	jwt.Claims
}

type JwtManager struct {
	secretKey string
	ttl       time.Duration
}

func NewJwtManager(secretKey string, expire time.Duration) *JwtManager {
	return &JwtManager{
		secretKey: secretKey,
		ttl:       expire,
	}
}

func (m *JwtManager) Generate(tokenType string, userID string) (string, error) {
	createdAt := time.Now()
	claims := &JwtClaims{UserID: userID, CreatedAt: createdAt}

	token := jwt.NewWithClaims(JwtAlgorithm, claims)
	return token.SignedString([]byte(m.secretKey))
}

func (m *JwtManager) Verify(userToken string, tokenType string) (*JwtClaims, error) {
	claims := &JwtClaims{}

	token, err := jwt.ParseWithClaims(userToken, claims,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrUnexpectedSigningMethod
			}

			return []byte(m.secretKey), nil
		},
	)
	if err != nil {
		return nil, ErrInvalidToken
	}

	if token.Valid {
		if claims.TokenType != tokenType {
			return nil, ErrInvalidTokenType
		}

		return claims, nil
	}

	return nil, ErrInvalidToken
}
