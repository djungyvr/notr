package label

import (
	"fmt"
	"github.com/djungyvr/notr/models"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func NewAttachLabelCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "attach LABEL PROBLEM_ID",
		Short: "Attaches a label to a problem",
		Long:  `Attaches a label to a problem`,
		Args:  cobra.ExactArgs(2),
		Run:   attachLabel,
	}
	return command
}

func attachLabel(cmd *cobra.Command, args []string) {
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

	if len(labels) == 1 && labels[0] == "" {
		labels = []string{}
	}

	for _, l := range labels {
		if l == label {
			fmt.Printf("Label %s already attached to problem %s\n", label, id)
			os.Exit(1)
		}
	}

	labels = append(labels, label)

	problem.Labels = strings.Join(labels, ",")

	if err := models.Repo().SaveProblem(problem); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Label %s attached to problem %s\n", label, id)
}
