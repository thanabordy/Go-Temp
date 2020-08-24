package console

import (
	"log"

	"app/app/model"
	"app/lib/database"
	"github.com/spf13/cobra"
)

func migrateDown() *cobra.Command {
	cmd := &cobra.Command{
		Use: "down",
		Run: func(cmd *cobra.Command, args []string) {
			if err := model.Down(database.DB()); err != nil {
				log.Fatalln(err)
			}
		},
	}
	return cmd
}
