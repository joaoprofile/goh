package security

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	env "github.com/joaocprofile/goh/environment"
)

func TokenGenerate(userId, tenantId string) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 24).Unix()
	permissions["user_id"] = userId
	permissions["tenant_id"] = tenantId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(env.Get().Security.JWTSecret))
}

func TokenValidate(r *http.Request) error {
	tokenHeader := r.Header.Get("Authorization")
	tokenArray := strings.Split(tokenHeader, " ")
	var tokenString string
	if len(tokenArray) == 2 {
		tokenString = tokenArray[1]
	}
	if tokenString == "" {
		return errors.New("Token missing from request")
	}
	token, err := jwt.Parse(tokenString, getVerifyKey)
	if err != nil {
		return errors.New("Invalid: " + err.Error())
	}
	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := permissions["user_id"]
		tenantId := permissions["tenant_id"]
		if userId == nil {
			return errors.New("Invalid user authorization " + err.Error())
		}
		if tenantId == nil {
			return errors.New("Invalid Tenant authorization " + err.Error())
		}
		Session().User(fmt.Sprint(userId))
		Session().Tenant(fmt.Sprint(tenantId))
		return nil
	}
	return errors.New("Invalid token")
}

func getVerifyKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil,
			fmt.Errorf("Invalid subscription method, %v", token.Header["alg"])
	}
	return env.Get().Security.JWTSecret, nil
}
