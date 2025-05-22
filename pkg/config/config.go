package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"

	cmn "ng_autocode_back_service/common/model"
	cmnUtils "ng_autocode_back_service/common/util"
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

func GetConfig() *viper.Viper {
	return config
}

func LoadConfig(appType cmn.APP_TYPE, env string, configPaths ...string) *viper.Viper {
	// Set the environment variable for the config file

	config = viper.New()
	config.SetConfigName(env)
	config.SetConfigType("yaml")

	cwd, _ := os.Getwd()
	sep := string(cmnUtils.GetPathSeparator())
	cwd = strings.Replace(cwd, sep+"cmd", "", -1)
	cwd = strings.Replace(cwd, sep+"app", "", -1)
	cwd = strings.Replace(cwd, sep+"bin", "", -1)
	cwd = strings.Replace(cwd, sep+"win", "", -1)
	cwd = strings.Replace(cwd, sep+"linux", "", -1)
	cwd = strings.Replace(cwd, sep+"debug", "", -1)
	cwd = strings.Replace(cwd, sep+"release", "", -1)
	envPath := cwd + sep + "app" + sep + "conf" + sep

	config.AddConfigPath(envPath)
	config.AddConfigPath(".")
	if len(configPaths) > 0 {
		for _, path := range configPaths {
			config.AddConfigPath(path)
		}
	} else {
		config.AddConfigPath(envPath)
	}

	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	return config
}
