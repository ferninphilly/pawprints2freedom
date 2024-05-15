package facebook

import util "github.com/ferninphilly/pawprints2freedom/utils"

type FBEndpoints struct {
	FacebookConfig  *util.FacebookConfig
	UserID          string
	PageAccessData  []PageAccessData `json:"data"`
	UserToken       string
	PermAccessToken string
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
type AccessTokenResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	TokenType   string `json:"token_type,omitempty"`
	ExpiresIn   int32  `json:"expires_in,omitempty"`
}
