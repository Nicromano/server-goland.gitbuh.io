package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

//Encripta datos con una clave establecida
func Encrypt(text string) []byte {

	key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	/* fmt.Printf("%x\n", ciphertext) */
	return ciphertext
}

//Desencriptar datos
func Decrypt(hash []byte) []byte {
	key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")

	/* ciphertext, _ := hex.DecodeString(hex.EncodeToString(hash)) */
	nonce, _ := hex.DecodeString("64a9433eae7ccceee2fc0eda")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, hash, nil)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%s\n", plaintext)
	return plaintext
}
