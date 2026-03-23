package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(ConfigDir())

	viper.SetEnvPrefix("HA")
	viper.AutomaticEnv()

	viper.SetDefault("url", "")

	_ = viper.ReadInConfig()
}

func ConfigDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "ha-cli")
}

func URL() string {
	return viper.GetString("url")
}

func Token() string {
	return viper.GetString("token")
}

func Save() error {
	dir := ConfigDir()
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}
	return viper.WriteConfigAs(filepath.Join(dir, "config.yaml"))
}
