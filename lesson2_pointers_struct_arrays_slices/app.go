package main

import "fmt"

func main() {
	pointers()
	structs()
	arrays()
	slices()
}

func pointers() {

	a := "Hello world"
	b := &a         // с помощью & получили ссылку на переменную a и присвоили эту ссылку переменной b.
	fmt.Println(b)  // вывели ссылку.
	fmt.Println(*b) // с помощью * получили значение ссылки.
	*b = "bye"      // с помощью * присвоили новое значение ссылке.

}

func structs() {

	/*
		в goLang struct является аналогом класса.
	*/

	p1 := myStruct{
		X: 15,
		Y: 20,
	}
	fmt.Println(p1)
	fmt.Println(p1.X)
	fmt.Println(p1.Y)
	fmt.Println(p1.line)
	//p1.Y = 99
	fmt.Println(p1.Y)
	p1.myStructMethod()

	p2 := myStruct{
		X:    778,
		line: "hhh",
	}
	p2.myStructMethod()
}

func arrays() {
	/*
		Создание пустого массива myArr на два элемента.
		Массивы используются редко.
		У массива статическая длина.
	*/
	var myArr [2]string
	myArr[0] = "Jack"
	myArr[1] = "Ivanov"
	fmt.Println(myArr)
	fmt.Println(myArr[0])
	fmt.Println(myArr[1])

	/*
		Создание массива с инициализацией.
		Оператор ... позволяет нам не записывать длину самим, вместо этого она будет автоматически выведена
		компилятором из количества элементов.
	*/
	numbers := [...]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Println(numbers[3])
	fmt.Println(numbers)
}

func slices() {
	/*
		slice это массив с динамической длиной. Очень похож на ArrayList в java.
		slice является обёрткой над обычным массивом.
		При создании slice нам не нужно указывать изначальную длину.
		У slice есть различные методы в отличие от массива.
	*/

	firstSlice := []string{"1", "2", "3", "4", "5"}
	firstSlice2 := append(firstSlice, "a", "b")
	fmt.Println(firstSlice)
	fmt.Println(firstSlice2)

	/*
		Создали новый slice с типом int, c начальной длинной 3.
		len - возвращает длину слайса (количество элементов в слайсе)
		cap - возвращает емкость слайса (количество элементов, до которых слайс может увеличиваться или уменьшаться)
	*/
	newSlice := make([]int, 3)
	fmt.Println("len: ", len(newSlice))
	fmt.Println("cap: ", cap(newSlice))
	newSlice = append(newSlice, 3)
	fmt.Println(newSlice)
	fmt.Println("len: ", len(newSlice))
	fmt.Println("cap: ", cap(newSlice))

	/*
		У слайса есть возможность выбрать часть слайса с помощью оператора [].
		Внутри себя слайс содержит массив. Соответственно если мы в одной части слайса изменим какое то значение, то это
		значение изменится во всех слайсах в основе которых лежит один и тот же массив.
	*/
	var animalsSlice = []string{"cat", "dog", "mouse", "horse"}
	var first = animalsSlice[0:2]
	var second = animalsSlice[1:4]
	var sliceFromZero = animalsSlice[:3]     // не указав начальное значение по дефолту будет стоять ноль.
	var sliceToEnd = animalsSlice[2:]        // не указав значение по дефлту будет срез до последнего элемента слайса.
	var sliceFromZeroToEnd = animalsSlice[:] // создастся слайс с нулевого элемента до последнего

	fmt.Println(animalsSlice)
	fmt.Println(first)
	fmt.Println(second)
	second[0] = "pig"
	fmt.Println(animalsSlice)
	fmt.Println(first)
	fmt.Println(second)

	fmt.Println(sliceFromZero)
	fmt.Println(sliceToEnd)
	fmt.Println(sliceFromZeroToEnd)
}

type myStruct struct {
	X    int
	Y    int
	line string
}

func (p myStruct) myStructMethod() {
	// Метод которйы относится к структуре myStruct

	if p.X > 0 {
		fmt.Println(p.X)
	}

	if p.Y > 0 {
		fmt.Println(p.Y)
	}
	if p.line != "" {
		fmt.Println(p.line)
	}

}
