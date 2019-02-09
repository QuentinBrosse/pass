package aescbc

import (
	"bytes"
	"crypto/aes"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	key                   = []byte("0123456789ABCDEF")
	expectedDecryptedData = []byte("Hello, World!")
	expectedEncryptedData = []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xc5, 0x53, 0x7d, 0xbf, 0x47, 0x9f, 0x80, 0xa3, 0x73, 0x2f, 0x9a, 0x77, 0x54, 0xce, 0x88, 0x8e}
)

func TestEncrypt(t *testing.T) {
	iv := bytes.Repeat([]byte{0}, aes.BlockSize)
	encryptedData, _ := Encrypt(expectedDecryptedData, key, bytes.NewBuffer(iv))
	assert.Equal(t, expectedEncryptedData, encryptedData)
}

func TestDecrypt(t *testing.T) {
	decryptedData, _ := Decrypt(expectedEncryptedData, key)
	assert.Equal(t, expectedDecryptedData, decryptedData)
}
