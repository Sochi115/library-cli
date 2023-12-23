package main

import (
	"github.com/Sochi115/library-cli/app"
	"github.com/Sochi115/library-cli/initializer"
)

func main() {
	initializer.GenerateEnvironmentFolder()
	envFolderpath := initializer.GetEnvFolderPath()
	app.StartApp(envFolderpath)
}
