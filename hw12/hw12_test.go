package hw12

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func Test_CheckReafEnvVarFromFile(t *testing.T) {

	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	keyEnv, err := ReadDir(dir)
	if err != nil {
		fmt.Println(err.Error())
	}

	if val, ok := keyEnv["port"]; !ok || val != "8080" {
		t.Error("не найденна переменная port или не соответствует значению ", val)
	}

	if val, ok := keyEnv["server"]; !ok || val != "192.168.1.1;192.168.1.2" {
		t.Error("не найденна переменная server или не соответствует значению ", val)
	}

}

func Test_RunFile(t *testing.T) {

	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	keyEnv, err := ReadDir(dir)
	if err != nil {
		fmt.Println(err.Error())
	}
	s := []string{"writeListEnvVar.exe", "server", "port"}
	RunCmd(s, keyEnv)

	var result strings.Builder
	file, err := os.Open("listEnv.txt")
	if err != nil {
		t.Error("не создал тестовый файл.")
	}
	defer file.Close()

	data := make([]byte, 128)

	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		result.WriteString(string(data[:n]))
	}
	checkString := `server=>192.168.1.1;192.168.1.2;port=>8080`
	res := strings.TrimSpace(result.String())
	res = strings.Replace(res, "\n", ";", -1)
	res = strings.Replace(res, "\r", "", -1)
	if res != checkString {
		t.Error("не сответстие ожидаемой строки с тестовой.")
	}
}
