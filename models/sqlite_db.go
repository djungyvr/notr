package models

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

var (
	repo *SqlDb
)

func Repo() *SqlDb {
	return repo
}

type SqlDb struct {
	db *sql.DB
}

// Init connects to the sqlite databse
// And generates the required tables
func Init(path string) error {
	dbPath, err := createDbIfNotExist(path)
	if err != nil {
		fmt.Printf("Error opening file %q\n", dbPath)
		return err
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Printf("Error opening database at %q\n", dbPath)
		return err
	}

	if repo == nil {
		repo = &SqlDb{db}
	}

	if err = repo.initSchema(); err != nil {
		fmt.Println("Error initializing database")
		return err
	}

	return nil
}

func createDbIfNotExist(path string) (string, error) {
	dbPath := path + "/notr.db"

	// Check to make sure the database exists
	if _, err := os.Stat(dbPath); !os.IsNotExist(err) {
		return dbPath, nil
	}

	// Create the file
	if _, err := os.Create(dbPath); err != nil {
		return "", err
	}

	return dbPath, nil
}

// Close releases the connections
func (m *SqlDb) Close() {
	m.db.Close()
}

func (m *SqlDb) initSchema() error {
	query := `
	CREATE TABLE IF NOT EXISTS
	problem (
		id TEXT PRIMARY KEY,
		problem_statement TEXT,
		solution_statement TEXT,
		is_open TEXT,
		opened_at TEXT,
		closed_at TEXT,
		labels TEXT
	)
	`
	statement, err := m.db.Prepare(query)
	if err != nil {
		fmt.Println("Statement error")
		return err
	}

	if _, err := statement.Exec(); err != nil {
		fmt.Println("Execution error")
		return err
	}

	// Mainly so that labels remain unique
	query = `
	CREATE TABLE IF NOT EXISTS
	label (
		name TEXT PRIMARY KEY
	)
	`

	statement, err = m.db.Prepare(query)
	if err != nil {
		fmt.Println("Statement error")
		return err
	}

	if _, err := statement.Exec(); err != nil {
		fmt.Println("Execution error")
		return err
	}

	return nil
}

func (m *SqlDb) DeleteProblem(id string) error {
	id = id + "%"
	query := `
	DELETE FROM problem
	WHERE id LIKE ?
	`

	statement, err := m.db.Prepare(query)

	if err != nil {
		fmt.Println("Statement error")
		return err
	}

	if _, err = statement.Exec(id); err != nil {
		fmt.Println("Execution error")
		return err
	}

	return err
}

func (m *SqlDb) SaveProblem(problem *Problem) error {
	query := `
	INSERT OR REPLACE INTO 
	problem (id, problem_statement, solution_statement, is_open, opened_at, closed_at, labels)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	statement, err := m.db.Prepare(query)
	if err != nil {
		fmt.Println("Statement error")
		return err
	}

	_, err = statement.Exec(
		problem.Id,
		problem.ProblemStatement,
		problem.SolutionStatement,
		problem.IsOpen,
		problem.OpenedAt,
		problem.ClosedAt,
		problem.Labels,
	)

	if err != nil {
		fmt.Println("Execution error")
		return err
	}

	return err
}

func (m *SqlDb) ListProblems() ([]Problem, error) {
	query := `
	SELECT 
	id, problem_statement, solution_statement, is_open, opened_at, closed_at, labels
	FROM problem`

	rows, err := m.db.Query(query)

	if err != nil {
		return []Problem{}, err
	}

	var problem Problem
	var problems []Problem

	for rows.Next() {
		err = rows.Scan(
			&problem.Id,
			&problem.ProblemStatement,
			&problem.SolutionStatement,
			&problem.IsOpen,
			&problem.OpenedAt,
			&problem.ClosedAt,
			&problem.Labels,
		)

		if err != nil {
			return []Problem{}, err
		}
		problems = append(problems, problem)
	}

	rows.Close()

	return problems, nil
}

func (m *SqlDb) GetProblem(id string) (*Problem, error) {
	id = id + "%"
	query := `
	SELECT 
	id, problem_statement, solution_statement, is_open, opened_at, closed_at, labels
	FROM problem
	WHERE id LIKE ?`

	row := m.db.QueryRow(query, id)

	problem := &Problem{}

	err := row.Scan(
		&problem.Id,
		&problem.ProblemStatement,
		&problem.SolutionStatement,
		&problem.IsOpen,
		&problem.OpenedAt,
		&problem.ClosedAt,
		&problem.Labels,
	)

	if err != nil {
		fmt.Println("Scan error")
		return nil, err
	}

	return problem, nil
}

func (m *SqlDb) SaveLabel(label string) error {
	query := `
	INSERT OR REPLACE INTO 
	label (name)
	VALUES (?)
	`

	statement, err := m.db.Prepare(query)
	if err != nil {
		fmt.Println("Statement error")
		return err
	}

	_, err = statement.Exec(label)

	if err != nil {
		fmt.Println("Execution error")
		return err
	}

	return err
}

func (m *SqlDb) DeleteLabel(label string) error {
	query := `
	DELETE FROM label
	WHERE name = ?
	`

	statement, err := m.db.Prepare(query)

	if err != nil {
		fmt.Println("Statement error")
		return err
	}

	if _, err = statement.Exec(label); err != nil {
		fmt.Println("Execution error")
		return err
	}

	return err
}

func (m *SqlDb) IsLabelAvailable(label string) bool {
	query := `
	SELECT *
	FROM label
	WHERE name = ?`

	var row string

	err := m.db.QueryRow(query, label).Scan(&row)

	return err != sql.ErrNoRows
}

func (m *SqlDb) ListLabels() ([]string, error) {
	query := `
	SELECT name 
	FROM label`

	rows, err := m.db.Query(query)

	if err != nil {
		return []string{}, err
	}

	var label string
	var labels []string

	for rows.Next() {
		err = rows.Scan(
			&label,
		)

		if err != nil {
			return []string{}, err
		}
		labels = append(labels, label)
	}

	rows.Close()

	return labels, nil
}
