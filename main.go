package main

import (
	"context"
	"database/sql"
	"log"
	"orm/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func main() {
	client := Open("postgresql://kyh0703:123qwe@127.0.0.1/testent")

	// Your code. For example:
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}
	users, err := client.User.Query().All(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(users)

	u, err := client.User.Create().
		SetName("10").
		SetAge(10).Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(u)
}
