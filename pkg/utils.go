package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"os/exec"
	"strings"
)

func encrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte("examplekey123456"))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, []byte("examplekey123456"))
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func decrypt(encryptedText string) (string, error) {
	block, err := aes.NewCipher([]byte("examplekey123456"))
	if err != nil {
		return "", err
	}
	cipherText, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}
	cfb := cipher.NewCFBDecrypter(block, []byte("examplekey123456"))
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

func CopyToClipboard(text string) error {
	cmd := exec.Command("xclip", "-selection", "clipboard")
	cmd.Stdin = strings.NewReader(text)
	return cmd.Run()
}
