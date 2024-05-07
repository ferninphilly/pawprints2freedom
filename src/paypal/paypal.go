package paypal

import (
	"fmt"
	"os"

	"github.com/plutov/paypal/v4"
	"gopkg.in/ini.v1"
)

type ConfigData struct {
	ClientID string `ini:"clientID"`
	Secret   string `ini:"secret"`
}

func main() {
	var config ConfigData
	populateConfig(config)
	// Create a client instance
	c, err := paypal.NewClient(config.ClientID, config.Secret, paypal.APIBaseSandBox)
	if err != nil {
		fmt.Println("Error initializing paypal client due to ", err)
	}
	c.SetLog(os.Stdout)
}

func populateConfig(config ConfigData) {
	inidata, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println("Error reading the config due to ", err.Error())
	}
	inidata.MapTo(&config)
}
