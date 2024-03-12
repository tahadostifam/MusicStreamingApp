package random

import (
	"crypto/rand"
	"encoding/hex"
	"math"
	"math/big"
	"strconv"
)

const OtpLength = 6

func GenerateOTP() int {
	min := int64(math.Pow(10, OtpLength-1))
	max := int64(math.Pow(10, OtpLength)) - 1

	randomNumber, err := rand.Int(rand.Reader, big.NewInt(max-min))
	if err != nil {
		panic(err)
	}

	number := int(randomNumber.Int64()) + int(min)

	if len(strconv.Itoa(number)) != OtpLength {
		number = GenerateOTP()
	}

	return number
}

func GenerateRandomFileName(n int) string {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}

	return hex.EncodeToString(bytes)
}
