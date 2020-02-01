package hw1

import (
	"github.com/beevik/ntp"
	"net"
	"time"
)

func CheckConnectionNTPServer(hostName string, port string, typeConnection string) (bool, error) {
	seconds := 5
	timeOut := time.Duration(seconds) * time.Second
	_, err := net.DialTimeout("udp", hostName+":"+port, timeOut)
	if err != nil {
		return false, err
	}
	return true, err
}

func GetTime(hostName string) (time.Time, error) {
	portNum := "123"
	t := time.Now()
	_, err := CheckConnectionNTPServer(hostName, portNum, "udp")
	if err != nil {
		return t, err
	}
	response, err := ntp.Query(hostName)
	if err != nil {
		return t, err
	}
	t = time.Now().Add(response.ClockOffset)
	return t, nil
}
