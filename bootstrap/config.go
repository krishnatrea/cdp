package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppEnv              string `mapstructure:"APP_ENV"`
	PostgresUser        string `mapstructure:"POSTGRES_USER"`
	PostgresPassword    string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDB          string `mapstructure:"POSTGRES_DB"`
	PostgresHost        string `mapstructure:"POSTGRES_HOST"`
	ContextTimeout      int    `mapstructure:"CONTEXT_TIMEOUT"`
	ServerAddress       string `mapstructure:"SERVER_ADDRESS"`
	AccessTokenSecret   string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret  string `mapstructure:"REFRESH_TOKEN_SECRET"`
	RefreshTokenTimeOut int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenTimeOut  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`

	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     int    `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
}

func NewConfig() *Config {
	env := Config{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
