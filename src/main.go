package main

import (
	"log"
	"os"

	fb "github.com/ferninphilly/pawprints2freedom/facebook"
)

func main() {
	fbe := fb.NewFBEndpoints()
	if err := fbe.GetPageAccessToken(); err != nil {
		log.Fatalf("There was an error getting the facebook access token due to %s", err.Error())
		os.Exit(1)
	}
}
