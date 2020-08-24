package console

import (
	"log"

	"app/app"

	"github.com/spf13/cobra"
)

func http() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "http",
		Short: "Run server on http protocal",
		Run: func(cmd *cobra.Command, args []string) {
			if err := app.RunHTTP(); err != nil {
				log.Fatalln(err)
			}
		},
	}
	return cmd
}
