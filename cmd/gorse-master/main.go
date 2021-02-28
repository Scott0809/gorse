// Copyright 2020 gorse Project Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/zhenghaoz/gorse/cmd/version"
	"github.com/zhenghaoz/gorse/config"
	"github.com/zhenghaoz/gorse/master"
)

var masterCommand = &cobra.Command{
	Use:   "gorse-master",
	Short: "The master node of gorse recommender system.",
	Run: func(cmd *cobra.Command, args []string) {
		// Show version
		if showVersion, _ := cmd.PersistentFlags().GetBool("version"); showVersion {
			fmt.Println(version.VersionName)
			return
		}
		// Start master
		configPath, _ := cmd.PersistentFlags().GetString("config")
		log.Infof("master: load config from %v", configPath)
		conf, meta, err := config.LoadConfig(configPath)
		if err != nil {
			log.Fatal(err)
		}
		l := master.NewMaster(conf, meta)
		l.Serve()
	},
}

func init() {
	masterCommand.PersistentFlags().StringP("config", "c", "/etc/gorse.toml", "configuration file path")
	masterCommand.PersistentFlags().BoolP("version", "v", false, "gorse version")
	masterCommand.PersistentFlags().Int("port", 8086, "port of master node")
	masterCommand.PersistentFlags().String("host", "127.0.0.1", "host of master node")
}

func main() {
	if err := masterCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
