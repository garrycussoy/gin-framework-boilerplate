package config

import (
	"gin-framework-boilerplate/internal/constants"

	"github.com/spf13/viper"
)

var AppConfig Config

// Let's define all config variables here
type Config struct {
	AppHost       string `mapstructure:"APP_HOST"`
	Port          int    `mapstructure:"PORT"`
	Environment   string `mapstructure:"ENVIRONMENT"`
	Debug         bool   `mapstructure:"DEBUG"`
	DatabaseDebug bool   `mapstructure:"DATABASE_DEBUG"`

	WebHost            string `mapstructure:"WEB_HOST"`
	ChangePasswordPath string `mapstructure:"CHANGE_PASSWORD_PATH"`

	DBPostgreDriver string `mapstructure:"DB_POSTGRE_DRIVER"`
	DBPostgreDsn    string `mapstructure:"DB_POSTGRE_DSN"`
	DBPostgreURL    string `mapstructure:"DB_POSTGRE_URL"`

	JWTSecret                   string `mapstructure:"JWT_SECRET"`
	JWTExpired                  int    `mapstructure:"JWT_EXPIRED"`
	JWTIssuer                   string `mapstructure:"JWT_ISSUER"`
	ChangePasswordTokenLifespan int    `mapstructure:"CHANGE_PASSWORD_TOKEN_LIFESPAN"`

	ServerReadTimeout  int `mapstructure:"SERVER_READ_TIMEOUT"`
	ServerWriteTimeout int `mapstructure:"SERVER_WRITE_TIMEOUT"`
	HandlerTimeout     int `mapstructure:"HANDLER_TIMEOUT"`

	REDISHost     string `mapstructure:"REDIS_HOST"`
	REDISPassword string `mapstructure:"REDIS_PASS"`
	REDISExpired  int    `mapstructure:"REDIS_EXPIRED"`

	AllowedCORS string `mapstructure:"ALLOWED_CORS"`

	EmailSender   string `mapstructure:"EMAIL_SENDER"`
	EmailPassword string `mapstructure:"EMAIL_PASSWORD"`
}

// Function to load config from stated source
func InitializeAppConfig(isTesting bool) error {
	if isTesting {
		// Load config for testing environment
		viper.SetConfigName("test.env")
		viper.SetConfigType("env")
		viper.AddConfigPath(".")
		viper.AllowEmptyEnv(true)
		viper.AutomaticEnv()
	} else {
		// Default config source
		viper.SetConfigName(".env") // Allow directly reading from .env file
		viper.SetConfigType("env")
		viper.AddConfigPath(".")
		viper.AllowEmptyEnv(true)
		viper.AutomaticEnv()
	}

	// Process to load the config
	err := viper.ReadInConfig()
	if err != nil {
		return constants.ErrLoadConfig
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		return constants.ErrParseConfig
	}

	// Ensure some required variables have value
	if AppConfig.Port == 0 || AppConfig.Environment == "" || AppConfig.JWTSecret == "" || AppConfig.JWTExpired == 0 || AppConfig.JWTIssuer == "" || AppConfig.EmailSender == "" || AppConfig.EmailPassword == "" || AppConfig.REDISHost == "" || AppConfig.REDISPassword == "" || AppConfig.REDISExpired == 0 || AppConfig.DBPostgreDriver == "" {
		return constants.ErrEmptyVar
	}

	switch AppConfig.Environment {
	case constants.EnvironmentDevelopment:
		if AppConfig.DBPostgreDsn == "" {
			return constants.ErrEmptyVar
		}
	case constants.EnvironmentProduction:
		if AppConfig.DBPostgreURL == "" {
			return constants.ErrEmptyVar
		}
	}

	return nil
}
