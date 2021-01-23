// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// apigwCmd command
var apigwCmd = &cobra.Command{
	Use:   "apigw",
	Short: "Interact with a Remote Helmet API Gateway",
}

// authCmd command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage Auth Methods",
}

// endpointsCmd command
var endpointsCmd = &cobra.Command{
	Use:   "endpoints",
	Short: "Manage Configured Endpoints on API Gateway",
}

// endpointsListCmd command
var endpointsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List API Gateway Endpoints",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("-->>")
	},
}

// authListCmd command
var authListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Auth Methods",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("-->>")
	},
}

// authCreateCmd command
var authCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an Auth Method",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("-->>")
	},
}

func init() {
	rootCmd.AddCommand(apigwCmd)

	// APIGW command sub commands
	apigwCmd.AddCommand(authCmd)
	apigwCmd.AddCommand(endpointsCmd)

	// Endpoints command sub commands
	endpointsCmd.AddCommand(endpointsListCmd)

	// Auth command sub commands
	authCmd.AddCommand(authListCmd)
	authCmd.AddCommand(authCreateCmd)
}
