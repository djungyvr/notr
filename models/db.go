package models

type DB interface {
	InitSchema() error
	SaveOpenedProblem(problem Problem) error
	SaveClosedProblem(problemId Problem) error
	ListProblems() ([]Problem, error)
}
