package enums

type Environment string

const (
	EnvironmentNull        Environment = ""
	EnvironmentDevelopment Environment = "development"
	EnvironmentTest        Environment = "test"
	EnvironmentProduction  Environment = "production"
)

func (env Environment) IsNull() bool {
	return env == EnvironmentNull
}

func (env Environment) IsDevelopment() bool {
	return env == EnvironmentDevelopment
}

func (env Environment) IsTest() bool {
	return env == EnvironmentTest
}

func (env Environment) IsProduction() bool {
	return env == EnvironmentProduction
}

func (env Environment) GetEnvironmentName() string {
	return string(env)
}

func (env Environment) IsValid() bool {
	_, ok := environmentMap[env]
	return ok
}

var environmentMap = map[Environment]string{
	EnvironmentDevelopment: "development",
	EnvironmentTest:        "test",
	EnvironmentProduction:  "production",
}

const (
	AppEnv         = "APP_ENV"
	AppRootPath    = "APP_ROOT"
	ProjectNameEnv = "PROJECT_NAME"
	ConfigPath     = "CONFIG_PATH"
	Json           = "json"
)
