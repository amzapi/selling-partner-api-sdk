package decryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func Decrypt(encryptionKey string, initVector string, dataToDecrypt []byte) ([]byte, error) {
	iv, err := base64.StdEncoding.DecodeString(initVector)
	if err != nil {
		return nil, err
	}

	key, err := base64.StdEncoding.DecodeString(encryptionKey)
	if err != nil {
		return nil, err
	}

	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	decryptedData := make([]byte, len(dataToDecrypt))
	cipher.NewCBCDecrypter(cipherBlock, iv).CryptBlocks(decryptedData, dataToDecrypt)
	return decryptedData, nil
}
