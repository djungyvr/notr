package label

import (
	"fmt"
	"github.com/djungyvr/notr/models"
	"github.com/spf13/cobra"
	"os"
)

func NewListLabelCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "ls",
		Short: "Lists available labels",
		Long:  `Lists available labels`,
		Run:   listLabels,
		Args:  cobra.NoArgs,
	}
	return command
}

func listLabels(cmd *cobra.Command, args []string) {
	labels, err := models.Repo().ListLabels()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, label := range labels {
		fmt.Println(label)
	}
}
