package security

import (
	"database/sql"
	"time"
)

type JWT struct {
	Header    JWTHeader  `json:"header"`
	Payload   JWTPayload `json:"payload"`
	Signature string     `json:"signature"`
}

type JWTHeader struct {
	Alg string
	Typ string
}

type JWTPayload struct {
	Sub string `json:"sub"`
	Iat int64  `json:"iat"`
	Exp int64  `json:"exp"`
}

func CreateJWTTokenFromJSON(inputJSON string) (string, error) {

}

func CreateJWTToken() {
	Header := JWTHeader{"RS256", "JWT"}
	Payload := CreateJWTPayload()

}

// Index из бд, время выдачи токена, время истечения токена
func CreateJWTPayload() JWTPayload {
	var db sql.DB

	connStr := ""



	Payload := JWTPayload{
		Sub: ,
		Iat: time.Now().Unix(),
		Exp: time.Now().Add(time.Hour * 12).Unix(),
	}

	return Payload
}
