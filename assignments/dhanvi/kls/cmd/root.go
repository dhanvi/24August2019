package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile, kubeConfig, namespace string

var rootCmd = &cobra.Command{
	Use:   "kls",
	Short: "list pods and deployment of k8s",
	Long:  `List the name and status of pods and deployment based on the config and namespace passed`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&kubeConfig, "kube-config", "k", os.ExpandEnv("$HOME/.kube/config"), "config path to user")
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "kube-system", "namespace to use")
	viper.BindPFlag("kube-config", rootCmd.PersistentFlags().Lookup("kube-config"))
	viper.BindPFlag("namespace", rootCmd.PersistentFlags().Lookup("namespace"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

}
