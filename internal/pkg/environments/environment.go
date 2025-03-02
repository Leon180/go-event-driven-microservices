package environments

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/spf13/viper"
)

func InitEnv() enums.Environment {
	env := enums.EnvironmentDevelopment

	// setup viper to read from os environment with `viper.Get`
	viper.AutomaticEnv()

	if err := loadEnvFilesRecursive(enums.FileNameExtensionEnv); err != nil {
		log.Printf(".env file cannot be found, err: %v", err)
	}

	if err := setRootWorkingDirectoryEnvironment(); err != nil {
		log.Printf("Failed to set root working directory environment, err: %v", err)
	}

	if err := fixProjectRootWorkingDirectoryPath(); err != nil {
		log.Printf("Failed to fix project root working directory path, err: %v", err)
	}

	manualEnv := os.Getenv(enums.AppEnv)

	if manualEnv != "" && enums.Environment(manualEnv).IsValid() {
		env = enums.Environment(manualEnv)
	}

	return env
}

func FindProjectRootWorkingDirectory() (string, error) {
	var rootDirectory string
	projectNameEnv := viper.GetString(enums.ProjectNameEnv)
	if projectNameEnv != "" {
		rootDir, err := findProjectRootDirectoryFromProjectName(projectNameEnv)
		if err != nil {
			return "", err
		}
		rootDirectory = rootDir
	} else {
		currecntDir, _ := os.Getwd()
		rootDir, err := findRootDirectory(currecntDir)
		if err != nil {
			return "", err
		}
		rootDirectory = rootDir
	}

	absoluteRootWorkingDirectory, _ := filepath.Abs(rootDirectory)

	return absoluteRootWorkingDirectory, nil
}
