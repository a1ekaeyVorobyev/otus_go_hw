package storage

import (
	"context"
	"fmt"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/calendar/event"
	"log"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)


type Postgres struct {
	Config  Config
	Logger  *logrus.Logger
	db      *sqlx.DB
	ctxExec context.Context
}

func (s *Postgres) Init() (err error) {
	ctxConnect, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(s.Config.TimeoutConnect))
	s.db, err = sqlx.ConnectContext(ctxConnect, "pgx", fmt.Sprintf("postgres://%s:%s@%s/%s", s.Config.User, s.Config.Pass, s.Config.Server, s.Config.Database))
	if err != nil {
		return err
	}
	s.ctxExec, _ = context.WithCancel(context.Background())
	//s.ctxExec, _ = context.WithTimeout(context.Background(),time.Minute*time.Duration(s.Config.DBTimeoutExecute))
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
	sql := "INSERT INTO events (StartTime, EndTime, Duration, TypeDuration,Title,Note) VALUES (:starttime, :endtime, :duration, :typeduration,:title,:note);"
	_, err = s.db.NamedExecContext(s.ctxExec, sql, e)
	//_, err = s.db.NamedExec(sql,e)
	return err
}

func (s *Postgres) Delete(id int) (err error) {
	sql := "DELETE FROM events WHERE id = :id;"
	res, err := s.db.NamedExecContext(s.ctxExec, sql, struct {
		Id int `db:"id"`
	}{Id: id})
	if err != nil {
		return err
	}
	cntRow, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if cntRow != 1 {
		return fmt.Errorf("Fail to delete event id: %d", id)
	}
	return nil
}

func (s *Postgres) Get(id int) (e event.Event, err error) {
	sql := "SELECT * from events WHERE id = :id;" // :id from `db:"id"`
	rows, err := s.db.NamedQueryContext(s.ctxExec, sql, struct {
		Id int `db:"id"`
	}{Id: id})
	if err != nil {
		return e, err
	}

	rows.Next()
	if err := rows.StructScan(&e); err != nil {
		return e, err
	}

	return e, err
}

func (s *Postgres) GetAll() (events []event.Event, err error) {
	sql := "SELECT * FROM events;" // :id from `db:"id"`
	err = s.db.SelectContext(s.ctxExec, &events, sql)
	if err != nil {
		return events, err
	}
	return events, err
}

func (s *Postgres) CountRecord() (cnt int) {
	row := s.db.QueryRow("SELECT COUNT(*) FROM events")
	err := row.Scan(&cnt)
	if err != nil {
		log.Fatal(err)
	}
	return cnt
}

func (s *Postgres) IsBusy(e event.Event) (exist bool, err error) {
	var rows *sqlx.Rows
	sql := "SELECT true FROM events WHERE start_time < :end_time AND endtime > :starttime and id != :id;"
	rows, err = s.db.NamedQueryContext(s.ctxExec, sql, e)
	if err != nil {
		return exist, err
	}
	exist = rows.Next()
	return exist, err
}

func (s *Postgres) Clear() (err error) {
	sql := "DELETE FROM events;"
	res, err := s.db.NamedExecContext(s.ctxExec, sql, struct {
	}{})
	if err != nil {
		return err
	}
	cntRow, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if cntRow != 1 {
		return fmt.Errorf("Fail to delete all event ")
	}
	return nil
}

func (s *Postgres) Edit(e event.Event) (err error) {
	sql := "UPDATE public.events SET starttime=:starttime, endtime=:endtime, duration=:duration, typeduration=:typeduration, title=:title, note=:note WHERE  id=:id;"
	_, err = s.db.NamedExecContext(s.ctxExec, sql, e)
	return err
}
