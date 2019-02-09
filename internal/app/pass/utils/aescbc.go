package aescbc

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"io"
)

// Add bytes to data so len(data) is a multiple of aes.BlockSize
// Padding size can be anywhere between [1 ; aes.BlockSize]
func pad(data []byte) []byte {
	paddingSize := aes.BlockSize - len(data)%aes.BlockSize
	padding := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)

	return append(data, padding...)
}

// Remove the padding added by pad in data
func unpad(data []byte) []byte {
	dataLen := len(data)
	paddingSize := int(data[dataLen-1])

	return data[:dataLen-paddingSize]
}

// Encrypt data using AES-CBC with the specified key and an IV read from ivReader
// data is a byte array to encrypt
// key is a 16 or 32-byte array used to encrypt the data
// ivReader allows to pass either a fixed IV (e.g. bytes.NewBuffer(...)) or generating a random one (e.g. rand.Reader)
func Encrypt(data, key []byte, ivReader io.Reader) ([]byte, error) {
	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)

	if _, err := io.ReadFull(ivReader, iv); err != nil {
		return nil, err
	}

	cbcEncrypter := cipher.NewCBCEncrypter(aesCipher, iv)

	data = pad(data)

	cbcEncrypter.CryptBlocks(data, data)

	return append(iv, data...), nil
}

// Decrypt the AES-CBC encrypted data with the specified key
// data is a byte array to decrypt
// key is a 16 or 32-byte array used to decrypt the data
func Decrypt(data, key []byte) ([]byte, error) {
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]
	aesCipher, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	cbcDecrypter := cipher.NewCBCDecrypter(aesCipher, iv)

	cbcDecrypter.CryptBlocks(data, data)

	return unpad(data), nil
}
