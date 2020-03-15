package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Printf("Set adress server and port: example 127.0.0.1  25 --timeout=10s , default timeout = 10s")
		os.Exit(2)
	}
	host := os.Args[1]
	port := os.Args[2]

	var timeout int
	flag.IntVar(&timeout, "timeout", 10, "time out in seconds")
	flag.Parse()

	connect, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), time.Second*time.Duration(timeout))
	defer connect.Close()
	if err != nil {
		fmt.Printf("Can't connect to telnet server: %v\n", err)
		os.Exit(2)
	}
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan bool, 1)
	go hanlerRead(ctx, connect, ch)
	go handlerwriter(ctx, connect, ch)
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
			break exit
		case <-ch:
			break exit
		}

	}
	log.Println("Exit telnet.")
}

func hanlerRead(ctx context.Context, connect net.Conn, ch chan bool) {
	defer func() { ch <- true }()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if !scanner.Scan() {
				return
			}
			str := scanner.Text()
			_, err := connect.Write([]byte(fmt.Sprintf("%s\n", str)))
			if err != nil {
				return
			}
			if str == "quit" || str == "exit" {
				return
			}
		}
	}
}

func handlerwriter(ctx context.Context, connect net.Conn, ch chan bool) {
	defer func() { ch <- true }()
	scanner := bufio.NewScanner(connect)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if !scanner.Scan() {
				return
			}
			text := scanner.Text()
			fmt.Println(text)
		}
	}
}
