// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package config contains convenience functions for reading and managing viper configs.
package config

import (
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
)

var (
	// Logrus structured logging setup
	cfgLog = logrus.WithFields(logrus.Fields{
		"app":       "openmatch",
		"component": "config",
	})

	// OpenCensus
	cfgVarCount = stats.Int64("config/vars_total", "Number of config vars read during initialization", "1")
	// CfgVarCountView is the Open Census view for the cfgVarCount measure.
	CfgVarCountView = &view.View{
		Name:        "config/vars_total",
		Measure:     cfgVarCount,
		Description: "The number of config vars read during initialization",
		Aggregation: view.Count(),
	}
)

// View is a read-only view of the Open Match configuration.
// New accessors from Viper should be added here.
type View interface {
	IsSet(string) bool
	GetString(string) string
	GetInt(string) int
	GetInt64(string) int64
	GetStringSlice(string) []string
	GetBool(string) bool
	GetDuration(string) time.Duration
	GetStringMap(string) map[string]interface{}
}

// Read reads a config file into a viper.Viper instance and associates environment vars defined in
// config.envMappings
func Read() (View, error) {
	cfg := viper.New()
	// Viper config management initialization
	// Support either json or yaml file types (json for backwards compatibility
	// with previous versions)
	cfg.SetConfigType("json")
	cfg.SetConfigType("yaml")
	cfg.SetConfigName("matchmaker_config")
	cfg.AddConfigPath(".")
	cfg.AddConfigPath("config")

	// Read in config file using Viper
	err := cfg.ReadInConfig()
	if err != nil {
		cfgLog.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatal("Fatal error reading config file")
	}

	// Look for updates to the config; in Kubernetes, this is implemented using
	// a ConfigMap that is written to the matchmaker_config.yaml file, which is
	// what the Open Match components using Viper monitor for changes.
	// More details about Open Match's use of Kubernetes ConfigMaps at:
	// https://open-match.dev/open-match/issues/42
	cfg.WatchConfig() // Watch and re-read config file.
	// Write a log when the configuration changes.
	cfg.OnConfigChange(func(event fsnotify.Event) {
		cfgLog.WithFields(logrus.Fields{
			"filename":  event.Name,
			"operation": event.Op,
		}).Info("Server configuration changed.")
	})
	return cfg, err
}
