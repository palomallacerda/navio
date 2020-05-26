package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/container"
)

func init() {
	rootCmd.AddCommand(ps())
}

func ps() *cobra.Command {
	return &cobra.Command{
		Use:   "ps",
		Short: "List containers",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("ID\t\tNAME\t\t\tIMAGE\t\t\tCOMMAND\t\t\tSTATUS")
			containers, _ := container.Ps()
			fmt.Println(containers)
		},
	}
}
