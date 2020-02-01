package main

import (
	"testing"
)

func Test_CheckcConnection(t *testing.T) {
	listServer := []string{"0.pool.ntp.org",
		"1.pool.ntp.org",
		"2.pool.ntp.org",
		"3.pool.ntp.org"}
	for _, v := range listServer {
		_, err := CheckConnectionNTPServer(v, "123", "udp")
		if err != nil {
			t.Error("нет cooединения с сервером ", v)
		}
	}
}
