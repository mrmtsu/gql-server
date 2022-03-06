//go:generate go run github.com/99designs/gqlgen generate
package graph

import (
	"graphql-server/adapter/db"
	"graphql-server/graph/generated"
	"time"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	dbc  db.Connector
	nowf func() time.Time
}

func NewResolver(
	dbc db.Connector,
	nowf func() time.Time,
) generated.ResolverRoot {
	return &Resolver{
		dbc:  dbc,
		nowf: nowf,
	}
}
