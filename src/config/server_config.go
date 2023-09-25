package config

import (
	"github.com/caarlos0/env/v9"
)

type ServerConfig struct {
	Environment          string              `yaml:"environment" env:"environment"`
	ServiceAdminBaseURL  string              `yaml:"service_admin_base_url" env:"service_admin_base_url"`
	ServiceFromEmail     string              `yaml:"service_from_email" env:"service_from_email"`
	ServiceToEmail       string              `yaml:"service_to_email" env:"service_to_email"`
	JwtIssuer            string              `yaml:"jwt_issuer" env:"jwt_issuer"`
	SentryDSN            string              `yaml:"sentry_dns" env:"sentry_dsn"`
	Slack                SlackConfig         `yaml:"slack" envPrefix:"slack_"`
	ElasticSearch        ElasticSearchConfig `yaml:"elastic_search" envPrefix:"elastic_search_"`
	ElasticSearchProject ElasticSearchConfig `yaml:"elastic_search_project"  envPrefix:"elastic_search_project_"`
	Cognito              CognitoConfig       `yaml:"cognito" envPrefix:"cognito_"`
	ExternalFrontEndURL  string              `yaml:"external_front_end_url" env:"external_front_end_url"`
	V2FrontEndURL        string              `yaml:"v2_front_end_url" env:"v2_front_end_url"`
}

type SlackConfig struct {
	WebhookURL string `yaml:"webhook_url" env:"webhook_url"`
}

type ElasticSearchConfig struct {
	Endpoint           string `yaml:"endpoint" env:"endpoint"`
	AWSAccessKeyId     string `yaml:"aws_access_key_id" env:"aws_access_key_id"`
	AWSSecretAccessKey string `yaml:"aws_secret_access_key" env:"aws_secret_access_key"`
}

type CognitoConfig struct {
	AdminClientId string `yaml:"admin_client_id" env:"admin_client_id"`
	UserClientId  string `yaml:"user_client_id" env:"user_client_id"`
}

func LoadServerConfigFromENV() (*ServerConfig, error) {
	var conf ServerConfig
	err := env.ParseWithOptions(&conf, env.Options{
		RequiredIfNoDef: true,
	})
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
