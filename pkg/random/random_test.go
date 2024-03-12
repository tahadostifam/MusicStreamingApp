package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateOTP(t *testing.T) {
	otp := GenerateOTP()
	assert.NotZero(t, otp)

	t.Log(otp)
}

func TestGenerateRandomFileName(t *testing.T) {
	username := GenerateRandomFileName(20)

	t.Log(username)
}
