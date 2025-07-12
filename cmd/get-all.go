package cmd

import (
	"github.com/PandaX185/pass-man/consts"
	"github.com/PandaX185/pass-man/pkg"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var getAllCmd = &cobra.Command{
	Use:     "get-all <domain>",
	Aliases: []string{"ga"},
	Short:   "Retrieve all passwords for a specific domain",
	Long:    `This command allows you to retrieve all passwords associated with a specific domain and copy them to the clipboard.`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		domain := args[0]
		bolt := &pkg.BoltDB{}
		err := bolt.OpenBoltDB()
		if err != nil {
			cmd.Println(consts.RED("Error opening database:" + err.Error()))
			return
		}
		defer bolt.DB.Close()

		password, err := bolt.GetPasswordsByDomain(domain)
		if err != nil {
			cmd.Println(consts.RED("Error retrieving password:" + err.Error()))
			return
		}

		if err := clipboard.WriteAll(password); err != nil {
			cmd.Println(consts.RED("Error copying passwords to clipboard:" + err.Error()))
			return
		}
		cmd.Println(consts.GREEN("All passwords copied to clipboard successfully."))

	},
}

func init() {
	rootCmd.AddCommand(getAllCmd)
	getAllCmd.SetUsageTemplate("Usage: pass-man get-all <domain>\n\n" +
		"Retrieve all passwords for a specific domain.\n\n")
}
