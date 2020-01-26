package main

import (
	"fmt"
	"io/ioutil"

	//"io"
	//"io/ioutil"
	"os"
)

//noinspection GoMissingReturn
func readFile(nameFile string )([]byte, error) {
	fmt.Println(nameFile)
	data, err := ioutil.ReadFile(nameFile)
	if err != nil {
		fmt.Println("File reading error", err)
		return data,err
	}
	return data,err
}




func main() {

	argsWithProg := os.Args
	indexInFile := -1
	indexOutFile := -1
	indexNTPServer := -1
	indexHelp :=-1

	for index, value := range argsWithProg {
		switch value {
		case "-f":
			indexInFile = index
		case "-o":
			indexOutFile = index
		case "-s":
			indexNTPServer = index
		case "-h":
			indexHelp = index
		}
	}
	if indexInFile>-1 {
		fmt.Println("файл загузки %[1]s", argsWithProg[indexInFile+1])
		data,err  := readFile(argsWithProg[indexInFile+1])
		fmt.Println("Contents of file:", string(data))
		fmt.Println("Ошибка",err)
	}
	if indexOutFile>-1 {
		fmt.Println("файл загузки %[1]s", argsWithProg[indexOutFile+1])
	}
	if indexNTPServer>-1 {
		fmt.Println("файл загузки %[1]s", argsWithProg[indexNTPServer+1])
	}
	if indexHelp>-1 {
		fmt.Println("файл загузки %[1]s", argsWithProg[indexHelp+1])
	}
	fmt.Println("Hello, playground")
}
