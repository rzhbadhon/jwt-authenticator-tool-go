package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func DecodeJWTParts(token string) (
	headerJSON,
	payloadJSON json.RawMessage,
	headerPart,
	payloadPart,
	sigPart string,
	err error) {

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, nil, "", "", "", errors.New("invalid token format")
	}

	rawHeaderJSON, err := B64urlDecode(parts[0])
	if err != nil {
		return nil, nil, "", "", "", fmt.Errorf("header decode error %w", err)
	}

	rawPayloadJSON, err := B64urlDecode(parts[1])
	if err != nil {
		return nil, nil, "", "", "", fmt.Errorf("payload decode error %w", err)
	}
	var temp any
	if err := json.Unmarshal(rawHeaderJSON, &temp); err != nil {
		return nil, nil, "", "", "", fmt.Errorf("header json error: %w", err)
	}

	if err := json.Unmarshal(rawPayloadJSON, &temp); err != nil {
		return nil, nil, "", "", "", fmt.Errorf("payload json error: %w", err)
	}

	return json.RawMessage(rawHeaderJSON), json.RawMessage(rawPayloadJSON), parts[0], parts[1], parts[2], nil

}
