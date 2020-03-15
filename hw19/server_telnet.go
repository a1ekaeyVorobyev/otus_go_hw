package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	adress := fmt.Sprintf("Welcome to %s, friend from %s\n", conn.LocalAddr(), conn.RemoteAddr())
	conn.Write([]byte(adress))

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {

		text := scanner.Text()
		log.Printf("Recived text: %s", text)
		if text == "quit" || text == "exit" {
			log.Printf("Closing connection with keypress %s", text)
			conn.Write([]byte(fmt.Sprintf("Closing connection with keypress %s", text)))
			conn.Close()
			return
		}

		conn.Write([]byte(fmt.Sprintf("I get text: '%s'\n", text)))
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error connection  %s: %v", conn.RemoteAddr(), err)
	}

	log.Printf("Closing connection %s", conn.RemoteAddr())

}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	loader, err := net.ResolveTCPAddr("tcp", "0.0.0.0:3125")
	if err != nil {
		log.Fatalf("Cannot listen: %v", err)
	}
	l, err := net.ListenTCP("tcp", loader)
	if err != nil {
		log.Fatalf("Cannot listen: %v", err)
	}
	defer l.Close()

	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
		<-sigs
		cancel()
	}()

exit:
	for {

		select {
		case <-ctx.Done():
			log.Println("Exit telnet.")
			break exit
		default:
			err := l.SetDeadline(time.Now().Add(time.Second))
			if err != nil {
				log.Println(err.Error())
				continue
			}
			conn, err := l.Accept()
			if os.IsTimeout(err) {
				continue
			}
			if err != nil {
				log.Fatalf("Cannot accept: %v", err)
			}
			log.Println("Connection telnet.")
			go handleConn(conn)
		}
	}

}
