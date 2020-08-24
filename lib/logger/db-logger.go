package logger

import (
	"context"
	"time"

	"github.com/go-pg/pg/v9"
)

type DBLogger struct{}

func (d DBLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d DBLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	qr, err := q.FormattedQuery()
	Debugf("RTime: %d ns, Query: %s, QError: %s, Error:%s", time.Since(q.StartTime).Nanoseconds(), qr, err, q.Err)
	return nil
}
