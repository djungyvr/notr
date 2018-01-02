package label

import (
	"fmt"
	"github.com/djungyvr/notr/models"
	"github.com/spf13/cobra"
	"os"
)

func NewCreateLabelCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "create LABEL",
		Short: "Creates a label",
		Long:  `Creates a label`,
		Args:  cobra.ExactArgs(1),
		Run:   createLabel,
	}
	return command
}

func createLabel(cmd *cobra.Command, args []string) {
	if err := models.Repo().SaveLabel(args[0]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
