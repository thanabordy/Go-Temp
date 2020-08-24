package console

import (
	"log"

	model "app/app/model"
	"app/lib/database"
	"github.com/spf13/cobra"
)

func migrate() *cobra.Command {
	cmd := &cobra.Command{
		Use: "migrate",
		Run: func(cmd *cobra.Command, args []string) {
			if err := model.Up(database.DB()); err != nil {
				log.Fatalln(err)
			}
		},
	}
	cmd.AddCommand(migrateDown())
	cmd.AddCommand(migrateSeed())
	cmd.AddCommand(migrateRefresh())
	return cmd
}
