package label

import (
	"fmt"
	"github.com/djungyvr/notr/models"
	"github.com/spf13/cobra"
	"os"
)

func NewDeleteLabelCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "rm LABEL",
		Short: "Deletes a label",
		Long:  `Deletes a label`,
		Args:  cobra.ExactArgs(1),
		Run:   deleteLabel,
	}
	return command
}

func deleteLabel(cmd *cobra.Command, args []string) {
	label := args[0]
	fmt.Printf("Deleting %s\n", label)

	if err := models.Repo().DeleteLabel(label); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
