package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tomato3713/nulledge/graph"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

const defaultPort = "8080"

func main() {
	ctx := context.Background()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	sqldb, err := sql.Open("mysql", getDbConnStr())
	if err != nil {
		panic(err)
	}

	db := bun.NewDB(sqldb, mysqldialect.New())

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

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

func getDbConnStr() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_ADDR"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_TABLENAME"))
}
