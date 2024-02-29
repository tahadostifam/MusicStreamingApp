package jwt_manager

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

const JwtDefaultDuration = "5s"

var jwtManager = NewJwtManager("sample_secret", JwtDefaultDuration)

const TokenType = "access_token"

var StaticID = primitive.NewObjectID()

func TestJWTGenerateAndVerifyAccessToken(t *testing.T) {
	token, tokenErr := jwtManager.Generate(TokenType, StaticID.String())

	assert.Empty(t, tokenErr)
	assert.NotEmpty(t, token)

	cases := []struct {
		name  string
		token string
		err   error
	}{
		{
			name:  "valid",
			token: token,
			err:   nil,
		},
		{
			name:  "not_valid",
			token: "akdmakldmakldmaldalkdm",
			err:   ErrInvalidToken,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			_, verifyErr := jwtManager.Verify(tt.token, TokenType)

			assert.Equal(t, verifyErr, tt.err)
		})
	}
}
