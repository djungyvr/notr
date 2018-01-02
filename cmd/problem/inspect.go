package problem

import (
	"bytes"
	"fmt"
	"github.com/djungyvr/notr/models"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"os"
)

func NewInspectCommand() *cobra.Command {
	inspectCommand := &cobra.Command{
		Use:   "inspect PROBLEM_ID",
		Short: "Shows details of a problem",
		Long:  `Outputs all the data with regards to a problem`,
		Run:   inspectProblem,
	}
	return inspectCommand
}

func inspectProblem(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		id := args[0]

		problem, err := models.Repo().GetProblem(id)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var buffer bytes.Buffer

		body, err := yaml.Marshal(problem)

		buffer.Write(body)

		fmt.Println(buffer.String())

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
