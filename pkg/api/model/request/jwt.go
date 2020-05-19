package request

import (
	"github.com/dgrijalva/jwt-go"
)

// Custom claims structure
type CustomClaims struct {
	ID          uint
	NickName    string
	AuthorityId string
	jwt.StandardClaims
}
