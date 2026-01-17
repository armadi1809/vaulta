package config

import (
	"os"
	"path/filepath"
	"runtime"
)

func DefaultVaultPath() (string, error) {
	if p := os.Getenv("VAULTA_VAULT_PATH"); p != "" {
		return p, nil
	}

	var dataDir string

	switch runtime.GOOS {
	case "windows":
		// Windows
		dataDir = os.Getenv("LOCALAPPDATA")
		if dataDir == "" {
			dataDir = os.Getenv("APPDATA")
		}
		if dataDir == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				return "", err
			}
			dataDir = filepath.Join(home, "AppData", "Local")
		}
	case "darwin":
		// macOS
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		dataDir = filepath.Join(home, "Library", "Application Support")
	default:
		// Linux/Unix
		dataDir = os.Getenv("XDG_DATA_HOME")
		if dataDir == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				return "", err
			}
			dataDir = filepath.Join(home, ".local", "share")
		}
	}

	return filepath.Join(dataDir, "vaulta", "vault.json"), nil
}
