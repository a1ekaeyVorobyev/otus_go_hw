package hw4

import (
	"sort"

	//"sort"
	"strings"
	"unicode"
)

type wordCount struct {
	word  string
	count int
}

func deletePunctFinish(str string) string {

	v := []rune(str)
	index := 0
	for i := len(v) - 1; i >= 0; i-- {
		if !(unicode.IsPunct(v[i]) || unicode.IsControl(v[i])) {
			break
		}
		index++
	}

	return string(v[0 : len(v)-index])

}

func deletePunctStart(str string) string {

	v := []rune(str)

	index := 0

	for i :=  0 ; i < len(v); i++ {

		if !(unicode.IsPunct(v[i]) || unicode.IsControl(v[i])) {

			break

		}

		index++

	}

	return string(v[index:len(v)])

}

func deletePunct(str string) string {

	str = deletePunctStart(str)

	return deletePunctFinish(str)

}

func findindex(words []wordCount,value string)(int,bool) {

	for i, v := range words {
		if v.word == value{
			return i,true
		}
	}
	return 0, false
}

func CountWord(text string,counWord int) map[string]int{

	text = strings.ReplaceAll(text,"\n"," ")
	arrayString := strings.Split(text," ")
	wordCounts := make([]wordCount,len(arrayString))
	index :=0
	for _,v := range arrayString{
		r := strings.ToLower(deletePunct(v))
		if r != "" {
			if val, ok := findindex(wordCounts,r); ok {
				wordCounts[val].count++
			} else {
				wordCounts[index] = wordCount{r,1}
				index++
			}
		}
	}
	sort.Slice(wordCounts, func(i, j int) bool { return wordCounts[i].count > wordCounts[j].count })
	//fmt.Println(wordCounts)
	result := make(map[string]int)
	index = counWord
	for _, v := range wordCounts{
		if v.count ==0{
			break
		}else{
			result[v.word] = v.count
			index--
		}
		if index ==0{
			break
		}
	}
	return result
}