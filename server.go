package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tomato3713/nulledge/database"
	"github.com/tomato3713/nulledge/graph"
)

const defaultPort = "8080"

func main() {
	ctx := context.Background()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	resolver := graph.NewResolver(&ctx, logger, db)

	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{Resolvers: resolver}),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logger.Info("connect to http://localhost:" + port + "/ for GraphQL playground")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Warn("Exit", "err", err)
	}
}
