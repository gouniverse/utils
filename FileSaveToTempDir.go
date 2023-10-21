package utils

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

func FileSaveToTempDir(fileName string, file multipart.File) (string, error) {
	extension := filepath.Ext(fileName)
	out, pathError := os.CreateTemp(os.TempDir(), "temp-*"+extension)

	if pathError != nil {
		log.Println("Error Creating a file for writing", pathError)
		return "", pathError
	}
	defer out.Close()

	_, errCopy := io.Copy(out, file)
	if errCopy != nil {
		return "", errCopy
	}

	return out.Name(), nil
}
