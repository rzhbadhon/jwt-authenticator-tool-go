package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func VerifyHS256(headerPart, payloadPart, sigPart string, secret []byte) (bool, error) {

	signatureInput := headerPart + "." + payloadPart

	h := hmac.New(sha256.New, secret)
	h.Write([]byte(signatureInput))

	expectedSignature := h.Sum(nil)

	givenSignatureByte, err := B64urlDecode(sigPart) // String to Byte Array

	if err != nil {
		return false, fmt.Errorf("signature decoding error: %w", err)
	}

	return hmac.Equal(givenSignatureByte, expectedSignature), nil // sending boolean true if it matches and false if !
}
