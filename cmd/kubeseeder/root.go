/*
Copyright © 2024 Vishal Anarase <iamvishalanarase@gmail.com>
*/
package main

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile, kubeconfig string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kubeseeder",
	Short: "Kubeseeder is a tool for planting resource into a Kubernetes cluster",
	Long:  `Kubeseeder is a tool for planting resource into a Kubernetes cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "$HOME/.kubeseeder.yaml", "Path to config file")
	rootCmd.PersistentFlags().StringVarP(&kubeconfig, "kubeconfig", "k", "$HOME/.kube/config", "Path to kubeconfig file")
}
