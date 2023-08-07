package jwt

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Jonda-HR/Goland_twitter/v2/models"
	jwt "github.com/golang-jwt/jwt/v5"
)

var Email string
var IDUsuario string

func ProcessToken(tk string, JWTSing string) (*models.Claim, bool, string, error) {
	myKey := []byte(JWTSing)
	var claims models.Claim

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("formato de token Invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err == nil {
		fmt.Println("Hola mundo")
	}

	if !tkn.Valid {
		return &claims, false, string(""), errors.New("invalid Token")
	}

	return &claims, false, string(""), err

}
