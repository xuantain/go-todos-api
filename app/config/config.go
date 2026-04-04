package config

import (
	"bytes"
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

//go:embed config_default.yml config_dev.yml config_prod.yml config_test.yml
var embeddedConfigs embed.FS

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")

	// Load default config from embedded FS
	defaultData, err := embeddedConfigs.ReadFile("config_default.yml")
	if err != nil {
		log.Fatal("error reading embedded default config: ", err)
	}
	if err = config.ReadConfig(bytes.NewReader(defaultData)); err != nil {
		log.Fatal("error parsing default config: ", err)
	}

	// Merge env-specific config from embedded FS
	envFile := fmt.Sprintf("config_%s.yml", env)
	envData, err := embeddedConfigs.ReadFile(envFile)
	if err != nil {
		log.Fatal("error reading embedded env config: ", err)
	}
	envConfig := viper.New()
	envConfig.SetConfigType("yaml")
	if err = envConfig.ReadConfig(bytes.NewReader(envData)); err != nil {
		log.Fatal("error parsing env config: ", err)
	}
	config.MergeConfigMap(envConfig.AllSettings())

	// Merge parameters.yml from disk (secrets / deployment-specific overrides)
	// Search order: next to the binary, then ./config/, then config/ relative to cwd
	paramsPaths := []string{
		"parameters.yml",
		"config/parameters.yml",
	}
	if exe, err := os.Executable(); err == nil {
		paramsPaths = append([]string{exe + "/../parameters.yml"}, paramsPaths...)
	}

	params := viper.New()
	params.SetConfigType("yaml")
	loaded := false
	for _, p := range paramsPaths {
		params.SetConfigFile(p)
		if err := params.ReadInConfig(); err == nil {
			config.MergeConfigMap(params.AllSettings())
			loaded = true
			break
		}
	}
	if !loaded {
		log.Println("warning: parameters.yml not found, using defaults only")
	}
}

func GetConfig() *viper.Viper {
	return config
}
