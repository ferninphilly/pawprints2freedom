package facebook

import (
	util "github.com/ferninphilly/pawprints2freedom/utils"
)

type PageAccessToken struct {
	Data []PageAccessData `json:"data"`
}

type PageAccessData struct {
	AccessToken  string         `json:"access_token"`
	Category     string         `json:"category"`
	CategoryList []CategoryList `json:"category_list"`
	Name         string         `json:"name"`
	ID           string         `json:"id"`
	Tasks        []string       `json:"tasks"`
}

type CategoryList struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserID struct {
	ID string `json:"id"`
}

type FBEndpoints struct {
	Config *util.Config
	UserID string
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}
