package utils

import (
	"log"
	"os"

	"github.com/gouniverse/envenc"
)

func EnvEncInitialize(password string) {
	if FileExists(".env.vault") {
		keys, err := envenc.EnvList(".env.vault", password)
		if err != nil {
			log.Println(err.Error())
		}

		for k, v := range keys {
			os.Setenv(k, v)
		}
	}
}
