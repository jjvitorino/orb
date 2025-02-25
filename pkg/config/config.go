/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type MFSDKConfig struct {
	BaseURL      string `mapstructure:"base_url"`
	ThingsPrefix string `mapstructure:"things_prefix"`
}

type GRPCConfig struct {
	Service    string
	URL        string `mapstructure:"url"`
	Port       string `mapstructure:"port"`
	Timeout    string `mapstructure:"timeout"`
	CaCerts    string `mapstructure:"ca_certs"`
	ClientTLS  string `mapstructure:"client_tls"`
	ServerCert string `mapstructure:"server_cert"`
	ServerKey  string `mapstructure:"server_key"`
}
type NatsConfig struct {
	URL             string `mapstructure:"url"`
	ConsumerCfgPath string `mapstructure:"config_path"`
}

type EsConfig struct {
	URL      string `mapstructure:"url"`
	Pass     string `mapstructure:"pass"`
	DB       string `mapstructure:"db"`
	Consumer string `mapstructure:"consumer"`
}

type JaegerConfig struct {
	URL string `mapstructure:"url"`
}

type BaseSvcConfig struct {
	LogLevel       string `mapstructure:"log_level"`
	HttpPort       string `mapstructure:"http_port"`
	HttpServerCert string `mapstructure:"server_cert"`
	HttpServerKey  string `mapstructure:"server_key"`
}

type PostgresConfig struct {
	Host        string `mapstructure:"host"`
	Port        string `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Pass        string `mapstructure:"pass"`
	DB          string `mapstructure:"db"`
	SSLMode     string `mapstructure:"ssl_mode"`
	SSLCert     string `mapstructure:"ssl_cert"`
	SSLKey      string `mapstructure:"ssl_key"`
	SSLRootCert string `mapstructure:"ssl_root_cert"`
}

func LoadMFSDKConfig(prefix string) MFSDKConfig {

	cfg := viper.New()
	cfg.SetEnvPrefix(fmt.Sprintf("%s_sdk", prefix))

	cfg.SetDefault("base_url", "http://localhost")
	cfg.SetDefault("things_prefix", "")

	cfg.AllowEmptyEnv(true)
	cfg.AutomaticEnv()
	var nC MFSDKConfig
	cfg.Unmarshal(&nC)

	return nC
}

func LoadNatsConfig(prefix string) NatsConfig {

	cfg := viper.New()
	cfg.SetEnvPrefix(fmt.Sprintf("%s_nats", prefix))

	cfg.SetDefault("url", "nats://localhost:4222")
	cfg.SetDefault("config_path", "/config.toml")

	cfg.AllowEmptyEnv(true)
	cfg.AutomaticEnv()
	var nC NatsConfig
	cfg.Unmarshal(&nC)

	return nC
}

func LoadPostgresConfig(prefix string, db string) PostgresConfig {

	cfg := viper.New()
	cfg.SetEnvPrefix(fmt.Sprintf("%s_db", prefix))

	cfg.SetDefault("host", "localhost")
	cfg.SetDefault("port", "5432")
	cfg.SetDefault("user", "orb")
	cfg.SetDefault("pass", "orb")
	cfg.SetDefault("db", db)
	cfg.SetDefault("ssl_mode", "verify-full")
	cfg.SetDefault("ssl_cert", "")
	cfg.SetDefault("ssl_key", "")
	cfg.SetDefault("ssl_root_cert", "")

	cfg.AutomaticEnv()
	cfg.AllowEmptyEnv(true)
	var jC PostgresConfig
	cfg.Unmarshal(&jC)

	return jC
}

func LoadJaegerConfig(prefix string) JaegerConfig {

	cfg := viper.New()
	cfg.SetEnvPrefix(fmt.Sprintf("%s_jeager", prefix))

	cfg.SetDefault("url", "localhost:6831")

	cfg.AllowEmptyEnv(true)
	cfg.AutomaticEnv()
	var jC JaegerConfig
	cfg.Unmarshal(&jC)

	return jC
}

func LoadEsConfig(prefix string) EsConfig {
	cfg := viper.New()
	cfg.SetEnvPrefix(fmt.Sprintf("%s_es", prefix))

	cfg.SetDefault("url", "localhost:6379")
	cfg.SetDefault("pass", "")
	cfg.SetDefault("db", "0")
	cfg.SetDefault("consumer", fmt.Sprintf("%s-es-consumer", prefix))

	cfg.AllowEmptyEnv(true)
	cfg.AutomaticEnv()
	var esC EsConfig
	cfg.Unmarshal(&esC)
	return esC
}

func LoadBaseServiceConfig(prefix string, httpPort string) BaseSvcConfig {
	cfg := viper.New()
	cfg.SetEnvPrefix(prefix)

	cfg.SetDefault("log_level", "error")
	cfg.SetDefault("http_port", httpPort)
	cfg.SetDefault("server_cert", "")
	cfg.SetDefault("server_key", "")

	cfg.AllowEmptyEnv(true)
	cfg.AutomaticEnv()
	var svcC BaseSvcConfig
	cfg.Unmarshal(&svcC)
	return svcC
}

func LoadGRPCConfig(prefix string, svc string) GRPCConfig {
	cfg := viper.New()
	cfg.SetEnvPrefix(fmt.Sprintf("%s_%s_grpc", prefix, svc))

	cfg.SetDefault("url", "localhost:8181")
	cfg.SetDefault("port", "")
	cfg.SetDefault("timeout", "1s")
	cfg.SetDefault("client_tls", "false")
	cfg.SetDefault("ca_certs", "")
	cfg.SetDefault("server_cert", "")
	cfg.SetDefault("server_key", "")

	cfg.AllowEmptyEnv(true)
	cfg.AutomaticEnv()
	var aC GRPCConfig
	aC.Service = svc
	cfg.Unmarshal(&aC)
	return aC
}
