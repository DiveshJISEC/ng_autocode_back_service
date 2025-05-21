package config

import (
	"log"
	"os"

	"github.com/spf13/viper"

	cmn "ng_autocode_back_service/common/model"
)

var config *viper.Viper

const (
	APP_TYPE      = "app_type"
	USERID        = "userId"
	REQUESTID     = "requestID"
	TOKEN         = "token"
	UCC           = "ucc"
	AUTHORIZATION = "Authorization"
	XLENGTH       = "X-Length"
	SCOPE         = "scope"
	AccessToken   = "accessToken"
	IV256         = "iv"
)

func getConfig() *viper.Viper {
	return config
}

func LoadConfig(appType cmn.APP_TYPE, env string, configPaths ...string) *viper.Viper {

	var err error
	// Set the environment variable for the config file

	config = viper.New()
	config.SetConfigName(env)
	config.SetConfigType("yaml")

	cwd, _ := os.Getwd()
	sep := string(os.PathSeparator)

	config.AddConfigPath(envPath)

	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Set default values
	setDefaultValues()

	// Read environment variables
	readEnvVariables()
}
