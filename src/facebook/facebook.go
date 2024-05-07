package facebook

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	util "github.com/ferninphilly/pawprints2freedom/utils"
)

func NewFBEndpoints() *FBEndpoints {
	var cfg util.Config
	c := &cfg
	if err := util.GetConfigData(c); err != nil {
		fmt.Println("There was an error getting the fb config data due to ", err.Error())
	}
	return &FBEndpoints{Config: c}
}

func (fb *FBEndpoints) GetPageAccessToken() interface{} {
	fb.getCurrentUserID()
	if &fb.UserID == nil {
		fmt.Println("There is no UserID. Please run GetCurrentUserID() and try again!")
		os.Exit(1)
	}
	url := fmt.Sprintf("%s/%s/accounts?access_token=%s", fb.Config.Facebook.BaseURL, fb.UserID, fb.Config.Facebook.UserToken)
	var accessToken interface{}
	fbRequest("GET", url, &accessToken)
	return accessToken
}

func (fb *FBEndpoints) getCurrentUserID() {
	url := fmt.Sprintf("%s/me?fields=id&access_token=%s", fb.Config.Facebook.BaseURL, fb.Config.Facebook.UserToken)
	var currentUserID UserID
	fbRequest("GET", url, &currentUserID)
	fb.UserID = currentUserID.ID
}

func fbRequest(requestType string, url string, responseStruct interface{}) error {
	req, err := http.NewRequest(requestType, url, nil)
	if err != nil {
		fmt.Printf("Error getting this url: %s due to %s", url, err.Error())
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making call to facebook due to ", err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error received from FB: ", err.Error())
	}
	fmt.Println("Here is the return from FB ", string(bodyBytes))
	json.Unmarshal(bodyBytes, &responseStruct)
	return nil
}
