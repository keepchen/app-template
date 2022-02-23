package jwt

import (
	"errors"
	"fmt"

	jwtLib "github.com/golang-jwt/jwt"
)

//AppClaims App票据声明
type AppClaims struct {
	Scopes   []string
	ScopeIDs []int64
	Name     string
	Email    string
	//more...
	jwtLib.StandardClaims
}

const (
	defaultTokenIssuer = "authority"
)

type SigningMethod string

const (
	SigningMethodRS256 SigningMethod = "RS256"
	SigningMethodRS512 SigningMethod = "RS512"
	SigningMethodHS512 SigningMethod = "HS512"
)

func (sm SigningMethod) getSigningMethod() jwtLib.SigningMethod {
	return jwtLib.GetSigningMethod(string(sm))
}

//Sign 签名
func Sign(claim jwtLib.Claims, conf Conf) (string, error) {
	appClaims, ok := claim.(AppClaims)
	if !ok {
		return "", errors.New("not supported claims")
	}

	switch conf.Algorithm {
	case string(SigningMethodRS256):
		return appClaims.GetToken(SigningMethodRS256, conf.privateKey)
	case string(SigningMethodRS512):
		return appClaims.GetToken(SigningMethodRS512, conf.privateKey)
	case string(SigningMethodHS512):
		return appClaims.GetToken(SigningMethodHS512, []byte(conf.HmacSecret))
	default:
		return "", errors.New("jwt secret not config")
	}
}

//Verify 验证
func Verify(tokenString string, conf Conf) (AppClaims, error) {
	token, err := jwtLib.ParseWithClaims(tokenString, &AppClaims{}, func(token *jwtLib.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtLib.SigningMethodHMAC); ok {
			return []byte(conf.HmacSecret), nil
		}

		if _, ok := token.Method.(*jwtLib.SigningMethodRSA); ok {
			//私钥加密，公钥解密，因此这里返回公钥
			return conf.privateKey.Public(), nil
		}

		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	})

	if err != nil {
		return AppClaims{}, err
	}

	if claims, ok := token.Claims.(*AppClaims); ok && token.Valid {
		return *claims, nil
	}

	return AppClaims{}, errors.New("token verify failed")
}

//GetToken 获取token
func (ac *AppClaims) GetToken(method SigningMethod, secret interface{}) (string, error) {
	if secret == nil {
		return "", errors.New("unsupported secret")
	}
	sm := method.getSigningMethod()
	token := jwtLib.NewWithClaims(sm, *ac)
	tokenStr, err := token.SignedString(secret)

	return tokenStr, err
}
