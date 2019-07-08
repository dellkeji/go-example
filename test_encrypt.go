package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	mrand "math/rand"
	"time"
)

// baseCryptObject := BaseCrypt{RootKey: []byte("8NT0Di5Zu1A*Czir")}

type BaseCrypt struct {
	RootKey []byte
}

func GenerateRandomRootKey() []byte {
	mrand.Seed(time.Now().UnixNano())
	const supportChars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*()_+-=?,.<>"
	rootKey := make([]byte, aes.BlockSize)
	for i := range rootKey {
		rootKey[i] = supportChars[mrand.Int63()%int64(len(supportChars))]
	}
	fmt.Printf("rootKey = %s\n", string(rootKey))
	return rootKey
}

// 使用CFB模式的AES做基本加密
func (baseCrypt *BaseCrypt) aesEncrypt(key, plaintext []byte) ([]byte, error) {
	aesBlockEncrypter, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return []byte{}, err
	}
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

// 使用CFB模式的AES做基本解密
func (baseCrypt *BaseCrypt) aesDecrypt(key, ciphertext []byte) ([]byte, error) {
	aesBlockDecrypter, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
	plaintext := make([]byte, len(ciphertext))
	aesDecrypter.XORKeyStream(plaintext, ciphertext)
	return plaintext, nil
}

// 生成随机密钥，并用base64编码
func (baseCrypt *BaseCrypt) GenerateInstanceKey() ([]byte, string, error) {
	// generate InstanceKey
	instanceKey := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, instanceKey); err != nil {
		return []byte{}, "", err
	}
	// AES[RootKey](InstanceKey)
	aesEncryptInstanceKey, err := baseCrypt.aesEncrypt(baseCrypt.RootKey, instanceKey)
	if err != nil {
		return []byte{}, "", err
	}
	return instanceKey, base64.StdEncoding.EncodeToString(aesEncryptInstanceKey), nil
}

// 加密 return: BASE64(AESRootKey(InstanceKey))， BASE64(AESInstanceKey(Password))
func (baseCrypt *BaseCrypt) Encrypt(plaintext string) (string, string, error) {
	// 实例密钥，以及其被加密并base64的值
	instanceKey, encryptInstanceKey, err := baseCrypt.GenerateInstanceKey()
	if err != nil {
		return "", "", err
	}
	// 使用实例密钥对所需内容加密
	secretTxt, err := baseCrypt.aesEncrypt(instanceKey, []byte(plaintext))
	if err != nil {
		return "", "", err
	}
	base64Txt := base64.StdEncoding.EncodeToString(secretTxt)
	fmt.Printf("encryptInstanceKey = %s\n", encryptInstanceKey)
	fmt.Printf("encryptPassword = %s\n", base64Txt)
	return encryptInstanceKey, base64Txt, nil
}

// 解密
func (baseCrypt *BaseCrypt) Decrypt(encryptInstanceKey string, ciphertext string) (string, error) {
	// 先使用根密钥解密实例密钥
	decodeEncryptInstanceKey, err := base64.StdEncoding.DecodeString(encryptInstanceKey)
	if err != nil {
		return "", err
	}
	instanceKey, err := baseCrypt.aesDecrypt(baseCrypt.RootKey, decodeEncryptInstanceKey)
	if err != nil {
		return "", err
	}
	decodeCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	plaintext, err := baseCrypt.aesDecrypt(instanceKey, decodeCiphertext)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func main() {
	// rootKey := GenerateRandomRootKey()
	// fmt.Println(string(rootKey))
	// baseCrypt := BaseCrypt{RootKey: rootKey}
	baseCrypt := BaseCrypt{RootKey: []byte(">oZtVqFclz3lR.sW")}
	// encryptInstanceKey, ciphertext, _ := baseCrypt.Encrypt("bEp1a2CcB9rNfHns")
	// fmt.Println(encryptInstanceKey, ciphertext)
	plaintext, _ := baseCrypt.Decrypt("bG2jnhMO86PE230WZPqy0FNkbY4B5hz+umcfgkJ8YiU=", "nzrG82Ewu31jIMn4T3dEuB8R3Ae/fTxwPC6jKfABT28=")
	fmt.Printf("plaintext = %s\n", plaintext)
}
