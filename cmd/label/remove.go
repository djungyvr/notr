package label

import (
	"fmt"
	"github.com/djungyvr/notr/models"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func NewRemoveLabelCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "remove LABEL PROBLEM_ID",
		Short: "Removes a label from a problem",
		Long:  `Removes a label from a problem`,
		Args:  cobra.ExactArgs(2),
		Run:   removeLabel,
	}
	return command
}

func removeLabel(cmd *cobra.Command, args []string) {
	label := args[0]
	id := args[1]

	if !models.Repo().IsLabelAvailable(label) {
		fmt.Printf("Label %s unavailable, create before using\n", label)
		os.Exit(1)
	}

	problem, err := models.Repo().GetProblem(id)

	if err != nil {
		fmt.Printf("Problem %s unavailable\n", id)
		os.Exit(1)
	}

	labels := strings.Split(problem.Labels, ",")

	for i, l := range labels {
		if l == label {
			labels[i] = labels[len(labels)-1]

			problem.Labels = strings.Join(labels[:len(labels)-1], ",")

			if err := models.Repo().SaveProblem(problem); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Printf("Label %s removed from problem %s\n", label, id)
			return
		}
	}

	fmt.Printf("Label %s not attached to problem %s\n", label, id)
	os.Exit(1)
}
