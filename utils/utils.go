package utils

import (
	"bytes"
	"crypto/aes"
	"dailylife/conf"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// 加密函数 开始
// Pad applies PKCS7 padding to the input data.
func Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// Encrypt encrypts plaintext using AES-ECB mode and returns a base64 encoded string.
func Encrypt(t, key string) (string, error) {
	// Convert key to bytes
	keyBytes := []byte(key)
	// Convert plaintext to bytes and then pad it
	plaintextBytes := Pad([]byte(t), aes.BlockSize)

	// Create a new AES cipher with the provided key
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", fmt.Errorf("failed to create AES cipher: %v", err)
	}

	// Encrypt the padded plaintext
	ciphertext := make([]byte, len(plaintextBytes))
	for i := 0; i < len(plaintextBytes); i += aes.BlockSize {
		block.Encrypt(ciphertext[i:i+aes.BlockSize], plaintextBytes[i:i+aes.BlockSize])
	}

	// Encode the ciphertext to Base64 and return it
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// 加密函数结束

// 获取学校ID
func GetSchoolID(schoolName string) (string, error) {
	req, err := http.NewRequest("GET", conf.SchoolList_url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}
	// 设置请求头
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 16_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.6 Mobile/15E148 Safari/604.1 Edg/119.0.0.0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make GET request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	// 解析resp
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}
	var result struct {
		Data []conf.School `json:"data"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %v", err)
	}
	for _, school := range result.Data {
		if school.Name == schoolName {
			return school.ID, nil
		}
	}

	return "", nil
}
