package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

type Server struct {
	Addr    string  // TCP address to listen on; ":telnet" or ":telnets" if empty (when used with ListenAndServe or ListenAndServeTLS respectively).
	Port	string
	//Handler Handler // handler to invoke; telnet.EchoServer if nil
	//TLSConfig *tls.Config // optional TLS configuration; used by ListenAndServeTLS.
	ConnType  string
	TimeOut time.Duration
	//Logger Logger
}

func getTimeout(timeOut string) (time.Duration,error){
	fmt.Println(timeOut)
	if len(timeOut)<2{
		return 0,fmt.Errorf("Не верный формат timeout")
	}
	val := timeOut[: len(timeOut)-1]
	types := timeOut[len(timeOut)-1:]
	i, err := strconv.Atoi(val)
	if err != nil{
		return 0,fmt.Errorf("Не верный формат timeout ")
	}
	switch types {
	case "s":
		return time.Second*time.Duration(i), nil
	case "h":
		return time.Hour *time.Duration(i), nil
	case "m":
		return time.Minute *time.Duration(i), nil
	default:
		return 0,fmt.Errorf("Не верный формат timeout ")
	}
}

func main() {
	flagTimeOut := flag.String("timeout", "10s", "a string")
	flag.Parse()
	timeOut,err := getTimeout(*flagTimeOut)
	if err != nil{
		fmt.Println("Не правельнно введен timeout", flagTimeOut)
		os.Exit(2)
	}
	if len(os.Args)<3 && len(os.Args)>4 {
		fmt.Println("Не введен host и порт", flag.Args())
		os.Exit(2)
	}
	server := Server{
		Addr:     os.Args[1],
		Port:     os.Args[2],
		ConnType: "tcp",
		TimeOut:  timeOut,
	}
	fmt.Println(server)
	l, err := net.Listen(server.ConnType, server.Addr+":"+server.Port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	log.Printf("Listening on " + server.Addr + ":" + server.Port)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		//conn.SetReadDeadline(time.Now().add(server.TimeOut))
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	fmt.Println(reqLen)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}