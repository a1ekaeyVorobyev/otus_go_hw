package hw2

import (
	"testing"
)

func Test_CheckCreate(t *testing.T) {

	testStrings := map[string]string{
		"0123456789": "Не верный формат",
		"a4bc2d5e":   "aaaabccddddde",
		"abcd":       "abcd",
		"45":         "Не верный формат",
		`qwe\4\5`:    "qwe45",
		`qwe\45`:     "qwe44444",
		`qwe5\4`:     "qweeeee4",
		`qwe\\5`:     `qwe\\\\\`,
		`q\we`:       "Не верный формат",
		"ab/c":       "ab/c",
	}
	var check string

	for key, val := range testStrings {
		result, err := UnPackString(key)

		if err != nil {
			check = err.Error()
		} else {
			check = result
		}

		if check != val {
			t.Error(check, " Не cooтветствует ", val)
		}
	}
}
