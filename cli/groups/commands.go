package groups

import "github.com/spf13/cobra"

func LoadGroupCommands(rootCmd *cobra.Command) {
	groupCmd := &cobra.Command{
		Use:   "group",
		Short: "Manage groups",
	}
	rootCmd.AddCommand(groupCmd)

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add an group",
		Run:   AddGroup,
	}
	groupCmd.AddCommand(addCmd)

	removeCmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove an group",
		Run:   RemoveGroup,
	}

	groupCmd.AddCommand(removeCmd)

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List groups",
		Run:   ListGroups,
	}

	groupCmd.AddCommand(listCmd)
}
