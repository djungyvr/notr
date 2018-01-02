package problem

import (
	"fmt"
	"github.com/djungyvr/notr/models"
	"github.com/spf13/cobra"
	"os"
)

func NewDeleteCommand() *cobra.Command {
	deleteCommand := &cobra.Command{
		Use:   "rm",
		Short: "Delete a issue",
		Long:  `Deletes a local issue referenced by ID`,
		Run:   deleteProblem,
		Args:  cobra.ExactArgs(1),
	}
	return deleteCommand
}

func deleteProblem(cmd *cobra.Command, args []string) {
	id := args[0]
	fmt.Printf("Deleting %s\n", id)

	if err := models.Repo().DeleteProblem(id); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
