package cmd

import (
	"os"

	"github.com/PandaX185/pass-man/consts"
	"github.com/PandaX185/pass-man/pkg"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var getAllCmd = &cobra.Command{
	Use:     "get-all <domain> [--json]",
	Aliases: []string{"ga"},
	Short:   "Retrieve all passwords for a specific domain",
	Long:    `This command allows you to retrieve all passwords associated with a specific domain and copy them to the clipboard or output them in JSON format.`,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		domain := args[0]
		bolt := &pkg.BoltDB{}
		err := bolt.OpenBoltDB()
		if err != nil {
			cmd.Println(consts.RED("Error opening database:" + err.Error()))
			return
		}
		defer bolt.DB.Close()

		passwords, err := bolt.GetPasswordsByDomain(domain)
		if err != nil {
			cmd.Println(consts.RED("Error retrieving passwords:" + err.Error()))
			return
		}

		if cmd.Flag("json").Changed {
			jsonOutput, err := pkg.ConvertToJSON(passwords)
			if err != nil {
				cmd.Println(consts.RED("Error converting passwords to JSON:" + err.Error()))
				return
			}
			file, err := os.Create("passwords.json")
			if err != nil {
				cmd.Println(consts.RED("Error creating JSON file:" + err.Error()))
				return
			}
			defer file.Close()

			_, err = file.WriteString(jsonOutput)
			if err != nil {
				cmd.Println(consts.RED("Error writing to JSON file:" + err.Error()))
				return
			}
			cmd.Println(consts.GREEN("Passwords written to passwords.json successfully."))
		} else {
			if err := clipboard.WriteAll(passwords); err != nil {
				cmd.Println(consts.RED("Error copying passwords to clipboard:" + err.Error()))
				return
			}
			cmd.Println(consts.GREEN("All passwords copied to clipboard successfully."))
		}
	},
}

func init() {
	rootCmd.AddCommand(getAllCmd)
	getAllCmd.SetUsageTemplate("Usage: pass-man get-all <domain>\n\n" +
		"Retrieve all passwords for a specific domain.\n\n")
	getAllCmd.Flags().BoolP("json", "j", false, "Output passwords in JSON format")
}
