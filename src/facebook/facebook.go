package facebook

import (
	"encoding/json"
	"fmt"
	"log"

	util "github.com/ferninphilly/pawprints2freedom/utils"
)

func NewFBEndpoints() *FBEndpoints {
	var cfg util.FacebookConfig
	c := &cfg
	if err := c.PopulateConfig(); err != nil {
		log.Fatal("Error populating the facebook config due to ", err.Error())
	}
	return &FBEndpoints{FacebookConfig: c}
}

func (fb *FBEndpoints) GetPageAccessToken() error {
	url := fmt.Sprintf("%s/oauth/access_token?client_id=%s&client_secret=%s&grant_type=client_credentials", fb.FacebookConfig.BaseURL, fb.FacebookConfig.AppID, fb.FacebookConfig.AppSecret)
	bytesData, err := util.CallAPI("GET", url, "", "facebook")
	var atr AccessTokenResponse
	if err != nil {
		log.Fatalf("Error Calling the API for Facebook due to %s", err.Error())
		return err
	}
	fmt.Println("Here is our response from the fb api: ", string(bytesData))
	if err := json.Unmarshal(bytesData, &atr); err != nil {
		log.Fatalf("Error unmarshalling this %s into fb.PageAccessData due to %s", string(bytesData), err.Error())
		return err
	}
	fb.UserToken = fb.FacebookConfig.UserToken
	fb.PermAccessToken = atr.AccessToken
	return nil
}

func (fb *FBEndpoints) getCurrentUserID() error {
	url := fmt.Sprintf("%s/me?fields=id&access_token=%s", fb.FacebookConfig.BaseURL, fb.FacebookConfig.UserToken)
	var currentUserID UserID
	bytesData, err := util.CallAPI("GET", url, "", "facebook")
	if err != nil {
		fmt.Printf("Error Calling the API for Facebook due to %s", err.Error())
		return err
	}
	if err := json.Unmarshal(bytesData, &currentUserID); err != nil {
		log.Fatalf("Error unmarshalling this %s into currentUserID due to %s", string(bytesData), err.Error())
		return err
	}
	fb.UserID = currentUserID.ID
	return nil
}
