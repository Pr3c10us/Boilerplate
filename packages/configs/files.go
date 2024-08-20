package configs

import (
	"os"
	"regexp"
)

func GetRootPath() string {
	projectDirName := "Boilerplate"
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	return string(projectName.Find([]byte(currentWorkDirectory)))
}
