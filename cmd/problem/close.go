package problem

import (
	"fmt"
	"github.com/djungyvr/notr/io"
	"github.com/djungyvr/notr/models"
	"github.com/spf13/cobra"
	"os"
)

func NewCloseCommand() *cobra.Command {
	closeCommand := &cobra.Command{
		Use:   "close PROBLEM_ID",
		Short: "Closes a problem",
		Long:  `Closes a local problem referenced by ID`,
		Run:   closeProblem,
		Args:  cobra.ExactArgs(1),
	}
	closeCommand.Flags().StringVarP(&solutionStatement, "message", "m", "", "solution statement")
	return closeCommand
}

var solutionStatement string

func closeProblem(cmd *cobra.Command, args []string) {
	id := args[0]

	closedProblem, err := models.Repo().GetProblem(id)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if solutionStatement == "" {
		vimCapture, err := io.CaptureVim([]byte{})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		solutionStatement = string(vimCapture)

		// remove trailing newline
		if len(solutionStatement) > 0 {
			solutionStatement = solutionStatement[:len(solutionStatement)-1]
		}
	}

	models.CloseProblem(closedProblem, solutionStatement)

	if err := models.Repo().SaveProblem(closedProblem); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
