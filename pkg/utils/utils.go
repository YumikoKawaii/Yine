package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

var (
	TimeFormat string = "2006-01-02 15:04:05"
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

func messageResponse(content string) []byte {

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
	res := messageResponse(message)
	w.Header().Set(header_type, header_value)
	w.WriteHeader(status)
	w.Write(res)
}

// This will be put in env
// Key using for cipher
var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// This will be put in env, too
// Key using for aes
var key = "YumikoSekaideIchibanKawaii22@123vjpPro"

// Key set to create unique ID
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func EncryptData(data string) (string, error) {

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	originData := []byte(data)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherData := make([]byte, len(originData))
	cfb.XORKeyStream(cipherData, originData)
	return encode(cipherData), nil
}

func decode(s string) []byte {

	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}

	return data
}

func DecryptData(encryptedData string) (string, error) {

	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return "", err
	}

	cipherData := decode(encryptedData)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	originData := make([]byte, len(cipherData))
	cfb.XORKeyStream(originData, cipherData)
	return string(originData), nil
}

func RandomStringRunes(n int) string {

	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)

}
