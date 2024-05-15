package raisely

import (
	"fmt"

	util "github.com/ferninphilly/pawprints2freedom/utils"
)

func (rs *Raisely) GetDonations(cfg *util.RaiselyConfig, sinceDate string) error {
	url := fmt.Sprintf("%s/%s?createdAtFrom=%s", cfg.BaseURL, cfg.GetDonationsURL, sinceDate)
	fmt.Println("Here is the URL we're hitting for raisely ", url)
	if err := util.CallAPI("GET", url, cfg.ApiKey, &rs.Donations); err != nil {
		fmt.Println("Error getting the Donations from raisely due to ", err.Error())
		return err
	}
	return nil
}
