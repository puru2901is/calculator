/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// prodCmd represents the prod command
var prodCmd = &cobra.Command{
	Use:   "prod",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//jokeTer, _ := cmd.Flags().GetString("term")
		numbers, _ := cmd.Flags().GetFloat32Slice("number")
		//fmt.Println(number)
		var prod float64
		prod = 1

		for _, j := range numbers {
			//fmt.Println(i)
			prod *= float64(j)
		}
		fmt.Println(prod)
	},
}

func init() {
	rootCmd.AddCommand(prodCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// prodCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// prodCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
