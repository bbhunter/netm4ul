// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

package cmd

import (
	"fmt"
	"os"

	"github.com/netm4ul/netm4ul/core/api"
	"github.com/netm4ul/netm4ul/core/client"
	"github.com/netm4ul/netm4ul/core/config"
	"github.com/netm4ul/netm4ul/core/server"
	"github.com/netm4ul/netm4ul/core/session"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the requested service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("To few arguments !")
		cmd.Help()
		os.Exit(1)
	},
}

// startServerCmd represents the startServer command
var startServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {

		config.Config.IsServer = isServer
		config.Config.Nodes = make(map[string]config.Node)

		ss := session.NewSession(config.Config)

		// listen on all interface + Server port
		go server.CreateServer(ss)

		//TODO flag enable / disable api
		sa := session.NewSession(config.Config)
		go api.CreateAPI(sa)

		gracefulShutdown()

	},
}

// startClientCmd represents the startServer command
var startClientCmd = &cobra.Command{
	Use:   "client",
	Short: "Start the client",
	Run: func(cmd *cobra.Command, args []string) {

		config.Config.IsClient = isClient

		sc := session.NewSession(config.Config)

		go client.CreateClient(sc)

		gracefulShutdown()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.AddCommand(startServerCmd)
	startCmd.AddCommand(startClientCmd)
}
