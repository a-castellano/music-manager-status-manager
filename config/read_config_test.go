package config

import (
	"os"
	"testing"
)

func TestProcessNoConfigFilePresent(t *testing.T) {

	_, err := ReadConfig()
	if err == nil {
		t.Errorf("ReadConfig method without any valid config file should fail.")
	} else {
		if err.Error() != "Fatal error reading config file: Config File \"config\" Not Found in \"[/etc/music-manager]\"" {
			t.Errorf("Default config should be in /etc/music-manager/config.toml, not in other place, error was '%s'.", err.Error())
		}
	}
}

func TestProcessStorageNoDataInConfig(t *testing.T) {
	os.Setenv("MUSIC_MANAGER_SERVICE_CONFIG_FILE_LOCATION", "./config_files_test/storage_no_data/")
	_, err := ReadConfig()
	if err == nil {
		t.Errorf("ReadConfig method without storage data should fail.")
	} else {
		requiredError := "Fatal error reading config: no storage servicename was found."
		if err.Error() != requiredError {
			t.Errorf("Error should be \"%s\" but error was '%s'.", requiredError, err.Error())
		}
	}
}

func TestProcessWebserverNoData(t *testing.T) {
	os.Setenv("MUSIC_MANAGER_SERVICE_CONFIG_FILE_LOCATION", "./config_files_test/webserver_no_data/")
	_, err := ReadConfig()
	if err == nil {
		t.Errorf("ReadConfig method without webserver data should fail.")
	} else {
		requiredError := "Fatal error reading config: no webserver port was found."
		if err.Error() != requiredError {
			t.Errorf("Error should be \"%s\" but error was '%s'.", requiredError, err.Error())
		}
	}
}

func TestProcessNoWebserver(t *testing.T) {
	os.Setenv("MUSIC_MANAGER_SERVICE_CONFIG_FILE_LOCATION", "./config_files_test/no_webserver/")
	_, err := ReadConfig()
	if err == nil {
		t.Errorf("ReadConfig method without webserver config should fail.")
	} else {
		requiredError := "Fatal error reading config: no webserver config was found."
		if err.Error() != requiredError {
			t.Errorf("Error should be \"%s\" but error was '%s'.", requiredError, err.Error())
		}
	}
}

func TestProcessNoStorage(t *testing.T) {
	os.Setenv("MUSIC_MANAGER_SERVICE_CONFIG_FILE_LOCATION", "./config_files_test/no_storage/")
	_, err := ReadConfig()
	if err == nil {
		t.Errorf("ReadConfig method without storage config should fail.")
	} else {
		requiredError := "Fatal error reading config: no storage config was found."
		if err.Error() != requiredError {
			t.Errorf("Error should be \"%s\" but error was '%s'.", requiredError, err.Error())
		}
	}
}

func TestValisConfig(t *testing.T) {
	os.Setenv("MUSIC_MANAGER_SERVICE_CONFIG_FILE_LOCATION", "./config_files_test/valid_config/")
	config, err := ReadConfig()
	if err != nil {
		t.Errorf("ReadConfig method with valid config shouldn't fail.")
	}
	if config.Storage.ServiceName != "redis" {
		t.Errorf("config.Storage.ServiceName shold be 'redis' not '%s'", config.Storage.ServiceName)
	}
	if config.WebServer.Port != 8080 {
		t.Errorf("config.WebServer.Port shold be '8080' not '%d'", config.WebServer.Port)
	}

}
