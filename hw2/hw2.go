package hw2

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func repeat(v rune, r []rune) strings.Builder {
	var result strings.Builder
	cnt, _ := strconv.Atoi(string(r))
	for i := 1; i < cnt; i++ {
		result.WriteRune(v)
	}
	return result
}

func UnPackString(s string) (string, error) {
	var result strings.Builder
	v := []rune(s)
	index := 0
	isSkype := false
	if unicode.IsDigit(v[0]) {
		return "", fmt.Errorf("Не верный формат")
	}
	for i := 0; i < len(v); i++ {
		r := v[i]
		if r == '\\' && !isSkype {
			isSkype = true
			if index > 0 {
				s := repeat(v[index-1], v[index:i])
				result.WriteString(s.String())
				index = 0
			}
			continue
		}
		if r == '\\' && isSkype {
			isSkype = false
		}
		if !unicode.IsDigit(r) && r != '\\' && isSkype {
			return "", fmt.Errorf("Не верный формат")
		}
		if unicode.IsDigit(r) && index == 0 && !isSkype {
			index = i
		}
		if !unicode.IsDigit(r) && index > 0 {
			//fmt.Printf("r=%c, index=%v\n", r, index)
			s := repeat(v[index-1], v[index:i])
			result.WriteString(s.String())
			index = 0
		}
		if i == (len(v)-1) && index > 0 {
			s := repeat(v[index-1], v[index:i+1])
			result.WriteString(s.String())
		}
		if index > 0 {
			continue
		}
		isSkype = false
		result.WriteRune(r)
	}
	return result.String(), nil
}
