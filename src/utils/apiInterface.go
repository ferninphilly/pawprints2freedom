package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func CallAPI(requestType string, url string, authToken string, integration string) ([]byte, error) {
	fmt.Println("Beginning API call for ", integration)
	req, err := http.NewRequest(requestType, url, nil)
	if err != nil {
		log.Fatalf("Error getting this integration: %s this url: %s due to %s", integration, url, err.Error())
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	if len(authToken) > 0 {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authToken))
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making call for integration: %s API due to %s", integration, err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error received for integrations %s from API: %s", integration, err.Error())
	}
	return bodyBytes, nil
}
