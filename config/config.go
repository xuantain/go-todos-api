package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName("config_default")
	config.AddConfigPath("config/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing default configuration file")
	}

	envConfig := viper.New()
	envConfig.SetConfigType("yaml")
	envConfig.SetConfigName("config_" + env)
	envConfig.AddConfigPath("config/")
	err = envConfig.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing env configuration file")
	}

	params := viper.New()
	params.SetConfigType("yaml")
	params.SetConfigName("parameters")
	params.AddConfigPath("config/")
	err = params.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing parameters file")
	}

	config.MergeConfigMap(envConfig.AllSettings())
	config.MergeConfigMap(params.AllSettings())
}

// func relativePath(basedir string, path *string) {
// 	p := *path
// 	if len(p) > 0 && p[0] != '/' {
// 		*path = filepath.Join(basedir, p)
// 	}
// }

func GetConfig() *viper.Viper {
	return config
}
