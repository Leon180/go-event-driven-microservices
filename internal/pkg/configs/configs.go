package configs

import (
	"fmt"
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

	viper.SetDefault(enums.ConfigPath, "")
	configPath := findConfigPath(env)
	if configPath == "" {
		return *new(T), customizeerrors.DirectoryNotFoundError
	}
	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.AddConfigPath(configPath)
	viper.SetConfigType(string(enums.Json))

	if err := viper.ReadInConfig(); err != nil {
		return *new(T), err
	}

	cfg := reflect.GetInstance[T]()
	isPointer := reflect.IsPointerV2(cfg)
	if isPointer {
		if err := unmarchalConfig[T](cfg, configKey); err != nil {
			return *new(T), err
		}
	} else {
		if err := unmarchalConfig[T](&cfg, configKey); err != nil {
			return *new(T), err
		}
	}

	viper.AutomaticEnv()

	if isPointer {
		if err := parseEnv[T](cfg); err != nil {
			return *new(T), err
		}
	} else {
		if err := parseEnv[T](&cfg); err != nil {
			return *new(T), err
		}
	}

	return cfg, nil
}

func findConfigPath(env enums.Environment) string {
	configPath := viper.GetString(enums.ConfigPath)
	if configPath != "" {
		return configPath
	}
	return findConfigPathByRootPath(env)
}

func findConfigPathByRootPath(env enums.Environment) string {
	appRootPath := viper.GetString(enums.AppRootPath)
	if appRootPath == "" {
		appRootPath, _ = environments.FindProjectRootWorkingDirectory()
	}
	if appRootPath == "" {
		return ""
	}
	var result string
	filepath.Walk(appRootPath, func(path string, info os.FileInfo, err error) error {
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
	return result
}

// unmarchalConfig will unmarchal config from viper to struct T, the T should be a pointer
func unmarchalConfig[T any](cfg any, configKey string) error {
	if configKey == "" {
		return viper.Unmarshal(cfg)
	}
	return viper.UnmarshalKey(configKey, cfg)
}

// parseEnv will parse env from viper to struct T, the T should be a pointer
func parseEnv[T any](cfg any) error {
	return goenv.Parse(cfg)
}
