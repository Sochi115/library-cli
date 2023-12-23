package initializer

import (
	"log"
	"os"
)

func GenerateEnvironmentFolder() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	path := dirname + "\\library_cli"

	_, err = os.Stat(path)
	if err == nil {
		return
	}

	if os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}
	os.MkdirAll(path, os.ModePerm)
}

func GetEnvFolderPath() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return dirname + "\\library_cli"
}
