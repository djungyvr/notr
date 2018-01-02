package problem

import (
	"fmt"
	"github.com/djungyvr/notr/math"
	"github.com/djungyvr/notr/models"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var searchText string
var searchStatus string
var bePretty bool

func NewListCommand() *cobra.Command {
	listCommand := &cobra.Command{
		Use:   "ls",
		Short: "List problems",
		Long:  `List all local problems`,
		Run:   list,
	}
	listCommand.Flags().StringVarP(&searchText, "search", "f", "", "search text")
	listCommand.Flags().StringVarP(&searchStatus, "status", "s", "", "problem status")
	listCommand.Flags().BoolVarP(&bePretty, "pretty", "p", false, "pretty output with color")
	return listCommand
}

func list(cmd *cobra.Command, args []string) {
	problems, err := models.Repo().ListProblems()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	format := `
ID %s | Is Open %v | Opened At %s | Closed At %s | Labels %s
Problem  : %s
Solution : %s
`

	if strings.ToLower(searchStatus) == "open" || strings.ToLower(searchStatus) == "closed" {
		problems = filterProblemsByStatus(problems, strings.ToLower(searchStatus))
	}

	if searchText != "" {
		problems = filterProblemsByText(problems, searchText)
	}

	if bePretty {
		hdr := color.YellowString("ID: %s | Is Open %v | Opened At %s | Closed At %s | Labels ") + "%s\n"
		problemFmt := color.BlueString("Problem  : ") + "%s\n"
		solutionFmt := color.GreenString("Solution : ") + "%s\n"
		format = `
%s
%s
%s
`
		format = fmt.Sprintf(format, hdr, problemFmt, solutionFmt)
	}

	for _, p := range problems {
		openedAtIndex := math.MinInt(10, len(p.OpenedAt))
		closedAtIndex := math.MinInt(10, len(p.ClosedAt))
		problemIndex := math.MinInt(64, len(p.ProblemStatement))
		solutionIndex := math.MinInt(64, len(p.SolutionStatement))
		fmt.Printf(format, p.Id[:8], p.IsOpen, p.OpenedAt[:openedAtIndex], p.ClosedAt[:closedAtIndex], p.Labels, p.ProblemStatement[:problemIndex], p.SolutionStatement[:solutionIndex])
	}
}

func filterProblemsByStatus(problems []models.Problem, status string) []models.Problem {
	filteredProblems := []models.Problem{}

	statusFlag := status == "open"

	for _, p := range problems {
		if p.IsOpen == statusFlag {
			filteredProblems = append(filteredProblems, p)
		}
	}

	return filteredProblems
}

func filterProblemsByText(problems []models.Problem, text string) []models.Problem {
	filteredProblems := []models.Problem{}

	for _, p := range problems {
		if problemContains(p, text) {
			filteredProblems = append(filteredProblems, p)
		}
	}

	return filteredProblems
}

func problemContains(problem models.Problem, text string) bool {
	text = strings.ToLower(text)
	problemText := strings.ToLower(problem.ProblemStatement)
	solutionText := strings.ToLower(problem.SolutionStatement)

	return strings.Contains(problemText, text) || strings.Contains(solutionText, text)
}
