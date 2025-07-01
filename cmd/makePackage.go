/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli-boilerplate-generator/internal/domain/makepackages"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

// makePackageCmd represents the makePackage command
var makePackageCmd = &cobra.Command{
	Use:   "makePackage",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Wrong number of arguments")
			return
		}

		value, err := strconv.ParseBool(args[1])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		makePackage := makepackages.MakePackages{
			PackageName: args[0],
			UseUUID:     value,
		}

		makePackage.CreatePackage()
	},
}

func init() {
	rootCmd.AddCommand(makePackageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// makePackageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// makePackageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
