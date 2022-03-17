//go:build wireinject
// +build wireinject

package di

import (
	"graphql-server/adapter"
	"graphql-server/adapter/db"
	"graphql-server/cmd/todo/command"
	"graphql-server/cmd/todo/command/server"
	"graphql-server/graph"
	"os"

	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	command.NewCommand,
	server.NewCommand,
	db.NewConnector,
	graph.NewResolver,
	adapter.NewSystemNowFunc,
	newDbDsn,
)

func ResolveCommand() command.Command {
	wire.Build(providerSet)
	return nil
}

func newDbDsn() db.Dsn {
	return db.Dsn(os.Getenv("DB_DSN"))
}
