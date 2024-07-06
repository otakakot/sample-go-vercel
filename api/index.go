package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	pool, err := GetPool(r.Context(), GetDSN())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	defer pool.Close()

	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}

func GetDSN() string {
	return "postgres://default:" + os.Getenv("POSTGRES_PASSWORD") + "@" + os.Getenv("POSTGRES_HOST") + ":5432/" + os.Getenv("POSTGRES_DATABASE") + "?sslmode=require"
}

func GetPool(
	ctx context.Context,
	dsn string,
) (*pgxpool.Pool, error) {
	conn, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, conn)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}
