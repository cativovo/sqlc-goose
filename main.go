package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cativovo/sqlc-goose/db"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

// type Db struct {
// 	ctx     context.Context
// 	queries *queries.Queries
// }
//
// func NewDb(ctx context.Context, conn *pgx.Conn) *Db {
// 	return &Db{
// 		ctx:     ctx,
// 		queries: queries.New(conn),
// 	}
// }
//
// func (d *Db) getTodos() ([]queries.Todo, error) {
// 	return d.queries.GetTodos(d.ctx)
// }
//
// func (d *Db) insertTodos(arg queries.InsertTodoParams) (queries.Todo, error) {
// 	return d.queries.InsertTodo(d.ctx, arg)
// }

func run() error {
	ctx := context.Background()

	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_SSL_MODE"))
	fmt.Println(connString)
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := db.New(conn)

	todos, err := queries.GetTodos(ctx)
	if err != nil {
		return err
	}
	fmt.Println("todos", todos)

	insertArgs := db.InsertTodoParams{
		Task:        "task me",
		Description: pgtype.Text{String: "eyy description", Valid: true},
	}
	todo, err := queries.InsertTodo(ctx, insertArgs)
	if err != nil {
		return err
	}
	fmt.Println("added new todo", todo)

	{
		todos, err := queries.GetTodos(ctx)
		if err != nil {
			return err
		}
		fmt.Println("todos", todos)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
