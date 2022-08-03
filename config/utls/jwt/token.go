package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// JWTsecret token密钥
var JWTsecret = []byte("Krian")

// Claims token信息结构体
type Claims struct {
	Id       uint   `json:"id"`
	UserName string `json:"user_name"`
	//Password string `json:"password"` // token中存放密码不安全！！！
	jwt.StandardClaims
}

// GenerateToken
// @Func: 签发token
func GenerateToken(id uint, username string) (string, error) {
	notTime := time.Now()
	// 设置token过期时间：
	expireTime := notTime.Add(24 * time.Hour)
	// 设置token信息：
	claims := Claims{
		Id:       id,
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "awesomeProject",
		},
	}
	// 设置token加密：
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := tokenClaims.SignedString(JWTsecret)
	return token, err
}

// ParseToken
// @Func：验证token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTsecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, err
		}
	}
	return nil, err
}
