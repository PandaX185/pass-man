package cmd

import (
	"github.com/PandaX185/pass-man/consts"
	"github.com/PandaX185/pass-man/pkg"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add <email> <password>",
	Aliases: []string{"a"},
	Short:   "Add/Change a new password for certail email",
	Long:    `This command allows you to add or change a password for a specific email address.`,
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		email := args[0]
		password := args[1]
		bolt := &pkg.BoltDB{}
		err := bolt.OpenBoltDB()
		if err != nil {
			cmd.Println(consts.RED("Error opening database:" + err.Error()))
			return
		}
		defer bolt.DB.Close()

		err = bolt.AddPassword(email, password)
		if err != nil {
			cmd.Println(consts.RED("Error adding password:" + err.Error()))
			return
		}
		cmd.Println(consts.GREEN("Password added successfully for" + email))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.SetUsageTemplate("Usage: pass-man add <email> <password>\n\n" +
		"Add or change a password for a specific email address.\n\n")
}
