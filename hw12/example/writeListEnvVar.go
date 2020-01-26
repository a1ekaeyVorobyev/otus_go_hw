package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Create("listEnv.txt")
	if err != nil {
		fmt.Println("Не могу создать файл:", err)
		os.Exit(1)
	}
	defer file.Close()

	if len(os.Args) == 1 {
		for _, element := range os.Environ() {
			variable := strings.Split(element, "=")
			if variable[0] != "" {
				text := fmt.Sprintf("%v=>%v\n", variable[0], variable[1])
				file.WriteString(text)
			}
		}
	} else {
		r := os.Args[1:]
		for _, v := range r {
			user, ok := os.LookupEnv(v)
			if !ok {
				user = "не найден"
			}
			text := fmt.Sprintf("%v=>%v\n", v, user)
			file.WriteString(text) // rob
		}
	}
	fmt.Println("Done.")
}
