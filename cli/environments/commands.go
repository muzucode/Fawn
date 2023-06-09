package environments

import "github.com/spf13/cobra"

func LoadEnvironmentCommands(rootCmd *cobra.Command) {
	envCmd := &cobra.Command{
		Use:   "env",
		Short: "Manage ",
	}
	rootCmd.AddCommand(envCmd)

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add an environment",
		Run:   AddEnvironment,
	}
	envCmd.AddCommand(addCmd)

	removeCmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove an environment",
		Run:   RemoveEnvironment,
	}

	envCmd.AddCommand(removeCmd)

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List ",
		Run:   ListEnvironments,
	}

	envCmd.AddCommand(listCmd)
}
