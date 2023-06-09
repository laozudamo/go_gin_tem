package utils

import (
	response "goGinTem/Response"
	"goGinTem/middlewares"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"time"
)

func CreateToken(c *gin.Context, Id int, Tel string) string {
	//生成token信息
	j := middlewares.NewJWT()
	claims := middlewares.CustomClaims{
		ID:  uint(Id),
		Tel: Tel,
		// NickName: NickName,
		// AuthorityId: uint(Role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			// TODO 设置token过期时间
			ExpiresAt: time.Now().Unix() + 60*60*24*30, //token -->30天过期
			Issuer:    "test",
		},
	}
	//生成token
	token, err := j.CreateToken(claims)
	if err != nil {
		response.Success(c, 401, "token生成失败,重新再试", "test")
		return ""
	}
	return token
}
