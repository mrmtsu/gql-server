package server

import (
	"log"
	"net/http"
	"os"

	"graphql-server/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/spf13/cobra"
)

type Command func() *cobra.Command

func NewCommand(
	resolverRoot generated.ResolverRoot,
) Command {
	const defaultPort = "8080"
	return func() *cobra.Command {
		cmd := &cobra.Command{
			Use:   "server",
			Short: "listen and serve api",
			RunE: func(cmd *cobra.Command, args []string) error {
				port := os.Getenv("PORT")
				if port == "" {
					port = defaultPort
				}

				srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolverRoot}))

				http.Handle("/", playground.Handler("GraphQL playground", "/query"))
				http.Handle("/query", srv)

				log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
				log.Fatal(http.ListenAndServe(":"+port, nil))
				return http.ListenAndServe(":"+port, nil)
			},
		}

		return cmd
	}
}
