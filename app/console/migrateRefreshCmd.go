package console

import (
	"log"
	"app/lib/database"
	"app/app/model"
	"github.com/spf13/cobra"
)

  func migrateRefresh() *cobra.Command {
	cmd :=  &cobra.Command{
		Use:   "refresh",
		Run: func(cmd *cobra.Command, args []string) {
			if err := model.Down(database.DB()); err != nil { log.Fatalln(err) }
			if err := model.Up(database.DB()); err != nil { log.Fatalln(err) }
			// if err := model.Seed(database.DB()); err != nil { log.Fatalln(err) }
		},
	  }
	return cmd
  }