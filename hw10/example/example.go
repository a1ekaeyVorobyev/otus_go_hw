package main

import (
	"flag"
	"fmt"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw10"
)

func main() {

	from := flag.String("from", "", "source file name")
	to := flag.String("to", "", "destination file name")
	offset := flag.Int("offset", 0, "an offset")
	limit := flag.Int("limit", 64, "an limit")
	rewrite := flag.Bool("rewrite", true, "rewritefile")

	flag.Parse()

	if *from == "" || *to == "" {
		fmt.Print("Не заданны имена файлов")
	} else {
		fmt.Println(*from, *to, *limit, *offset, *rewrite)
		err := hw10.Copy(*from, *to, *limit, *offset, *rewrite)

		if err != nil {
			fmt.Print(err.Error())
		}
	}
}
