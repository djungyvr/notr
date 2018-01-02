package problem

import (
	"github.com/spf13/cobra"
)

func AddCommands(command *cobra.Command) {
	command.AddCommand(NewCloseCommand())
	command.AddCommand(NewDeleteCommand())
	command.AddCommand(NewEditCommand())
	command.AddCommand(NewInspectCommand())
	command.AddCommand(NewListCommand())
	command.AddCommand(NewOpenCommand())
}
