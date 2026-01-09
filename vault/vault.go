package vault

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"golang.org/x/crypto/argon2"
	"golang.org/x/term"
)

type VaultFile struct {
	Version int       `json:"version"`
	KDF     KDFConfig `json:"kdf"`
	Cipher  Cipher    `json:"cipher"`
}

type KDFConfig struct {
	Algorithm   string `json:"algorithm"`
	Salt        string `json:"salt"`
	Iterations  uint32 `json:"iterations"`
	Memory      uint32 `json:"memory"`
	Parallelism uint8  `json:"parallelism"`
}

type Cipher struct {
	Algorithm string `json:"algorithm"`
	Nonce     string `json:"nonce"`
	Data      string `json:"data"`
}

func InitVault(path string) error {
	fmt.Print("Choose a master password: \n")
	masterPwd, err := term.ReadPassword(int(os.Stdin.Fd()))

	salt, err := randomBytes(16)
	if err != nil {
		return err
	}

	key := argon2.IDKey(masterPwd, salt, 3, 64*1024, 2, 32)

	zero(masterPwd)

	plaintext := []byte(`{"entries":[]}`)

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonce, err := randomBytes(gcm.NonceSize())
	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)

	zero(plaintext)
	zero(key)

	vault := VaultFile{
		Version: 1,
		KDF: KDFConfig{
			Algorithm:   "argon2id",
			Salt:        base64.StdEncoding.EncodeToString(salt),
			Iterations:  3,
			Memory:      64 * 1024,
			Parallelism: 2,
		},
		Cipher: Cipher{
			Algorithm: "aes-256-gcm",
			Nonce:     base64.StdEncoding.EncodeToString(nonce),
			Data:      base64.StdEncoding.EncodeToString(ciphertext),
		},
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer file.Close()
	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	return enc.Encode(vault)

}

func readVault(path string) (*VaultFile, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var vault VaultFile
	if err = json.NewDecoder(f).Decode(&vault); err != nil {
		return nil, err
	}

	return &vault, nil

}

func unlockVault(path string) ([]byte, *VaultFile, error) {
	vault, err := readVault(path)
	if err != nil {
		return nil, nil, err
	}

	fmt.Print("Enter your master password: \n")
	masterPwd, err := term.ReadPassword(int(os.Stdin.Fd()))

	salt, err := base64.StdEncoding.DecodeString(vault.KDF.Salt)
	if err != nil {
		return nil, nil, err
	}

	key := argon2.IDKey(masterPwd, salt, 3, 64*1024, 2, 32)

	zero(masterPwd)

	nonce, err := base64.StdEncoding.DecodeString(vault.Cipher.Nonce)
	if err != nil {
		return nil, nil, err
	}

	cipherText, err := base64.StdEncoding.DecodeString(vault.Cipher.Data)
	if err != nil {
		return nil, nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}

	plaintext, err := gcm.Open(nil, nonce, cipherText, nil)
	zero(key)
	if err != nil {
		return nil, nil, errors.New("Invalid password or corrupted vault...")
	}

	return plaintext, vault, nil

}

func randomBytes(size int) ([]byte, error) {
	buf := make([]byte, size)
	_, err := rand.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func zero(buf []byte) {
	for i := range buf {
		buf[i] = 0
	}
}
