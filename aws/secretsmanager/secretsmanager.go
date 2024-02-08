package secretsmanager

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

var (
	ctx                   = context.TODO()
	cfg, _                = config.LoadDefaultConfig(ctx)
	SecretsManagerSession = secretsmanager.NewFromConfig(cfg)
)

func GetSecretValue(secretArn *string) (*secretsmanager.GetSecretValueOutput, error) {
	secretValue, err := SecretsManagerSession.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId: secretArn,
	})

	return secretValue, err
}

type DBCredentials struct {
	DBInstanceIdentifier string `json:"dbInstanceIdentifier"`
	DBName               string `json:"dbname"`
	Engine               string `json:"engine"`
	Host                 string `json:"host"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	Port                 int    `json:"port"`
}

func GetDBCredentials(secretArn *string) (DBCredentials, error) {
	// Get the db credentials value
	secretValue, err := GetSecretValue(secretArn)
	if err != nil {
		return DBCredentials{}, err
	}

	// Unmarshal the secret string into your struct
	var credentials DBCredentials
	json.Unmarshal([]byte(*secretValue.SecretString), &credentials)

	return credentials, nil
}
