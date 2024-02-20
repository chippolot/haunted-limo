package api

import (
	"os"
	"path"
)

func GetTemplatePath(filename string) string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return path.Join(cwd, "templates", filename)
}
