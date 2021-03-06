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
	config  Config
	logger  *logrus.Logger
	db      *sqlx.DB
	ctxExec context.Context
}

func NewPG(config  Config,	logger  *logrus.Logger)(s *Postgres,err error){
	s = &Postgres{
		config:  config,
		logger:  logger,
	}
	err = s.new()
	return
}

func (s *Postgres) new() (err error) {
	ctxConnect, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(s.config.TimeoutConnect))
	//ctxConnect, _ := context.WithCancel(context.Background())
	s.db, err = sqlx.ConnectContext(ctxConnect, "pgx", fmt.Sprintf("postgres://%s:%s@%s/%s", s.config.User, s.config.Pass, s.config.Server, s.config.Database))
	if err != nil {
		return err
	}
	s.ctxExec, _ = context.WithCancel(context.Background())
	//s.ctxExec, _ = context.WithTimeout(context.Background(),time.Minute*time.Duration(s.Config.TimeoutExecute))
	return err
}

func (s *Postgres) Shutdown() {
	s.logger.Infoln("Close Postgres connection...")
	err := s.db.Close()
	if err != nil { }
	s.logger.Infoln("Fail to close Postgres connection.")
}

func (s *Postgres) Add(e event.Event) (err error) {
	sql := "INSERT INTO events (StartTime, EndTime, Duration, TypeDuration,Title,Note,issending) VALUES (:starttime, :endtime, :duration, :typeduration,:title,:note,:issending);"
	_, err = s.db.NamedExecContext(s.ctxExec, sql, e)
	//_, err = s.db.NamedExec(sql,e)
	if (err!=nil){
		fmt.Print("add-",err.Error())
	}

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

func (s *Postgres) GetEventSending(endDаte time.Time) (events []event.Event, err error) {
	sql := fmt.Sprintf("SELECT * FROM events where issending = 0 and starttime < '%s';",endDаte.Format("2006-01-02 15:04:05")) // :id from `db:"id"`
	err = s.db.SelectContext(s.ctxExec, &events, sql)
	if err != nil {
		return events, err
	}
	return events, err
}

func (s *Postgres) MarkEventSentToQueue(id int) (err error) {
	sql := "update events set issending = 1 WHERE id = :id;"
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

func (s *Postgres) MarkEventSentToSubScribe(id int) (err error) {
	sql := "update events set issending = 2 WHERE id = :id;"
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
