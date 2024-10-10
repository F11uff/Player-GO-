package security

import "github.com/dgrijalva/jwt-go"

func CreateJWTToken(password, username string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	headerClaims := token.Claims.(jwt.MapClaims)

	headerClaims["username"] = username
	headerClaims["password"] = password

	secretKey := []byte("secretKey")

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return ""
	}

	return tokenString
}
