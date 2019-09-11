package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

var secret = "very secret secret"

func generateToken(content string) string {
	msg := base64.StdEncoding.EncodeToString([]byte(content))

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(msg))
	sig := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return msg + "." + sig
}

func getTokenContent(token string) string {
	splitToken := strings.Split(token, ".")
	msg, err := base64.StdEncoding.DecodeString(splitToken[0])
	if err != nil || len(splitToken) != 2 {
		//TODO log malformed token
		return "error"
	}
	return string(msg)
}

func verifyToken(token string) bool {
	if token != generateToken(getTokenContent(token)) {
		//TODO log invalid token
		return false
	}
	return true
}
