package handlers

import (
	"encoding/json"
	"jwt-tool/models"
	"jwt-tool/util"
	"net/http"
	"strings"
)

func CreateJWT(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")


	var req models.VerifyRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil{
		http.Error(w, `{"errors":["invalid json"]}`, http.StatusBadRequest)
		return
	}

	if req.Token == "" || req.Secret == "" || req.Alg == ""{
		http.Error(w, `{"errors":["token, secret, alg required"]}`, http.StatusBadRequest)
        return
	}

	if strings.ToUpper(req.Alg) != "HS256"{
		 http.Error(w, `{"errors":["only HS256 supported"]}`, http.StatusBadRequest)
        return
	}

	headerJSON, payloadJSON, headerPart, payloadPart, sigPart, err := util.DecodeJWTParts(req.Token)

	if err != nil{
		 response := models.VerifyResponse{Errors: []string{err.Error()}}
		 encoder := json.NewEncoder(w)
		 encoder.Encode(response)
		 return
	}

	isValid, err := util.VerifyHS256(headerPart, payloadPart, sigPart, []byte(req.Secret))

	response := models.VerifyResponse{
		ValidSignature: isValid,
		Header: headerJSON,
		Payload: payloadJSON,
	}
	if err != nil{
		response.Errors = []string{err.Error()}
	}

encoder := json.NewEncoder(w)
encoder.Encode(response)

}

