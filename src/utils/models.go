package utils

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type FacebookConfig struct {
	BaseURL   string `json:"baseURL"`
	AppID     string `json:"appID"`
	AppSecret string `json:"appSecret"`
	UserToken string `json:"userToken"`
}

type RaiselyConfig struct {
	BaseURL             string `json:"baseURL"`
	ApiKey              string `json:"apiKey"`
	CampaignUUID        string `json:"campaignUUID"`
	CampaignProfileUUID string `json:"campaignProfileUuid"`
	GetDonationsURL     string `json:"getDonationsURL"`
}

func populateConfig(secretName string, region string) (*secretsmanager.GetSecretValueOutput, error) {
	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region), config.WithSharedConfigProfile("pawprints2freedom"))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	svc := secretsmanager.NewFromConfig(config)
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}
	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return result, nil
}

func (f *FacebookConfig) PopulateConfig() error {
	res, err := populateConfig("pawprints2freedom/facebook", "eu-west-1")
	if err != nil {
		log.Fatal("Error Populating the facebook config due to ", err)
		return err
	}
	if err := json.Unmarshal([]byte(*res.SecretString), f); err != nil {
		log.Fatal("Error Unmarshaling res into JSON ", err)
		return err
	}
	return nil
}

func (r *RaiselyConfig) PopulateConfig() error {
	res, err := populateConfig("pawprints2freedom/raisely", "eu-west-1")
	if err != nil {
		log.Fatal("Error Populating the facebook config due to ", err)
		return err
	}
	if err := json.Unmarshal([]byte(*res.SecretString), r); err != nil {
		log.Fatal("Error Unmarshaling res into JSON ", err)
		return err
	}
	return nil
}
