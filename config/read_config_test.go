package config

import (
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
