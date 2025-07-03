/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	delimiter_constants "cli-boilerplate-generator/constants/delimiter-constants"
	filepathconstants "cli-boilerplate-generator/constants/file_path_constants"
	middleware2 "cli-boilerplate-generator/constants/middleware"
	"cli-boilerplate-generator/internal/domain/makeenums"
	"cli-boilerplate-generator/pkg/middleware"
	"cli-boilerplate-generator/pkg/utils/file_utils"
	"cli-boilerplate-generator/pkg/utils/file_writer"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var enumName string
var langName string
var projectName string

// makeEnumsCmd represents the makeEnums command
var makeEnumsCmd = &cobra.Command{
	Use:   "makeEnums",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		langName = strings.TrimSpace(langName)
		projectName = strings.TrimSpace(projectName)
		enumName = strings.TrimSpace(enumName)

		if enumName == "" || langName == "" {
			return fmt.Errorf("please provide both enum name and language name")
		}

		if projectName != "" {

			projectNameToStore := strings.ToUpper(langName) + delimiter_constants.CommaDelimiter + projectName
			err := file_writer.WriteStringInFile(projectNameToStore, filepathconstants.ProjectNameStorageFileAbsolutePath)
			if err != nil {
				return err
			}
		} else {
			allProjectNames, err := file_utils.GetAllFromFileAsSlices(filepathconstants.ProjectNameStorageFileAbsolutePath)

			if err != nil {
				return err
			}

			for _, projects := range allProjectNames {
				splitString := strings.Split(projects, delimiter_constants.CommaDelimiter)

				if len(splitString) != 2 {
					return fmt.Errorf("invalid project name being stored in database")
				}

				if splitString[0] == strings.ToUpper(langName) {
					projectName = splitString[1]
					break
				}
			}

			if projectName == "" {
				return fmt.Errorf("No project name found")
			}

		}

		return nil
	},
	Run: middleware.CommonMiddleware(func(cmd *cobra.Command, args []string) {
		cwd := cmd.Context().Value(middleware2.TERMINAL_PATH_GETTER).(string)

		makeEnums := makeenums.MakeEnums{
			EnumName:    enumName,
			Lang:        langName,
			CurrentPath: cwd,
			ProjectName: projectName,
		}
		makeEnums.CreateEnums()

	},
	),
}

func init() {
	rootCmd.AddCommand(makeEnumsCmd)

	makeEnumsCmd.Flags().StringVarP(&enumName, "name", "n", "", "Name of the enum")
	makeEnumsCmd.Flags().StringVarP(&langName, "lang", "l", "", "Language you're using")
	makeEnumsCmd.Flags().StringVarP(&projectName, "pname", "p", "", "Name of project")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// makeEnumsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// makeEnumsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
