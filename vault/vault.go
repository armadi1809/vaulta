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

// KDF configuration constants
const (
	kdfIterations  uint32 = 3
	kdfMemory      uint32 = 64 * 1024
	kdfParallelism uint8  = 2
	kdfKeyLength   uint32 = 32
	saltSize       int    = 16
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

type Entry struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Notes    string `json:"notes"`
}

type VaultData struct {
	Entries []Entry `json:"entries"`
}

// deriveKey derives an encryption key from a password and salt using Argon2id
func deriveKey(password, salt []byte) []byte {
	return argon2.IDKey(password, salt, kdfIterations, kdfMemory, kdfParallelism, kdfKeyLength)
}

// encrypt encrypts plaintext using AES-256-GCM and returns nonce and ciphertext
func encrypt(key, plaintext []byte) (nonce, ciphertext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}

	nonce, err = randomBytes(gcm.NonceSize())
	if err != nil {
		return nil, nil, err
	}

	ciphertext = gcm.Seal(nil, nonce, plaintext, nil)
	return nonce, ciphertext, nil
}

// decrypt decrypts ciphertext using AES-256-GCM
func decrypt(key, nonce, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, errors.New("invalid password or corrupted vault")
	}

	return plaintext, nil
}

// readVaultFile reads and parses the vault JSON file
func readVaultFile(path string) (*VaultFile, error) {
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

// writeVaultFile writes the vault to a JSON file
func writeVaultFile(path string, vault *VaultFile) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	return enc.Encode(vault)
}

// promptPassword prompts the user for a password
func promptPassword(prompt string) ([]byte, error) {
	fmt.Print(prompt)
	return term.ReadPassword(int(os.Stdin.Fd()))
}

func promptNormal(prompt string) (string, error) {
	fmt.Print(prompt)
	var input string
	_, err := fmt.Scanln(&input)
	return input, err
}

// newVaultFile creates a new VaultFile with the given salt, nonce, and ciphertext
func newVaultFile(salt, nonce, ciphertext []byte) *VaultFile {
	return &VaultFile{
		Version: 1,
		KDF: KDFConfig{
			Algorithm:   "argon2id",
			Salt:        base64.StdEncoding.EncodeToString(salt),
			Iterations:  kdfIterations,
			Memory:      kdfMemory,
			Parallelism: kdfParallelism,
		},
		Cipher: Cipher{
			Algorithm: "aes-256-gcm",
			Nonce:     base64.StdEncoding.EncodeToString(nonce),
			Data:      base64.StdEncoding.EncodeToString(ciphertext),
		},
	}
}

// updateVaultCipher updates the cipher data in an existing vault
func (v *VaultFile) updateCipher(nonce, ciphertext []byte) {
	v.Cipher.Nonce = base64.StdEncoding.EncodeToString(nonce)
	v.Cipher.Data = base64.StdEncoding.EncodeToString(ciphertext)
}

// decodeVaultCipher decodes the base64-encoded cipher components
func (v *VaultFile) decodeCipher() (salt, nonce, ciphertext []byte, err error) {
	salt, err = base64.StdEncoding.DecodeString(v.KDF.Salt)
	if err != nil {
		return nil, nil, nil, err
	}

	nonce, err = base64.StdEncoding.DecodeString(v.Cipher.Nonce)
	if err != nil {
		return nil, nil, nil, err
	}

	ciphertext, err = base64.StdEncoding.DecodeString(v.Cipher.Data)
	if err != nil {
		return nil, nil, nil, err
	}

	return salt, nonce, ciphertext, nil
}

func InitVault(path string) error {
	masterPwd, err := promptPassword("Choose a master password: \n")
	if err != nil {
		return err
	}

	salt, err := randomBytes(saltSize)
	if err != nil {
		return err
	}

	key := deriveKey(masterPwd, salt)
	zero(masterPwd)

	plaintext := []byte(`{"entries":[]}`)
	nonce, ciphertext, err := encrypt(key, plaintext)
	if err != nil {
		return err
	}

	zero(plaintext)
	zero(key)

	vault := newVaultFile(salt, nonce, ciphertext)
	return writeVaultFile(path, vault)
}

func unlockVault(path string) ([]byte, *VaultFile, []byte, error) {
	vault, err := readVaultFile(path)
	if err != nil {
		return nil, nil, nil, err
	}

	masterPwd, err := promptPassword("Enter your master password: \n")
	if err != nil {
		return nil, nil, nil, err
	}

	salt, nonce, ciphertext, err := vault.decodeCipher()
	if err != nil {
		return nil, nil, nil, err
	}

	key := deriveKey(masterPwd, salt)
	zero(masterPwd)

	plaintext, err := decrypt(key, nonce, ciphertext)
	if err != nil {
		zero(key)
		return nil, nil, nil, err
	}

	return plaintext, vault, key, nil
}

func AddEntry(path string) error {
	notes, err := promptNormal("What is this entry for: ")
	if err != nil {
		return err
	}
	username, err := promptNormal("Enter username: ")
	if err != nil {
		return err
	}
	password, err := promptPassword("Enter password: \n")
	if err != nil {
		return err
	}

	plaintext, vault, key, err := unlockVault(path)
	if err != nil {
		return err
	}
	defer zero(key)

	var data VaultData
	if err := json.Unmarshal(plaintext, &data); err != nil {
		return err
	}

	newEntry := Entry{
		Username: username,
		Password: string(password),
		Notes:    notes,
	}
	data.Entries = append(data.Entries, newEntry)

	updatedPlaintext, err := json.Marshal(data)
	if err != nil {
		return err
	}

	nonce, ciphertext, err := encrypt(key, updatedPlaintext)
	if err != nil {
		return err
	}

	vault.updateCipher(nonce, ciphertext)
	return writeVaultFile(path, vault)
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
