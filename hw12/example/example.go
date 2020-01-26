package main

import (
	"fmt"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw12"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		dir := os.Args[1]
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			fmt.Println("нет такой директории")
			os.Exit(0)
		}
		keyEnv, err := hw12.ReadDir(dir)
		if err != nil {
			fmt.Println(err.Error())
		}

		s := os.Args[2:]
		hw12.RunCmd(s, keyEnv)
		//RunCmdCheck(s)
	} else {
		fmt.Println("задайте директорию и имя запускаемого файла")
	}
}
