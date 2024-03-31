package main

import "fmt"

func main() {
	//rangeLoop()
	mapExample()
}

func rangeLoop() {
	/*
		range - это цикл.
		i - индекс массика
		l - элемент массива
	*/

	arr := []string{"a", "b", "c"}
	for i, l := range arr {
		fmt.Println(i, l)
	}
}

func mapExample() {
	map1 := map[string]Point{}  // создание map
	map2 := make(map[int]Point) // альтернативный вариант создания map
	var map3 map[string]Point   // объявили, но не проинициализировали
	map3 = map[string]Point{}   //  проинициализировали map3

	p1 := Point{1, 7}
	p2 := Point{4, 10}

	map1["a"] = p1 // добавление элемента в map
	map1["b"] = p2 // добавление элемента в map

	map2[1] = p1 // добавление элемента в map
	map2[2] = p2 // добавление элемента в map

	map4 := map[string]Point{ // создание и добавление в элемента в map
		"a": {5, 5},
	}

	fmt.Println(map1["a"])
	fmt.Println(map2[1])
	fmt.Println(map4)
	fmt.Println(map3)

	key := "a"
	value, ok := map1[key] // провекра на наличие ключа.
	if ok {
		fmt.Printf("map exist key \"%s\" with value: \n", key)
		fmt.Println(value)
	} else {
		fmt.Printf("map dont exist key \"%s\" with value: \n", key)
		fmt.Println(value)
	}

	for k, v := range map1 { // обход map с помощью цикла range
		fmt.Println(k, v)
	}

}

type Point struct {
	X int
	Y int
}
