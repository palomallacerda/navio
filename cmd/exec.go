package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/container"
	"github.com/viniciusbds/navio/utilities"
)

var (
	// Used for id flag.
	containerID string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&containerID, "id", "", "The ID of the container")
	rootCmd.AddCommand(exec())
}

func exec() *cobra.Command {
	return &cobra.Command{
		Use:   "exec",
		Short: "Run a command in a running container",
		RunE: func(cmd *cobra.Command, args []string) error {

			// navio exec [--id containerID | --name containerName] COMMAND PARAMS...

			if containerName == "" && containerID == "" {
				l.Log("WARNING", "You must insert a container name or a container id. (ex.: --name containerName or --id containerID)")
				return nil
			}

			if len(args) < 1 {
				l.Log("WARNING", "You must insert a command!")
				return nil
			}

			if utilities.IsEmpty(containerName) {
				containerName = container.GetContainerName(containerID)
			}

			if !container.RootfsExists(containerName) {
				l.Log("WARNING", fmt.Sprintf("%s is not a valid container. Run navio ps to see the available ones.", containerName))
				return nil
			}
			command := args[0]
			params := args[1:]

			l.Log("INFO", fmt.Sprintf("Container: %s, Command: %s, Params: %v", containerName, command, params))
			args = append([]string{containerName, command}, params...)
			container.Exec(args)

			return nil
		},
	}
}
