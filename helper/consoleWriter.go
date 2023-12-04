package helper

import (
	"io"
	"os"
)

func WriteStringToConsole(msg string) {
	io.WriteString(os.Stdout, msg)
	io.WriteString(os.Stdout, "\n")
}
