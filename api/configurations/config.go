package configurations

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver        string `mapstructure:"DB_DRIVER"`
	DBSource        string `mapstructure:"DB_SOURCE"`
	ServerAddress   string `mapstructure:"SERVER_ADDRESS"`
	DBAutoMigration bool   `mapstructure:"DB_AUTO_MIGRATION"`
}

func LoadConfig(path string) (Config, error) {
	var config Config
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		config.DBAutoMigration = viper.GetBool("DB_AUTO_MIGRATION")
		config.DBDriver = viper.GetString("DB_DRIVER")
		config.DBSource = viper.GetString("DB_SOURCE")
		config.ServerAddress = viper.GetString("SERVER_ADDRESS")
		return config, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}
	return config, nil
}
