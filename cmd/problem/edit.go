package problem

import (
	"fmt"
	"github.com/djungyvr/notr/io"
	"github.com/djungyvr/notr/models"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"os"
)

func NewEditCommand() *cobra.Command {
	editCommand := &cobra.Command{
		Use:   "edit PROBLEM_ID",
		Short: "Edits a problem",
		Long:  `Edits a local problem referenced by ID`,
		Run:   editProblem,
		Args:  cobra.ExactArgs(1),
	}
	return editCommand
}

type YamlProblem struct {
	ProblemStatement  string `yaml:"problem,omitempty"`
	SolutionStatement string `yaml:"solution,omitempty"`
}

func editProblem(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		id := args[0]

		problem, err := models.Repo().GetProblem(id)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		yamlProblem := YamlProblem{ProblemStatement: problem.ProblemStatement, SolutionStatement: problem.SolutionStatement}

		marshaled, err := yaml.Marshal(yamlProblem)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		content := []byte(marshaled)

		data, err := io.CaptureVim(content)

		err = yaml.Unmarshal(data, &yamlProblem)
		if err != nil {
			panic(err)
		}

		problem.ProblemStatement = yamlProblem.ProblemStatement
		problem.SolutionStatement = yamlProblem.SolutionStatement

		if err := models.Repo().SaveProblem(problem); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
