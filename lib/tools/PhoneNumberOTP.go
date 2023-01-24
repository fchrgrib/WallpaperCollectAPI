package tools

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

func PhoneNumberOTP(phoneNumber int) error {

	return nil
}

func getRandNumber() (string, error) {
	nBig, e := rand.Int(rand.Reader, big.NewInt(8999))
	if e != nil {
		return "", e
	}
	return strconv.FormatInt(nBig.Int64()+1000, 10), nil
}
