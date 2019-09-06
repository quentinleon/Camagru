package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

var secret = "very secret secret"

func generateToken(json string) string {
	msg := base64.StdEncoding.EncodeToString([]byte(json))

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(msg))
	sig := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return msg + "." + sig
}

func verifyToken(token string) bool {
	splitToken := strings.Split(token, ".")
	msg := splitToken[0]
	actsig := splitToken[1]

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(msg))
	expsig := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return actsig == expsig
}

func getTokenContent(token string) string {
	splitToken := strings.Split(token, ".")
	msg, err := base64.StdEncoding.DecodeString(splitToken[0])
	if err != nil {
		return "error"
	}
	return string(msg)
}
