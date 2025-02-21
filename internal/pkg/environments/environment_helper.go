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

func findProjectRootDirectoryFromProjectName(projectName string) (string, error) {
	currecntDir, _ := os.Getwd()
	parentDir := filepath.Dir(currecntDir)
	for {
		if strings.HasSuffix(currecntDir, projectName) {
			return currecntDir, nil
		}
		if currecntDir == "" || parentDir == currecntDir {
			log.Printf("project root directory not found")
			return "", customizeerrors.DirectoryNotFoundError
		}
		currecntDir, parentDir = parentDir, filepath.Dir(parentDir)
	}
}

func findRootDirectory(currentDirectory string) (string, error) {
	files, err := os.ReadDir(currentDirectory)
	if err != nil {
		log.Printf("error read directory: %v", err)
		return "", err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.EqualFold(file.Name(), "go.mod") {
			return currentDirectory, nil
		}
	}

	parentDir := filepath.Dir(currentDirectory)
	if parentDir == currentDirectory {
		log.Printf("root directory not found")
		return "", customizeerrors.DirectoryNotFoundError
	}

	return findRootDirectory(parentDir)
}
