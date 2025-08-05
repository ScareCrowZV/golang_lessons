package main

import "fmt"

type Namer interface {
	showMyName() string
}

type Developer interface {
	writeHelloWorld() string
}

type SuperDeveloper struct {
	DeveloperSQL
}

type DevLanguage struct {
	Name       string // Название языка
	Complexity int    // Значения от 1 до 10
}

func (d DevLanguage) showMyName() string {
	return "язык программирования " + d.Name
}

type DeveloperSQL struct {
	Name string
	DevLanguage
}

func (a DeveloperSQL) writeHelloWorld() string {
	return "SELECT 'Hello World!'"
}

func (a DeveloperSQL) showMyName() string {
	return "разработчик SQL. Меня зовут " + a.Name
}

func (a DeveloperSQL) String() string {
	return fmt.Sprintf("%s %s", "Developer", a.Name)
}

type DeveloperGO struct {
	Name string
	DevLanguage
	Grade string // junior, middle, senior
}

func (a DeveloperGO) writeHelloWorld() string {
	return "fmt.PrintLn(\"Hello World!\")"
}

func (a DeveloperGO) showMyName() string {
	return "разработчик Go. Меня зовут " + a.Name
}

func writeCode(d Developer) {
	fmt.Println("im developer", d.writeHelloWorld())
}

func sayHello(t Namer) {
	fmt.Println("Всех приветствую! Я", t.showMyName())
}

func main() {
	// Создаём экземпляр типа языка SQL
	SQLLang := DevLanguage{
		Name:       "SQL",
		Complexity: 4,
	}
	// Создаём экземпляр типа разработчика SQL
	devSQL := DeveloperSQL{
		Name:        "Petya",
		DevLanguage: SQLLang,
	}

	// Создаём экземпляр типа языка Go
	GOLang := DevLanguage{
		Name:       "GoLang",
		Complexity: 5,
	}
	// Создаём экземпляр типа разработчика Go
	devGO := DeveloperGO{
		Name:        "Vasya",
		DevLanguage: GOLang,
		Grade:       "senior",
	}
	// Пишем код с помощью общей функции writeCode()
	writeCode(devSQL)
	writeCode(devGO)

	// Позволяет избавиться от копипасты ...
	sayHello(devSQL)
	sayHello(devGO)
	sayHello(SQLLang)
	sayHello(GOLang)

	// ...иначе это было бы так
	fmt.Println("Всех приветствую! Я", GOLang.showMyName())
	fmt.Println("Всех приветствую! Я", SQLLang.showMyName())
	fmt.Println("Всех приветствую! Я", devGO.showMyName())
	fmt.Println("Всех приветствую! Я", devSQL.showMyName())

	// Полиморфизм интерфейса. Интерфейс может принимать разные типы, если они имплементировали интерфейс.
	var dev Developer = devSQL
	fmt.Println(dev.writeHelloWorld())
	dev = devGO
	fmt.Println(dev.writeHelloWorld())

	// dev = SQLLang // Ошибка. Тип DevLanguage не имплементирует интерфейс Developer, а именно не реализует его метод writeHelloWorld()

	// Ещё полиморфизм интерфейса.
	var nm Namer
	nm = GOLang
	fmt.Println(nm.showMyName())
	nm = devGO
	fmt.Println(nm.showMyName())
	sayHello(nm)

	// Композиция типов и соответствие интерфейсам
	var sd SuperDeveloper = SuperDeveloper{devSQL}
	sayHello(sd)

	// Проверка типа
	var checkType interface{} = sd

	// Вариант 1: Явное приведение к какому-либо типу
	val, ok := checkType.(SuperDeveloper)
	if ok {
		fmt.Println(val.Complexity)
	}

	// Вариант 2: Проверяем множество типов через switch case
	switch checkType.(type) {
	case DevLanguage:
		fmt.Println("SQL Lang")
	case SuperDeveloper:
		fmt.Println("Super Developer")
	default:
		fmt.Println("неизвестный тип")
	}

	// Так же есть 3 вариант через пакет reflect

	// Пример подстраивание под интерфейс. Выше для DeveloperSQL есть реализация метода String() интерфейса Stringer() из стандартного пакета fmt
	fmt.Println(devSQL)
	fmt.Println(devGO)
}
