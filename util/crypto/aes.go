package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"

	"github.com/aokoli/goutils"
)

const nonce_length = 16

func AESGCMEncrypt(origin, seed string) (ciphertext string, err error) {
	rr, err := goutils.CryptoRandom(16, 0, 127, false, false)
	if err != nil {
		return
	}
	randNonce := hex.EncodeToString([]byte(rr))
	ciphertext, err = aesgcmEncryptWithNonce(origin, seed, randNonce)
	if err != nil {
		return
	}
	return ciphertext + randNonce, err
}

func aesgcmEncryptWithNonce(origin, seed, nonce string) (ciphertext string, err error) {
	seedByte, err := hex.DecodeString(seed)
	if err != nil {
		return
	}
	block, err := aes.NewCipher(seedByte)
	if err != nil {
		return
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}
	nonceByte, err := hex.DecodeString(nonce)
	if err != nil {
		return
	}
	ciphertext = hex.EncodeToString(aesgcm.Seal(nil, nonceByte, []byte(origin), nil))
	return
}

func AESGCMDecrypt(ciphertextStr, seed string) (ciphertext string, err error) {
	nonceStart := len(ciphertextStr) - nonce_length*2
	nonce := ciphertextStr[nonceStart:]
	return aesgcmDecryptWithNonce(ciphertextStr[:nonceStart], seed, nonce)
}

func aesgcmDecryptWithNonce(ciphertextStr, seed, nonce string) (origin string, err error) {
	ciphertext, err := hex.DecodeString(ciphertextStr)
	if err != nil {
		return
	}
	seedByte, err := hex.DecodeString(seed)
	if err != nil {
		return
	}
	block, err := aes.NewCipher(seedByte)
	if err != nil {
		return
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}
	nonceByte, err := hex.DecodeString(nonce)
	if err != nil {
		return
	}
	originByte, err := aesgcm.Open(nil, nonceByte, ciphertext, nil)
	if err != nil {
		return
	}
	origin = string(originByte)
	return
}
