package main

import "fmt"

type DevLanguage struct {
	Name       string // Название языка
	Complexity int    // Значения от 1 до 10
}

type DeveloperSQL struct {
	Name string
	DevLanguage
}

type DeveloperGO struct {
	Name          string
	LanguagesProf DevLanguage
}

func (l DevLanguage) helloWorld() string {
	return "Hello, World!"
}

func (a DeveloperSQL) WriteCode(taskComplexity int) int {
	return taskComplexity + a.DevLanguage.Complexity
}

func (a DeveloperSQL) Optimize(CodeLines int) bool {
	switch {
	case CodeLines < 1000:
		return true
	case CodeLines >= 1000:
		return false
	}

	return true
}

func (a DeveloperGO) WriteCode(taskComplexity int) int {
	return taskComplexity + a.LanguagesProf.Complexity
}

func (a DeveloperGO) Optimize(CodeLines int) bool {
	switch {
	case CodeLines < 100:
		return true
	case CodeLines >= 100:
		return false
	}

	return true
}

func main() {

	// Вариант 1: Создаёт экземпляр структуры со значениями по умолчанию
	var SQLLanguage1 DevLanguage

	// Вариант 2: Так же как и в варианте 1
	SQLLanguage2 := DevLanguage{}

	SQLLang := DevLanguage{
		Name:       "SQL",
		Complexity: 4,
	}

	devSQL := DeveloperSQL{
		Name:        "Petya",
		DevLanguage: SQLLang,
	}

	fmt.Println(devSQL.WriteCode(5))
	fmt.Println(devSQL.Optimize(100500))

	// Демонстрация встраивания структур в структуры
	fmt.Println(SQLLang.helloWorld())
	// Теперь из разработчика мы можем обращаться к методу встроенной структуры DevLanguage
	fmt.Println(devSQL.helloWorld())

	GOLang := DevLanguage{
		Name:       "GoLang",
		Complexity: 5,
	}

	devGO := DeveloperGO{
		Name:          "Vasya",
		LanguagesProf: GOLang,
	}

	fmt.Println(devGO.WriteCode(5))
	fmt.Println(devGO.Optimize(99))
}
