package environments

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/joho/godotenv"
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
	if enums.Environment(manualEnv).IsValid() {
		env = enums.Environment(manualEnv)
	}

	return env
}

func loadEnvFilesRecursive(envFileNameExtension enums.FileNameExtension) error {
	currecntDir, err := os.Getwd()
	if err != nil {
		log.Printf("error get working directory: %v", err)
		return err
	}

	for {
		envFilePath := filepath.Join(currecntDir, envFileNameExtension.String())
		if err := godotenv.Load(envFilePath); err == nil {
			return nil
		}
		parentDir := filepath.Dir(currecntDir)
		if parentDir == currecntDir {
			break
		}
		currecntDir = parentDir
	}

	return customizeerrors.FileNotFoundError
}

func setRootWorkingDirectoryEnvironment() error {
	rootDir, err := FindProjectRootWorkingDirectory()
	if err != nil {
		log.Printf("error find project root working directory: %v", err)
		return err
	}
	viper.Set(enums.AppRootPath, rootDir)
	return nil
}

func fixProjectRootWorkingDirectoryPath() error {
	rootDir, err := FindProjectRootWorkingDirectory()
	if err != nil {
		log.Printf("error find project root working directory: %v", err)
		return err
	}
	return os.Chdir(rootDir)
}

func FindProjectRootWorkingDirectory() (string, error) {
	rootDirectory := findProjectRootDirectory()
	if rootDirectory == "" {
		return "", customizeerrors.DirectoryNotFoundError
	}
	absoluteRootWorkingDirectory, _ := filepath.Abs(rootDirectory)
	return absoluteRootWorkingDirectory, nil
}

func findProjectRootDirectory() string {
	projectNameEnv := viper.GetString(enums.ProjectNameEnv)
	if projectNameEnv != "" {
		return findProjectRootDirectoryFromProjectName(projectNameEnv)
	}
	currecntDir, _ := os.Getwd()
	return findRootDirectory(currecntDir)
}

func findProjectRootDirectoryFromProjectName(projectName string) string {
	currecntDir, _ := os.Getwd()
	parentDir := filepath.Dir(currecntDir)
	for {
		if strings.HasSuffix(currecntDir, projectName) {
			return currecntDir
		}
		if currecntDir == "" || parentDir == currecntDir {
			log.Printf("project root directory not found")
			return ""
		}
		currecntDir, parentDir = parentDir, filepath.Dir(parentDir)
	}
}

func findRootDirectory(currentDirectory string) string {
	files, err := os.ReadDir(currentDirectory)
	if err != nil {
		log.Printf("error read directory: %v", err)
		return ""
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.EqualFold(file.Name(), "go.mod") {
			return currentDirectory
		}
	}

	parentDir := filepath.Dir(currentDirectory)
	if parentDir == currentDirectory {
		log.Printf("root directory not found")
		return ""
	}

	return findRootDirectory(parentDir)
}
