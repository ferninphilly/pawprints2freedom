package tests

import (
	"fmt"
	"reflect"
	"testing"

	fb "github.com/ferninphilly/pawprints2freedom/facebook"
	util "github.com/ferninphilly/pawprints2freedom/utils"
	"github.com/stretchr/testify/mock"
)

type SendFunc struct {
	mock.Mock
}

func TestNewFBEndpoints(t *testing.T) *fb.FBEndpoints {
	var cfg util.Config
	c := &cfg
	if err := util.GetConfigData(c, "../utils/config_test.ini"); err != nil {
		fmt.Println("There was an error getting the fb config data due to ", err.Error())
	}
	want := &util.Config{struct {
		BaseURL   string "ini:\"baseURL\""
		AppID     string "ini:\"appID\""
		AppSecret string "ini:\"appSecret\""
		UserToken string "ini:\"userToken\""
	}{"https://iamabaseurl", "iamanappid", "shhhIamASecret", "iAmAUserToken"}}
	got := &fb.FBEndpoints{Config: c}
	isTrue := reflect.DeepEqual(want, got)
	if isTrue != true {
		fmt.Errorf("Want: %v\nGot: %v", want, got)
	}
}
