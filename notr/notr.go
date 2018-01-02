package notr

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "notr",
	Short: "notr keeps track of your problems and their solutions",
	Long:  `notr is a lightweight solution tracker and aims to solve the age old question "How did I?"`,
}

func Execute() {
	RootCmd.Execute()
}
