package tests

import (
	utils "github.com/ferninphilly/pawprints2freedom/utils"

	"reflect"
	"testing"
)

func TestReadConfig(t *testing.T) {
	var cfg utils.Config
	c := &cfg
	if err := utils.GetConfigData(c, "config_test.ini"); err != nil {
		t.Errorf("Error getting config data due to %s", err.Error())
	}
	want := &utils.Config{struct {
		BaseURL   string "ini:\"baseURL\""
		AppID     string "ini:\"appID\""
		AppSecret string "ini:\"appSecret\""
		UserToken string "ini:\"userToken\""
	}{"https://iamabaseurl", "iamanappid", "shhhIamASecret", "iAmAUserToken"}}
	got := c
	matchOrNot := reflect.DeepEqual(got, want)
	if matchOrNot != true {
		t.Errorf("GetConfig failed\nGot: %v\nWant:%v", got, want)
	}
}
