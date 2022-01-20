package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

//RSAEncode rsa加密
func RSAEncode(rawString string, publicKey []byte) (string, error) {
	block, _ := pem.Decode(publicKey)

	if block == nil {
		return "", errors.New("public key invalid")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	pub := pubInterface.(*rsa.PublicKey)

	encodedByte, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(rawString))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(encodedByte), nil
}

//RSADecode rsa解密
func RSADecode(encodedString string, privateKey []byte) (string, error) {
	encodedByte, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		return "", errors.New("invalid base64 encode string")
	}

	block, _ := pem.Decode(privateKey)
	if block == nil {
		return "", errors.New("private key invalid")
	}

	pri, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	priObj, ok := pri.(*rsa.PrivateKey)
	if !ok {
		return "", errors.New("private key invalid")
	}

	decodedByte, err := rsa.DecryptPKCS1v15(rand.Reader, priObj, encodedByte)
	if err != nil {
		return "", err
	}

	return string(decodedByte), nil
}
