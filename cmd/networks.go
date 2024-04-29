package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"nemesis-cli/src/app"
)

// networksCmd represents the networks command
var networksCmd = &cobra.Command{
	Use:   "networks",
	Short: "Get networks infos",
	Long:  `Use this command to get networks infos from the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		name := cmd.Flag("name").Value.String()
		id := cmd.Flag("id").Value.String()

		if name != "" && id != "" {
			fmt.Println("Please provide only a name or an id")
			return
		}

		if all, _ := cmd.Flags().GetBool("all"); all {
			if err := app.Networks("", ""); err != nil {
				fmt.Println(err)
			}
			return
		}

		if name == "" && id == "" {
			fmt.Println("Please provide a name or an id")
			return
		}

		if id, _ := cmd.Flags().GetString("id"); id != "" {
			if err := app.Networks(id, ""); err != nil {
				fmt.Println(err)
			}
			return
		}

		if name, _ := cmd.Flags().GetString("name"); name != "" {
			if err := app.Networks("", name); err != nil {
				fmt.Println(err)
			}
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(networksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// networksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// networksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	networksCmd.Flags().BoolP("all", "a", false, "Get all Networks")
	networksCmd.Flags().StringP("id", "i", "", "Get network by id")
	networksCmd.Flags().StringP("name", "n", "", "Filter networks by name")

}
