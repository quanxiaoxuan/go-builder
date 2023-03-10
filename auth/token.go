package auth

import (
	"encoding/json"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// 秘钥
const SecretKey = "1bM1qeHqi33rdiXr9SbrCCmkOeUXTw6vJpzzjhzKmG5g5bhILjOy48pWCpZMCMHU76qZ203KW2tEFNIxTLPxD16AEqEdKWcRznl6"

// 用户token参数
type Param struct {
	UserId     string `json:"userId"`      // 用户ID
	UserName   string `json:"userName"`    // 用户ID
	Phone      string `json:"phone"`       // 登录手机
	LoginIp    string `json:"loginIp"`     // 登录IP
	ExpireTime int64  `json:"sessionTime"` // 到期时间
}

// 生成token
func BuildAuthToken(param *Param) (string, error) {
	bytes, err := json.Marshal(param)
	var mapClaims jwt.MapClaims
	err = json.Unmarshal(bytes, &mapClaims)
	if err != nil {
		return "", err
	}
	return GenerateToken(mapClaims, SecretKey)
}

// 解析token
func ParseAuthToken(token string) (*Param, error) {
	var tp Param
	resultMap, err := ParseTokenToJson(token, SecretKey)
	if err != nil {
		log.Error("解析token失败", err)
		return nil, err
	}
	tp.UserId = resultMap["userId"].(string)
	tp.UserName = resultMap["userName"].(string)
	tp.Phone = resultMap["phone"].(string)
	tp.LoginIp = resultMap["loginIp"].(string)
	tp.ExpireTime = int64(resultMap["sessionTime"].(float64))
	return &tp, nil
}

// 获取用户ID
func GetUserId(context *gin.Context) (userId int64) {
	auth := context.Request.Header.Get("Authorization")
	tp, err := ParseAuthToken(auth)
	if err != nil {
		return
	}
	userId, err = strconv.ParseInt(tp.UserId, 10, 64)
	if err != nil {
		return
	}
	return
}

// 生成Token值
func GenerateToken(mapClaims jwt.MapClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	return token.SignedString([]byte(secret))
}

// 解析token
func ParseTokenToJson(token string, secret string) (map[string]interface{}, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	var result = make(map[string]interface{})
	result = claim.Claims.(jwt.MapClaims)
	return result, nil
}
