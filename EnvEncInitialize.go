package utils

import (
	"errors"
	"log"
	"os"

	"github.com/gouniverse/envenc"
)

func EnvEncInitialize(password string, vaultFilePath string) error {
	if vaultFilePath == "" {
		vaultFilePath = ".env.vault"
	}

	if !FileExists(vaultFilePath) {
		return errors.New("Vault file not found: " + vaultFilePath)
	}

	keys, err := envenc.EnvList(vaultFilePath, password)

	if err != nil {
		log.Println(err.Error())
	}

	for k, v := range keys {
		err := os.Setenv(k, v)

		if err != nil {
			return err
		}
	}

	return nil
}
