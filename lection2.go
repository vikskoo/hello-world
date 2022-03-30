package main

import (
	"fmt"
	"math"
)

func main() {
	// Вывод аргумента + '\n'
	fmt.Println("Hello world")
	fmt.Println("Second line")

	//print - простой вывод
	fmt.Print("first")
	fmt.Print("second")
	fmt.Print("third")

	//Форматированный вывод Printf - стандартный вывод os.Stdout c флагами форматрирования
	fmt.Printf("\nHello, my name is %s\nMy age is %d\n", "Bob", 32)

	//Декларирование переменных
	var age int
	fmt.Println("My age is:", age)
	age = 32
	fmt.Println("Age after assignment:", age)

	//Декларирование и инициализация пользовательским значением
	var height int = 183
	fmt.Println("My height is:", height)

	//В чем полустрогость типизации?
	var uid = 12345
	fmt.Println("My uid:", uid)
	fmt.Printf("%T\n", uid) //посмотреть тип переменной uid

	//Декларирование и инициализация переменных одного типа
	var firstVar, secondVar int = 20, 30
	fmt.Printf("firstVar: %d SecondVar: %d\n", firstVar, secondVar)

	//Декларирование блока переменных
	var (
		personName string = "Bob"
		personAge  int    = 42
		personUID  int
	)
	fmt.Printf("Name: %s\nAge %d\nUID %d\n", personName, personAge, personUID)

	//Немного странного
	var a, b = 30, "Masha"
	fmt.Println(a, b)

	//повторное декларирование переменной приводит к ошибке компиляции
	//var a = 200
	//!Всегда типизировать переменные явно

	//Короткая декларация
	count := 10
	fmt.Println("Count", count)
	count++
	fmt.Println("Count", count)

	//Множественное присваивание через :=
	aArg, bArg := 10, 40
	fmt.Println(aArg, bArg)

	//Исключение из правила повторного декларирования переменных
	//(Если в блоке хотя бы одна новая переменная старые будут перезаписаны)
	bArg, cArg := 10, 60
	fmt.Println(aArg, bArg, cArg)

	//Пример
	width, length := 20.5, 30.4
	fmt.Printf("Min dimensional of rectangle is: %.2f\n", math.Min(width, length))

}
