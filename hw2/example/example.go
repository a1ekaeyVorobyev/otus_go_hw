package main

import (
	"fmt"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw2"
)

func main() {

	testStrings := map[string]string{
		"0123456789": "Не верный формат",
		"a4bc2d5e":   "aaaabccddddde",
		"abcd":       "abcd",
		"45":         "Не верный формат",
		`qwe\4\5`:    "qwe45",
		`qwe\45`: 		"qwe44444",
		`qwe5\4`:     "qweeeee4",
		`qwe\\5`:     `qwe\\\\\`,
		`q\we`:       "Не верный формат",
		"ab/c":       "ab/c",
	}

	for key, val := range testStrings {
		result, err := hw2.UnPackString(key)

		if err != nil {
			fmt.Println(key, err.Error())
		} else {
			fmt.Println(key, result, val)
		}
	}

}
