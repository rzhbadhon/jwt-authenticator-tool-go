package models

import "encoding/json"

type VerifyRequest struct {
	Token  string `json:"token"`
	Secret string `json:"secret"`
	Alg    string `json:"alg"`
}

type VerifyResponse struct {
	ValidSignature bool            `json:"valid_signature"`
	Header         json.RawMessage `json:"header,omitempty"`
	Payload        json.RawMessage `json:"payload,omitempty"`
	Errors         []string        `json:"errors,omitempty"`
}
