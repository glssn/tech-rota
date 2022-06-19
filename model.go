package main

import (
	"fmt"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/georgysavva/scany/pgxscan"
)

type event struct {
	ID            int       `db:"id"`
	EventType     string    `db:"event_type"`
	DateStart     time.Time `db:"date_start"`
	DateEnd       time.Time `db:"date_end"`
	AllDay        bool      `db:"all_day"`
	RecurType     string    `db:"recur_type"`
	RecurInterval string    `db:"recur_interval"`
}

func (p *pgx) GetEvents() {
	var events []*event

	if err := pgxscan.Select(
		p.ctx, p.pool, &events, `select id, event_type, date_start, date_end, all_day, recur_type, recur_interval FROM public.events`,
	); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	spew.Dump(events)
}
