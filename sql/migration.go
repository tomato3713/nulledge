package main

import (
	"context"
	"embed"
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tomato3713/nulledge/database"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

var (
	Migrations = migrate.NewMigrations()

	rollbackFlag = flag.Bool("rollback", false, "run rollback")
)

//go:embed *.sql
var sqlMigrations embed.FS

func init() {
	if err := Migrations.Discover(sqlMigrations); err != nil {
		panic(err)
	}

}

func rollback(ctx context.Context, db *bun.DB, migrator *migrate.Migrator) error {
	group, err := migrator.Rollback(ctx)
	if err != nil {
		return err
	}

	if group.IsZero() {
		fmt.Println("no new migrations to run. database is latest")
		return nil
	}

	fmt.Printf("migrated to %s\n", group)
	return nil
}

func migration(ctx context.Context, db *bun.DB, migrator *migrate.Migrator) error {
	group, err := migrator.Migrate(ctx)
	if err != nil {
		return err
	}

	if group.IsZero() {
		fmt.Println("no new migrations to run. database is latest")
		return nil
	}

	fmt.Printf("migrated to %s\n", group)
	return nil
}

func usage() {
	fmt.Printf("usage: \nmigration [flags]\n")
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	ctx := context.Background()

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	migrator := migrate.NewMigrator(db, Migrations)
	if err := migrator.Init(ctx); err != nil {
		panic(err)
	}

	if *rollbackFlag {
		fmt.Println("start database rollback")
		if err := rollback(ctx, db, migrator); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("start database migration")
		if err := migration(ctx, db, migrator); err != nil {
			panic(err)
		}
	}
}
