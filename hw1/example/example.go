package main

import (
	"fmt"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw1"
	"log"
	"os"
	"time"
)

func main() {
	listServer := []string{"0.pool.ntp.org",
		"1.pool.ntp.org",
		"2.pool.ntp.org",
		"3.pool.ntp.org"}
	t := time.Now()
	fmt.Printf("Локальное время %s\n", t.Format(time.RFC3339))
	for _, v := range listServer {
		t, err := hw1.GetTime(v)
		if err != nil {
			log.Fatalf(err.Error())
			os.Exit(1)
		}
		fmt.Printf("Время с ntp сервера %s время %s\n", v, t.Format(time.RFC3339))
	}
}
