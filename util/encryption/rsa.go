package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

// 加密 加密完后 因为是乱码的  所以要base64之后再传递
func RsaEncrypt(origData []byte, publicKeyName string) ([]byte, error) {
	b, err := ioutil.ReadFile(publicKeyName)
	if err != nil {
		log.Fatalf("yamlFile.Get err %v ", err)
	}
	fmt.Println(string(b))
	//解密pem格式的公钥
	block, _ := pem.Decode(b)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	cr, err := rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
	if err != nil {
		return nil, err
	}
	//base64
	return []byte(base64.StdEncoding.EncodeToString(cr)), nil
}

// 解密 接收的是base64编码过后的密码  所以要先base64解码
func RsaDecrypt(base64Ciphertext []byte, privateKeyName string) ([]byte, error) {
	c, err := base64.StdEncoding.DecodeString(string(base64Ciphertext))
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadFile(privateKeyName)
	if err != nil {
		log.Fatalf(err.Error())
	}
	//解密
	block, _ := pem.Decode([]byte(b))
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, c)
}
