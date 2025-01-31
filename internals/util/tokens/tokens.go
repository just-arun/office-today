package tokens

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/just-arun/office-today/internals/boot/config"
	"github.com/just-arun/office-today/internals/util/stringutil"
)

// JWTTokenType json web token type
type JWTTokenType int

const (
	// AccessToken fot token
	AccessToken JWTTokenType = iota
	// RefreshToken fot token
	RefreshToken
)

// GenerateToken create jwt token
func GenerateToken(userID string, tokenType JWTTokenType) (string, error) {
	// create the token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := make(jwt.MapClaims)

	// set some claims
	claims["id"] = stringutil.HashFromString(userID)
	if tokenType == AccessToken {
		claims["exp"] = time.Now().Add(time.Minute * config.JWTAccessTokenTime).Unix()
	} else {
		claims["exp"] = time.Now().Add(time.Minute * config.JWTRefreshTokenTime).Unix()
	}

	fmt.Printf("[claims %v]: %v \n", tokenType, claims["exp"])

	fmt.Printf("Access: %v, Refresh: %v \n", time.Minute*config.JWTAccessTokenTime, time.Minute*config.JWTRefreshTokenTime)

	token.Claims = claims

	//Sign and get the complete encoded token as string
	return (token.SignedString([]byte(config.TokenSignature)))
}

// DecodeJWTToken unwrap token
func DecodeJWTToken(tokenString string) (token interface{}, clain map[string]interface{}, err error) {
	claims := jwt.MapClaims{}
	tok, error := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.TokenSignature), nil
	})

	if err != nil || !tok.Valid {
		return nil, nil, errors.New("Token is not valid")
	}

	return token, claims, error
}

// GetTokenRemainingValidity token validate
func GetTokenRemainingValidity(timestamp interface{}) int {
	var expireOffset = 0
	if validity, ok := timestamp.(int64); ok {
		tm := time.Unix(int64(validity), 0)
		remainder := tm.Sub(time.Now())

		if remainder > 0 {
			return int(remainder.Minutes())
		}
	} else {
		fmt.Println("not ok")
	}
	return expireOffset
}

// GetTokenFromHeader get token from header
func GetTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header["Authorization"]

	if authHeader == nil {
		return "", errors.New("Token not found")
	}

	authHeaderArr := strings.Split(authHeader[0], " ")
	if len(authHeaderArr) != 2 {
		return "", errors.New("Token not found")
	}
	if authHeaderArr[0] != "Bearer" {
		return "", errors.New("Token not valied")
	}

	return authHeaderArr[1], nil
}

// ValidateToken validates token
func ValidateToken(token string) (jwtClaim map[string]interface{}, valide bool) {
	claim := jwt.MapClaims{}
	tkn, err := jwt.ParseWithClaims(
		token,
		claim,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.TokenSignature), nil
		})
	if err != nil {
		fmt.Println("[ERROR]", err.Error())
		return nil, false
	}
	return claim, tkn.Valid
}
