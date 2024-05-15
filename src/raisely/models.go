package raisely

type Raisely struct {
	Donations RaiselyDonationReturn
}

type RaiselyDonationReturn struct {
	Data []RaiselyDonation `json:"data"`
}

type RaiselyDonation struct {
	UUID                 string  `json:"uuid"`
	PreferredName        string  `json:"preferredName"`
	Type                 string  `json:"type"`
	Method               string  `json:"method"`
	Currency             string  `json:"currency"`
	Message              string  `json:"message"`
	Anonymous            bool    `json:"anonymous"`
	Amount               int64   `json:"amount"`
	Date                 string  `json:"date"`
	UpdatedAt            string  `json:"updatedAt"`
	Mode                 string  `json:"mode"`
	Public               Public  `json:"public"`
	Status               string  `json:"status"`
	Processing           string  `json:"processing"`
	GatewayVersion       string  `json:"gatewayVersion"`
	CreatedAt            string  `json:"createdAt"`
	IncludesGiftAid      bool    `json:"includesGiftAid"`
	UsedSwiftAid         bool    `json:"usedSwiftAid"`
	DisplayTotal         int64   `json:"displayTotal"`
	CampaignDisplayTotal int64   `json:"campaignDisplayTotal"`
	DonationTotal        int64   `json:"donationTotal"`
	ProfileUUID          string  `json:"profileUuid"`
	CampaignUUID         string  `json:"campaignUuid"`
	SubscriptionUUID     string  `json:"subscriptionUuid"`
	FirstName            string  `json:"firstName"`
	CampaignAmount       int64   `json:"campaignAmount"`
	PublicAmount         int64   `json:"publicAmount"`
	PublicFee            float64 `json:"publicFee"`
	PublicTotal          float64 `json:"publicTotal"`
	CurrencySymbol       string  `json:"currencySymbol"`
	User                 User    `json:"user"`
	Profile              Profile `json:"profile"`
}

type Profile struct {
	UUID                     string      `json:"uuid"`
	Path                     string      `json:"path"`
	Name                     string      `json:"name"`
	Pronouns                 string      `json:"pronouns"`
	Currency                 string      `json:"currency"`
	Goal                     int64       `json:"goal"`
	Status                   string      `json:"status"`
	Type                     string      `json:"type"`
	Description              string      `json:"description"`
	CreatedAt                string      `json:"createdAt"`
	UpdatedAt                string      `json:"updatedAt"`
	Public                   string      `json:"public"`
	Internal                 interface{} `json:"internal"`
	Paid                     bool        `json:"paid"`
	ExerciseGoal             string      `json:"exerciseGoal"`
	ExerciseGoalTime         string      `json:"exerciseGoalTime"`
	IsCampaignProfile        bool        `json:"isCampaignProfile"`
	FacebookFundraiserID     string      `json:"facebookFundraiserId"`
	FundraiserTheme          string      `json:"fundraiserTheme"`
	PhotoURL                 string      `json:"photoUrl"`
	IndividualProfileCount   int64       `json:"individualProfileCount"`
	TeamProfileCount         int64       `json:"teamProfileCount"`
	OrganisationProfileCount int64       `json:"organisationProfileCount"`
	Items                    []Item      `json:"items"`
}

type Item struct {
	UUID            string `json:"uuid"`
	Description     string `json:"description"`
	IncludeInTotals bool   `json:"includeInTotals"`
	Amount          int64  `json:"amount"`
	CampaignAmount  int64  `json:"campaignAmount"`
	Type            string `json:"type"`
	Public          string `json:"public"`
	Quantity        int64  `json:"quantity"`
}

type User struct {
	UUID          string `json:"uuid"`
	PreferredName string `json:"preferredName"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
	Public        string `json:"public"`
	Status        string `json:"status"`
	FirstName     string `json:"firstName"`
	PhotoURL      string `json:"photoUrl"`
	Permission    string `json:"permission"`
}

type Public struct {
	FixedDescription string `json:"fixed_description"`
	FixedAmount      string `json:"fixed_amount"`
}
