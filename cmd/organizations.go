/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"nemesis-cli/src/app"

	"github.com/spf13/cobra"
)

// organizationsCmd represents the organizations command
var organizationsCmd = &cobra.Command{
	Use:   "organizations",
	Short: "Get organizations infos",
	Long:  `Use this command to get organizations infos from the database.`,
	Run: func(cmd *cobra.Command, args []string) {

		if all, _ := cmd.Flags().GetBool("all"); all {
			if err := app.Organizations("", ""); err != nil {
				fmt.Println(err)
			}
			return
		}

		if id, _ := cmd.Flags().GetString("id"); id != "" {
			if err := app.Organizations(id, ""); err != nil {
				fmt.Println(err)
			}
			return
		}

		if name, _ := cmd.Flags().GetString("name"); name != "" {
			if err := app.Organizations("", name); err != nil {
				fmt.Println(err)
			}
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(organizationsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// organizationsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// organizationsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	organizationsCmd.Flags().BoolP("all", "a", false, "Get all Organizations")
	organizationsCmd.Flags().StringP("id", "i", "", "Get organization by id")
	organizationsCmd.Flags().StringP("name", "n", "", "Filter organization by name")
}
