package console

import (
	"log"
	"app/lib/database"
	"app/app/model"
	"github.com/spf13/cobra"
)


  func migrateSeed() *cobra.Command {
	cmd :=  &cobra.Command{
		Use:   "seed",
		Run: func(cmd *cobra.Command, args []string) {
			if err := model.Seed(database.DB()); err != nil {
				log.Fatalln(err)
			}
		},
	  }
	return cmd
  }