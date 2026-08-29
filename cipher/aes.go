// Package cipher provides symmetric encryption helpers.
package cipher

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"os"
)

var (
	// DefaultAesKey is the key used by the Aes*Default helpers. It can be
	// overridden at startup with the ctj.cipher.aes.key environment variable,
	// encoded as standard base64.
	DefaultAesKey, _ = base64.StdEncoding.DecodeString("PZLQ44ST8O2wquFo+GsD1DnHW/z/UA7Sy+6bl70MBaY=")
)

func init() {
	if encodedKey := os.Getenv("ctj.cipher.aes.key"); encodedKey != "" {
		if key, err := base64.StdEncoding.DecodeString(encodedKey); err == nil {
			DefaultAesKey = key
		}
	}
}

// AesEnc encrypts data with AES-CBC and PKCS#7 padding.
//
// For compatibility with the source utility, the IV is derived from the first
// AES block of key. Callers that need semantic security should use an
// authenticated encryption mode with a random nonce instead.
func AesEnc(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	padding := blockSize - len(origData)%blockSize
	origData = append(origData, bytes.Repeat([]byte{byte(padding)}, padding)...)

	cipherText := make([]byte, len(origData))
	cipher.NewCBCEncrypter(block, key[:blockSize]).CryptBlocks(cipherText, origData)
	return cipherText, nil
}

// AesEncDefault encrypts data with DefaultAesKey.
func AesEncDefault(origData []byte) ([]byte, error) {
	return AesEnc(origData, DefaultAesKey)
}

// AesEncDefaultStr encrypts a string with DefaultAesKey and encodes it using
// standard base64.
func AesEncDefaultStr(origData string) (string, error) {
	data, err := AesEncDefault([]byte(origData))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

// AesDec decrypts AES-CBC data encrypted by AesEnc and removes PKCS#7 padding.
func AesDec(cipherData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	if len(cipherData) == 0 || len(cipherData)%blockSize != 0 {
		return nil, errors.New("ciphertext must be a non-empty multiple of the AES block size")
	}

	origData := make([]byte, len(cipherData))
	cipher.NewCBCDecrypter(block, key[:blockSize]).CryptBlocks(origData, cipherData)

	padding := int(origData[len(origData)-1])
	if padding == 0 || padding > blockSize || padding > len(origData) {
		return nil, errors.New("invalid PKCS#7 padding")
	}
	for _, value := range origData[len(origData)-padding:] {
		if int(value) != padding {
			return nil, errors.New("invalid PKCS#7 padding")
		}
	}
	return origData[:len(origData)-padding], nil
}

// AesDecDefault decrypts data with DefaultAesKey.
func AesDecDefault(cipherData []byte) ([]byte, error) {
	return AesDec(cipherData, DefaultAesKey)
}

// AesDecDefaultStr decodes standard base64 data and decrypts it with DefaultAesKey.
func AesDecDefaultStr(cipherData string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(cipherData)
	if err != nil {
		return "", err
	}
	origData, err := AesDecDefault(data)
	if err != nil {
		return "", err
	}
	return string(origData), nil
}
