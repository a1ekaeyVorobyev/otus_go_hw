package hw4

import (
	"testing"
)


func Test_Check(t *testing.T) {
	text := `Если есть более 10 самых частотых слов (например 15 разных слов встречаются ровно 133 раза, остальные < 100), можно вернуть любые 10 из самых частотных.

Словоформы не учитываем. "нога", "ногу", "ноги" - это разные слова.
Слово с большой и маленькой буквы можно считать за разные слова. "Нога" и "нога" - это разные слова.
Знаки препиания можно считать "буквами" слова или отдельными словами. "-" (тире) - это отдельное слово. "нога," и "нога" - это разные слова.

Пример: "cat and dog one dog two cats and one man". "dog", "one", "and" - встречаются два раза, это топ-3.

Задание со звездочкой (*): учитывать большие/маленьгие буквы и знаки препинания. "Нога" и "нога" - это одинаковые слова, "нога," и "нога" - это одинаковые слова, "—" (тире) - это не слово.
Критерии оценки: Функция должна проходить все тесты
Код должен проходить проверки go vet и golint
У преподавателя должна быть возможность скачать и проверить пакет с помощью go get / go test
Задание (*) НЕ влияет на баллы, оно дано просто для развития навыков.`
	resut := CountWord(text,10)
	check := map[string]int{"and":3, "go":3, "one":3, "и":8, "не":3, "нога":9, "разные":4, "слова":7, "слово":3, "это":8}
	for key,val := range resut{
		if check[key] != val{
			t.Error("Не соответстует тестовому результату")
		}
	}
}