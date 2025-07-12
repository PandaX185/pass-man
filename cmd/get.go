package cmd

import (
	"github.com/PandaX185/pass-man/consts"
	"github.com/PandaX185/pass-man/pkg"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:     "get <email>",
	Aliases: []string{"g"},
	Short:   "Retrieve the password for a specific email",
	Long:    `This command allows you to retrieve the password associated with a specific email address.`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		email := args[0]
		bolt := &pkg.BoltDB{}
		err := bolt.OpenBoltDB()
		if err != nil {
			cmd.Println(consts.RED("Error opening database:" + err.Error()))
			return
		}
		defer bolt.DB.Close()

		password, err := bolt.GetPasswordByEmail(email)
		if err != nil {
			cmd.Println(consts.RED("Error retrieving password for email " + email + " : " + err.Error()))
			return
		}

		if err := clipboard.WriteAll(password); err != nil {
			cmd.Println(consts.RED("Error copying password to clipboard:" + err.Error()))
			return
		}
		cmd.Println(consts.GREEN("Password copied to clipboard successfully."))

	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.SetUsageTemplate("Usage: pass-man get <email>\n\n" +
		"Retrieve the password for a specific email address.\n\n")
}
