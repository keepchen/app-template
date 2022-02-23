package jwt

import (
	"crypto/rsa"
	"log"

	jwtLib "github.com/golang-jwt/jwt"

	"github.com/keepchen/app-template/pkg/utils"
)

//Conf 配置信息
type Conf struct {
	PublicKey   string `yaml:"public_key" toml:"public_key" json:"public_key"`       //公钥字符串或公钥文件地址
	PrivateKey  string `yaml:"private_key" toml:"private_key" json:"private_key"`    //私钥字符串或私钥文件地址
	Algorithm   string `yaml:"algorithm" toml:"algorithm" json:"algorithm"`          //加密算法: RS256 | RS512 | HS512
	HmacSecret  string `yaml:"hmac_secret" toml:"hmac_secret" json:"hmac_secret"`    //密钥
	TokenIssuer string `yaml:"token_issuer" toml:"token_issuer" json:"token_issuer"` //令牌颁发者
	privateKey  *rsa.PrivateKey
	publicKey   *rsa.PublicKey
}

//Load 载入配置
func (c *Conf) Load() {
	if utils.FileExists(c.PublicKey) {
		contents, err := utils.FileGetContents(c.PublicKey)
		if err != nil {
			c.PublicKey = string(contents)
		}
	}

	if utils.FileExists(c.PrivateKey) {
		log.Println("xxxxxxxxxx")
		contents, err := utils.FileGetContents(c.PrivateKey)
		if err != nil {
			log.Println("yyyyyyyyyyyy")
			c.PrivateKey = string(contents)
		}
	}

	if len(c.PrivateKey) == 0 || len(c.PublicKey) == 0 {
		return
	}

	pri, err := jwtLib.ParseRSAPrivateKeyFromPEM([]byte(c.PrivateKey))
	if err != nil {
		panic(err)
	}
	pub, err := jwtLib.ParseRSAPublicKeyFromPEM([]byte(c.PublicKey))
	if err != nil {
		panic(err)
	}
	c.privateKey = pri
	c.publicKey = pub
}
