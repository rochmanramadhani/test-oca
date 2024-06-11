package lib

import (
	"fmt"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Name     string
	HTTPPort int
}

type RateCarType struct {
	SUV int
	MPV int
}

type Config struct {
	App           AppConfig
	RateCarType   RateCarType
	DataFile      string
	MaxParkingLot int
	TimeZone      string
}

var Cfg Config

func getAppConfig() AppConfig {
	return AppConfig{
		Name:     getStringOrPanic("APP_NAME"),
		HTTPPort: getIntOrPanic("APP_HTTP_PORT"),
	}
}

func getRateCarType() RateCarType {
	return RateCarType{
		SUV: getIntOrPanic("RATE_SUV"),
		MPV: getIntOrPanic("RATE_MPV"),
	}
}

func LoadConfigByFile(path, fileName, fileType string) Config {
	viper.SetConfigName(fileName)
	viper.SetConfigType(fileType)
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file, use automatic environment instead %s\n", err)
	}

	Cfg = Config{
		App:           getAppConfig(),
		RateCarType:   getRateCarType(),
		DataFile:      viper.GetString("DATA_FILE"),
		MaxParkingLot: viper.GetInt("MAX_PARKING_LOT"),
		TimeZone:      viper.GetString("TIME_ZONE"),
	}

	return Cfg
}
