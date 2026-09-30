// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cli

import (
	"fmt"

	"github.com/norwik/helmet/core"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start helmet server",
	Run: func(cmd *cobra.Command, args []string) {
		err := core.Load(config)
		if err != nil {
			panic(err.Error())
		}

		_, err = core.SetupLogging()
		if err != nil {
			panic(err.Error())
		}

		handler, context := core.SetupServer()

		err = core.RunServer(handler, context)
		if err != nil {
			panic(fmt.Sprintf("Server error: %s", err.Error()))
		}
	},
}

func init() {
	serverCmd.Flags().StringVarP(
		&config,
		"config",
		"c",
		"config.prod.yml",
		"Absolute path to config file (required)",
	)
	serverCmd.MarkFlagRequired("config")
	rootCmd.AddCommand(serverCmd)
}
