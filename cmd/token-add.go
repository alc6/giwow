package cmd

import (
	"github.com/manifoldco/promptui"
	"github.com/nicolasdscp/giwow/internal/netrc"
	"github.com/nicolasdscp/giwow/logger"
	"github.com/nicolasdscp/giwow/terminal"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var tokenAddCmd = &cobra.Command{
	Use:   "add [machine]",
	Short: "Add a new machine to your .netrc file",
	Args:  cobra.ExactArgs(1),
	Long: `[machine] is the name of the machine to add to your .netrc file.

Usually, it is the name of the git host you want to access eg: private.gitlab.com 
or github.com/my-private.
This will basically add a new line in your $HOME/.netrc file. 

You can avoid the interactive terminal by using the --user (-u) and --password (-p) flags.
Note that it's highly recommended to generate a personal access token on your git host
instead of using your password.

If a similar entry already exists in your .netrc file, it will be overwritten.`,
	Run: runTokenAdd,
}

func init() {
	tokenCmd.AddCommand(tokenAddCmd)

	tokenAddCmd.Flags().StringP("user", "u", "", "Directly set the user")
	tokenAddCmd.Flags().StringP("password", "p", "", "Directly set the password")
}

func runTokenAdd(cmd *cobra.Command, args []string) {
	user := cmd.Flag("user").Value.String()
	password := cmd.Flag("password").Value.String()

	if user == "" {
		prompt := promptui.Prompt{Label: "> User", Validate: terminal.NotEmpty("User")}
		result, err := prompt.Run()
		cobra.CheckErr(err)
		user = result
	}

	if password == "" {
		prompt := promptui.Prompt{Label: "> Password", Mask: '*', Validate: terminal.NotEmpty("Password")}
		result, err := prompt.Run()
		cobra.CheckErr(err)
		password = result
	}

	netrc.Current.AddMachine(args[0], user, password)
	cobra.CheckErr(netrc.Current.Save())

	logger.Print("New entry added to your .netrc file")
}
