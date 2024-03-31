package main

import "fmt"

func main() {
	callBackExample()
	closures()
	methods()
}

func callBackExample() {

	/*
		в golang можно передавать функции в качестве параметра.
		Создали две функции addition и multiply и передали их в качестве параметра в метод doSomething
	*/

	addition := func(n1 int, n2 int) int {
		return n1 + n2
	}
	multiply := func(n1 int, n2 int) int {
		return n1 * n2
	}

	fmt.Println(doSomething(addition, "plus"))
	fmt.Println(doSomething(multiply, "multiply"))
}

func closures() {
	/*
		замыкание - в качестве возвращаемого значения возвращается функция.
		Метод totalPrice() возвращает функцию, которую мы сохранили в переменной total и в дальнейшем передали в неё
		параметр и вызвали несколько раз.
		В результате у нас в функции totalPrice() переменная sum менялась в зависимости от параметра который мы передаём
		в метод total(). Тоесть функция total() имеет доступ к переменным функции totalPrice().
	*/

	total := totalPrice(1)
	fmt.Println(total(1))  // sum = 1 + 1 = 2
	fmt.Println(total(1))  // sum = 2 + 1 = 3
	fmt.Println(total(1))  // sum = 3 + 1 = 4
	fmt.Println(total(1))  // sum = 4 + 1 = 5
	fmt.Println(total(1))  // sum = 5 + 1 = 6
	fmt.Println(total(10)) // sum = 6 + 10 = 16
}

func methods() {

	point := Point{1, 2}
	newPoint := point.movePoint(1, 1) // получить новый объект newPoint

	fmt.Println("point(до изм.): ", point)
	fmt.Println("newPoint: ", newPoint)

	point.movePointPtr(1, 1) // изменили существующий объект point. Так как передали параметр по ссылке (*)
	fmt.Println("point(после изм.): ", point)
}

type Point struct {
	x, y int
}

func (p Point) movePoint(x, y int) Point {
	/*
		Этот метод возвращает новый объект несмотря на то, что возращает тот же объект который находится в агументах.
	*/
	p.x += x
	p.y += y
	return p
}

func (p *Point) movePointPtr(x, y int) {
	/*
		метод ничего не возвращает, но меняет объект по ссылке.
	*/
	p.x += x
	p.y += y
}

func totalPrice(initPrice int) func(int) int {
	sum := initPrice
	return func(x int) int {
		sum += x
		return sum
	}
}

func doSomething(callBack func(int, int) int, s string) string {
	res := callBack(4, 6)
	return fmt.Sprintf("%s: %d", s, res)
}
