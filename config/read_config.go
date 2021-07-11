package config

import (
	"errors"

	viperLib "github.com/spf13/viper"
)

type Storage struct {
	ServiceName string
	ServicePort int
	ServiceDB   int
}

type WebServer struct {
	Port int
}

type Config struct {
	Storage   Storage
	WebServer WebServer
}

func ReadConfig() (Config, error) {
	var configFileLocation string
	var config Config

	var envVariable string = "MUSIC_MANAGER_SERVICE_CONFIG_FILE_LOCATION"

	storageVariables := []string{"servicename", "serviceport", "servicedb"}
	webServerVariables := []string{"port"}

	requiredConfigEntities := []string{"storage", "webserver"}

	viper := viperLib.New()

	//Look for config file location defined as env var
	viper.BindEnv(envVariable)
	configFileLocation = viper.GetString(envVariable)
	if configFileLocation == "" {
		// Get config file from default location
		configFileLocation = "/etc/music-manager/"
	}
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(configFileLocation)

	if err := viper.ReadInConfig(); err != nil {
		return config, errors.New(errors.New("Fatal error reading config file: ").Error() + err.Error())
	}

	for _, requiredConfigEntity := range requiredConfigEntities {
		if !viper.IsSet(requiredConfigEntity) {
			return config, errors.New("Fatal error reading config: no " + requiredConfigEntity + " config was found.")
		}
	}

	for _, storage_variable := range storageVariables {
		if !viper.IsSet("storage." + storage_variable) {
			return config, errors.New("Fatal error reading config: no storage " + storage_variable + " was found.")
		}
	}

	storage := Storage{ServiceName: viper.GetString("storage.servicename"), ServicePort: viper.GetInt("storage.serviceport"), ServiceDB: viper.GetInt("storage.servicedb")}

	config.Storage = storage

	for _, webServer_variable := range webServerVariables {
		if !viper.IsSet("webserver." + webServer_variable) {
			return config, errors.New("Fatal error reading config: no webserver " + webServer_variable + " was found.")
		}
	}

	webServer := WebServer{Port: viper.GetInt("webserver.port")}

	config.WebServer = webServer

	return config, nil
}
