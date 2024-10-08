package security

import (
	_ "database/sql"
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
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

func CreateJWTTokenFromJSON(inputJSON string) (string, error) {

}

func CreateJWTToken() JWT {

	JWTToken := JWT{
		Header:  JWTHeader{"RS256", "JWT"},
		Payload: CreateJWTPayload(),
	}

	//Header := JWTHeader{"RS256", "JWT"}
	//Payload := CreateJWTPayload()

	return JWTToken
}

// Время выдачи токена, время истечения токена
func CreateJWTPayload() JWTPayload {
	//var _ sql.DB // Изменить на db

	Payload := JWTPayload{}

	return Payload
}
