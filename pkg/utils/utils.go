package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {

	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		err := json.Unmarshal([]byte(body), x)
		if err != nil {
			return
		}
	}
}

func MessageResponse(content string) []byte {

	m := struct {
		Message string `json:"Message"`
	}{}

	m.Message = content

	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	return b

}

func Hashing(text string) string {

	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])

}

func ResponseWriter(w http.ResponseWriter, header_type string, header_value string, status int, message string) {
	res := MessageResponse(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(res)
}
