package label

import (
	"github.com/spf13/cobra"
)

func AddCommands(command *cobra.Command) {
	command.AddCommand(LabelCommand)
	LabelCommand.AddCommand(NewAttachLabelCommand())
	LabelCommand.AddCommand(NewRemoveLabelCommand())
	LabelCommand.AddCommand(NewCreateLabelCommand())
	LabelCommand.AddCommand(NewDeleteLabelCommand())
	LabelCommand.AddCommand(NewListLabelCommand())
}

var LabelCommand = &cobra.Command{
	Use:   "label attach|remove|create|delete",
	Short: "Provides label commands",
	Long:  `Provides label commands`,
}
