package console

import (
	"github.com/spf13/cobra"
)

// Execute cmd
func Execute() error {
	cmd := &cobra.Command{
		Use:   "app",
		Short: "Hugo is a very fast static site generator",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
	cmd.AddCommand(http())
	cmd.AddCommand(gRPC())
	cmd.AddCommand(migrate())
	cmd.AddCommand(keygen())
	return cmd.Execute()
}
