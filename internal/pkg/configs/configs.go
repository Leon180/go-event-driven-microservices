package configs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/environments"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/reflect"
	goenv "github.com/caarlos0/env/v8"
	"github.com/spf13/viper"
)

// BindConfig will bind config file with specific env to struct T
// the config file name should be config.[env].json
func BindConfig[T any](env enums.Environment) (T, error) {
	return BindConfigByKey[T]("", env)
}

// BindConfigByKey will bind specific config with key: configKey in config file with specific env to struct T
// the config file name should be config.[env].json
// the configKey is the key of the config in the config file
func BindConfigByKey[T any](configKey string, env enums.Environment) (T, error) {
	if !env.IsValid() {
		env = enums.EnvironmentDevelopment
	}

	cfg := reflect.GetInstance[T]()

	viper.SetDefault(enums.ConfigPath, "")

	configPath := viper.GetString(enums.ConfigPath)
	if configPath == "" {
		appRootPath := viper.GetString(enums.AppRootPath)
		if appRootPath == "" {
			appRootPath, _ = environments.FindProjectRootWorkingDirectory()
		}
		if appRootPath == "" {
			log.Printf("appRootPath is empty")
			return *new(T), customizeerrors.DirectoryNotFoundError
		}
		dir, err := findConfigFileDir(appRootPath, env)
		if err != nil {
			log.Printf("error find config file dir: %v", err)
			return *new(T), err
		}
		configPath = dir
	}

	// https://github.com/spf13/viper/issues/390#issuecomment-718756752
	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.AddConfigPath(configPath)
	viper.SetConfigType(string(enums.Json))

	if err := viper.ReadInConfig(); err != nil {
		return *new(T), err
	}

	isPointer := reflect.IsPointer[T]()

	if isPointer {
		if configKey == "" {
			if err := viper.Unmarshal(cfg); err != nil {
				log.Printf("error unmarshal config: %v", err)
				return *new(T), err
			}
		} else {
			if err := viper.UnmarshalKey(configKey, cfg); err != nil {
				log.Printf("error unmarshal config key: %v", err)
				return *new(T), err
			}
		}
	} else {
		if configKey == "" {
			if err := viper.Unmarshal(&cfg); err != nil {
				log.Printf("error unmarshal config: %v", err)
				return *new(T), err
			}
		} else {
			if err := viper.UnmarshalKey(configKey, &cfg); err != nil {
				log.Printf("error unmarshal config key: %v", err)
				return *new(T), err
			}
		}
	}

	viper.AutomaticEnv()

	if isPointer {
		if err := goenv.Parse(cfg); err != nil {
			return *new(T), err
		}
	} else {
		if err := goenv.Parse(&cfg); err != nil {
			return *new(T), err
		}
	}

	return cfg, nil
}

func findConfigFileDir(
	rootDir string,
	env enums.Environment,
) (string, error) {
	var result string
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.EqualFold(info.Name(), fmt.Sprintf("config.%s%s", env, enums.FileNameExtensionJson)) ||
			strings.EqualFold(info.Name(), fmt.Sprintf("config.%s%s", env, enums.FileNameExtensionYaml)) ||
			strings.EqualFold(info.Name(), fmt.Sprintf("config.%s%s", env, enums.FileNameExtensionYml)) {
			result = filepath.Dir(path)
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		log.Printf("error walk config file dir: %v", err)
		return "", err
	}
	if result == "" {
		log.Printf("config file dir not found")
		return "", customizeerrors.DirectoryNotFoundError
	}
	return result, nil
}
