package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type Config struct {
	Facebook struct {
		BaseURL   string `ini:"baseURL"`
		AppID     string `ini:"appID"`
		AppSecret string `ini:"appSecret"`
		UserToken string `ini:"userToken"`
	} `ini:"facebook"`
}

func GetConfigData(config *Config) error {
	inidata, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println("Error reading the config due to ", err.Error())
		return err
	}
	if err2 := inidata.MapTo(&config); err != nil {
		fmt.Println("Error mapping to the fb config data due to ", err2.Error())
		return err2
	}
	return nil
}
