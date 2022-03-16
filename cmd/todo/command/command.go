package command

import (
	"graphql-server/cmd/todo/command/server"

	"github.com/spf13/cobra"
)

type Command func() *cobra.Command

func NewCommand(
	serverCmd server.Command,
) Command {
	return func() *cobra.Command {
		cmd := &cobra.Command{
			Use:   "todo",
			Short: "todo cli",
		}
		cmd.AddCommand(serverCmd())
		return cmd
	}
}
