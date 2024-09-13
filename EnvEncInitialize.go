package utils

import (
	"log"
	"os"

	"github.com/gouniverse/envenc"
)

func EnvEncInitialize(password string, vaultFilePath string) {
	if vaultFilePath == "" {
		vaultFilePath = ".env.vault"
	}

	if FileExists(vaultFilePath) {
		keys, err := envenc.EnvList(vaultFilePath, password)
		if err != nil {
			log.Println(err.Error())
		}

		for k, v := range keys {
			os.Setenv(k, v)
		}
	}
}
