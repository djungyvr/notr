package models

import (
	"github.com/satori/go.uuid"
	"time"
)

type Problem struct {
	Id                string `yaml:"id"`
	ProblemStatement  string `yaml:"problem"`
	SolutionStatement string `yaml:"solution"`
	IsOpen            bool   `yaml:"is_open"`
	OpenedAt          string `yaml:"opened_at"`
	ClosedAt          string `yaml:"closed_at"`
	Labels            string `yaml:"labels"`
}

func OpenProblem(problemStatement string) *Problem {
	id := uuid.NewV4().String()

	t := time.Now().UTC()
	timestamp := t.String()

	problem := &Problem{Id: id, ProblemStatement: problemStatement, SolutionStatement: "", OpenedAt: timestamp, IsOpen: true}

	return problem
}

func CloseProblem(problem *Problem, solutionStatement string) {
	t := time.Now().UTC()
	timestamp := t.String()

	problem.SolutionStatement = solutionStatement
	problem.IsOpen = false
	problem.ClosedAt = timestamp
}

func OpenCloseProblem(problemStatement string, solutionStatement string) *Problem {
	problem := OpenProblem(problemStatement)
	CloseProblem(problem, solutionStatement)
	return problem
}
