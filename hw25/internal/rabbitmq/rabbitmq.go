package rabbitmq

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type Config struct {
	User     string `yaml:"user"`
	Pass     string `yaml:"pass"`
	HostPort string `yaml:"host_port"`
	Timeout  int    `yaml:"timeout_connect"`
	Queue1   string `yaml:"queue1"`
	Queue2   string `yaml:"queue2"`
}

type RMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
	q2   amqp.Queue
	l    *logrus.Logger
	c    *Config
}

func NewRMQ(conf Config, logger *logrus.Logger) (r *RMQ, err error) {
	if logger != nil {
		logger.Infoln("Start connect to RMQ")
	}
	r = &RMQ{l: logger, c: &conf}
	r.conn, err = amqp.DialConfig(fmt.Sprintf("amqp://%s:%s@%s/", conf.User, conf.Pass, conf.HostPort),
		amqp.Config{Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, time.Second*time.Duration(conf.Timeout))
		}})
	if err != nil {
		return r, err
	}

	r.ch, err = r.conn.Channel()
	if err != nil {
		return r, err
	}

	r.q, err = r.ch.QueueDeclare(conf.Queue1, true, false, false, false, nil)
	if err != nil {
		return r, err
	}
	r.q2, err = r.ch.QueueDeclare(conf.Queue2, true, false, false, false, nil)
	if err != nil {
		return r, err
	}
	if logger != nil {
		logger.Infoln("Success connected to RMQ")
	}

	return r, nil
}

func (r *RMQ) Reconnect() (err error) {
	r.l.Warningln("Start try reconnect to RMQ")
	r.conn, err = amqp.DialConfig(fmt.Sprintf("amqp://%s:%s@%s/", r.c.User, r.c.Pass, r.c.HostPort),
		amqp.Config{Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, time.Second*time.Duration(r.c.Timeout))
		}})
	if err != nil {
		return err
	}

	r.ch, err = r.conn.Channel()
	if err != nil {
		return err
	}

	r.q, err = r.ch.QueueDeclare(r.c.Queue1, true, false, false, false, nil)
	if err != nil {
		return err
	}
	r.q2, err = r.ch.QueueDeclare(r.c.Queue2, true, false, false, false, nil)
	if err != nil {
		return err
	}
	r.l.Infoln("Success reconnected to RMQ")
	return err
}

func (r *RMQ) Shutdown() (err error) {
	r.l.Infoln("Close RMQ connect...")
	err = r.ch.Close()
	if err != nil {
		_ = r.conn.Close()
		return err
	}
	err = r.conn.Close()
	if err == nil {
		r.l.Infoln("Success close RMQ connect")
	}

	return err
}

func (r *RMQ) Send(message []byte) error {

	err := r.ch.Publish("", r.q.Name, false, false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			Body:         message,
		})

	return err
}

func (r *RMQ) Send2(message []byte) error {

	err := r.ch.Publish("", r.q2.Name, false, false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			Body:         message,
		})

	return err
}

func (r *RMQ) GetMsgsCh() (msgsCh <-chan amqp.Delivery, err error) {
	return r.ch.Consume(r.q.Name, "", false, false, false, false, nil)
}

func (r *RMQ) GetChanel(nameChanel string) (msg amqp.Delivery, ok bool, err error) {
	msg, ok, err = r.ch.Get(nameChanel, true)
	return
}

func handlerError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", msg, err)
	}
}
