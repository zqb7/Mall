package util

import "github.com/dgrijalva/jwt-go"

type Claim struct {
	jwt.StandardClaims
	UserID   uint
	UserName string
	Password string
}

func GenerateToken(claim Claim) (token string, err error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return t.SignedString([]byte("bzyy"))
}

func ParseToken(token string) (Claim, error) {
	t, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("bzyy"), nil
	})
	if err != nil {
		return Claim{}, err
	}
	if claim, ok := t.Claims.(*Claim); ok && t.Valid {
		return *claim, nil
	} else {
		return Claim{}, err
	}

}
