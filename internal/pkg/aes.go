package pkg

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

// AES加密密钥必须是16、24或32字节
const AESKey = "01234567890etcdm"

// Encrypt AES加密
func Encrypt(plaintext string) (string, error) {
	// 创建加密块
	block, err := aes.NewCipher([]byte(AESKey))
	if err != nil {
		return "", err
	}

	// 填充明文
	padding := block.BlockSize() - len(plaintext)%block.BlockSize()
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	rawData := append([]byte(plaintext), padText...)

	// 加密
	ciphertext := make([]byte, len(rawData))
	mode := cipher.NewCBCEncrypter(block, []byte(AESKey)[:block.BlockSize()])
	mode.CryptBlocks(ciphertext, rawData)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt AES解密
func Decrypt(ciphertext string) (string, error) {
	// base64解码
	encrypted, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	// 判断加密文本是否为空
	if len(encrypted) == 0 {
		return "", errors.New("invalid ciphertext")
	}

	// 创建解密块
	block, err := aes.NewCipher([]byte(AESKey))
	if err != nil {
		return "", err
	}

	// 检查加密文本长度是否为块大小的整数倍
	if len(encrypted)%block.BlockSize() != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}

	// 解密
	mode := cipher.NewCBCDecrypter(block, []byte(AESKey)[:block.BlockSize()])
	mode.CryptBlocks(encrypted, encrypted)

	// 去除填充
	padding := encrypted[len(encrypted)-1]
	encrypted = encrypted[:len(encrypted)-int(padding)]

	return string(encrypted), nil
}
