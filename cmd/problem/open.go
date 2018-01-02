package problem

import (
	"fmt"
	"github.com/djungyvr/notr/io"
	"github.com/djungyvr/notr/models"
	"github.com/spf13/cobra"
	"os"
)

func NewOpenCommand() *cobra.Command {
	openCommand := &cobra.Command{
		Use:   "open",
		Short: "Open a new problem",
		Long:  `Opens up a new problem, can be closed at a later time`,
		Run:   openProblem,
	}
	openCommand.Flags().StringVarP(&problemStatement, "message", "m", "", "problem statement")
	return openCommand
}

var problemStatement string

func openProblem(cmd *cobra.Command, args []string) {
	if problemStatement == "" {
		vimCapture, err := io.CaptureVim([]byte{})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		problemStatement = string(vimCapture)

		// remove trailing newline
		if len(problemStatement) > 0 {
			problemStatement = problemStatement[:len(problemStatement)-1]
		}
	}

	openedProblem := models.OpenProblem(problemStatement)

	if err := models.Repo().SaveProblem(openedProblem); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(openedProblem.Id)
}
