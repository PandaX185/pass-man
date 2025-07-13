package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"os"
	"os/exec"
	"strings"
)

func generateIv(length int) ([]byte, error) {
	iv := make([]byte, length)
	_, err := rand.Read(iv)
	if err != nil {
		return nil, err
	}
	return iv, nil
}

func encrypt(text string) (string, error) {
	key := os.Getenv("ENCRYPTION_KEY")
	iv, err := generateIv(aes.BlockSize)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, iv)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	cipherText = append(iv, cipherText...)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func decrypt(encryptedText string) (string, error) {
	key := os.Getenv("ENCRYPTION_KEY")
	if key == "" {
		return "", nil
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	cipherText, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

func CopyToClipboard(text string) error {
	cmd := exec.Command("xclip", "-selection", "clipboard")
	cmd.Stdin = strings.NewReader(text)
	return cmd.Run()
}

func ConvertToJSON(plain string) (string, error) {
	data := strings.Split(plain, "\n")
	mp := make(map[string]string)
	for _, line := range data {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			mp[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}
	jsonData, err := json.Marshal(mp)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
