package cmd

import (
	"github.com/djungyvr/notr/cmd/label"
	"github.com/djungyvr/notr/cmd/problem"
	"github.com/spf13/cobra"
)

func AddCommands(command *cobra.Command) {
	label.AddCommands(command)
	problem.AddCommands(command)
}
