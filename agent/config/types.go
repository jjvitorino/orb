/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package config

type TLS struct {
	Verify bool `mapstructure:"verify"`
}

type APIConfig struct {
	Address string `mapstructure:"address"`
	Token   string `mapstructure:"token"`
}

type DBConfig struct {
	File string `mapstructure:"file"`
}

type MQTTConfig struct {
	Address   string `mapstructure:"address"`
	Id        string `mapstructure:"id"`
	Key       string `mapstructure:"key"`
	ChannelID string `mapstructure:"channel_id"`
}

type CloudConfig struct {
	AgentName     string `mapstructure:"agent_name"`
	AutoProvision bool   `mapstructure:"auto_provision"`
}

type Cloud struct {
	Config CloudConfig `mapstructure:"config"`
	API    APIConfig   `mapstructure:"api"`
	MQTT   MQTTConfig  `mapstructure:"mqtt"`
}

type OrbAgent struct {
	Backends map[string]map[string]string `mapstructure:"backends"`
	Tags     map[string]string            `mapstructure:"tags"`
	Cloud    Cloud                        `mapstructure:"cloud"`
	TLS      TLS                          `mapstructure:"tls"`
	DB       DBConfig                     `mapstructure:"db"`
}

type Config struct {
	Debug    bool
	Version  float64  `mapstructure:"version"`
	OrbAgent OrbAgent `mapstructure:"orb"`
}
