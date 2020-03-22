package storage

import (
	"context"
	"fmt"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw22/internal/calendar/event"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw22/internal/config"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/sirupsen/logrus"
	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	Config  config.Config
	Logger  *logrus.Logger
	db      *sqlx.DB
	ctxExec context.Context
}

func (s *Postgres) Init() (err error) {
	ctxConnect, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(s.Config.DBTimeoutConnect))
	s.db, err = sqlx.ConnectContext(ctxConnect, "pgx", fmt.Sprintf("postgres://%s:%s@%s/%s", s.Config.DBUser, s.Config.DBPass, s.Config.DBServer, s.Config.DBDatabase))
	if err != nil {
		return err
	}
	s.ctxExec, _ = context.WithCancel(context.Background())
	return err
}

func (s *Postgres) Shutdown() {
	s.Logger.Infoln("Close Postgres connection...")
	err := s.db.Close()
	if err != nil {
		s.Logger.Infoln("Success close Postgres connection.")
	}
	s.Logger.Infoln("Fail to close Postgres connection.")
}

func (s *Postgres) Add(e event.Event) (err error) {
	sql := "INSERT INTO events (start_time, end_time, title, description) VALUES (:start_time, :end_time, :title, :description);" // :start_time from `db:"start_time"` and so on
	_, err = s.db.NamedExecContext(s.ctxExec, sql, e)
	return err
}
