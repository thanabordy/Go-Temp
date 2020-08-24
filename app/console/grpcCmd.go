package console

import (
	"app/app"
	"fmt"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

func gRPC() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grpc",
		Args:  notReqArgs,
		Short: "Run server on grpc protocal",
		Run: func(cmd *cobra.Command, args []string) {
			grpc := app.GRPC()
			c := make(chan os.Signal, 1)
			signal.Notify(c)
			<-c
			grpc.GracefulStop()
		},
	}
	return cmd
}

func notReqArgs(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		return fmt.Errorf("Not required areuments")
	}
	return nil
}
