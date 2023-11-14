package authmanager

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func IsExpired(claims jwt.MapClaims) bool {
	if _, ok := claims["exp"]; !ok {
		return false
	}

	var tm time.Time
	switch exp := claims["exp"].(type) {
	case float64:
		tm = time.Unix(int64(exp), 0)
	case json.Number:
		v, _ := exp.Int64()
		tm = time.Unix(v, 0)
	}

	return time.Now().Add(10 * time.Second).After(tm)
}

func GetTokenClaims(accessToken string) (jwt.MapClaims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(accessToken, jwt.MapClaims{})
	if err != nil {
		return nil, fmt.Errorf("failed to parse accessToken: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed to convert token's claims to standard claims: %w", err)
	}
	return claims, nil
}
