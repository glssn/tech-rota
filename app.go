package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4/pgxpool"
)

type App struct {
	router *chi.Mux
	db     *pgxpool.Pool
}

type pgx struct {
	ctx  context.Context
	pool *pgxpool.Pool
}

func NewPGX() (*pgx, error) {
	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return &pgx{
		ctx:  ctx,
		pool: pool,
	}, nil
}

func (p *pgx) Close() {
	p.pool.Close()
}
