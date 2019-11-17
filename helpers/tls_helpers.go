package helpers

import (
	"log"
	"os"
	"path/filepath"
)

func TlsCerts(relPath, cert, key string) (string, string) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("unable to determine current directory")
	}

	certFile := filepath.Join(pwd, relPath, cert)
	keyFile := filepath.Join(pwd, relPath, key)

	return certFile, keyFile
}
