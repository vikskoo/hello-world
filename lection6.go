package main

import "fmt"

func main() {
	//for init; condition; post {
	// init -  блок инициализации переменной цикла
	// condition - условие ( если верно тело цикла выполняется, если нет то цикл завершается )
	// post - изменение переменной цикла инкремент, декремент
	//}

	for i := 0; i <= 5; i++ {
		fmt.Printf("Current value: %d\n", i)
	}
	// в качестве init может быть использовано выражение присваивания только через :=
	// переменная i находится только внутри цикла! за пределами её не видно.
	// break прервать цикл

	for i := 0; i < 15; i++ {
		if i > 12 {
			break
		}
		fmt.Printf("Current value: %d\n", i)
	}
	fmt.Println("After for loop with break")

	//continue - команда прерывающая текущее выполнение тела цикла и передающая следующей интерации

	for i := 0; i <= 20; i++ {
		if i > 10 && i <= 15 {
			continue
		}
		fmt.Printf("Current value: %d\n", i)
	}
	fmt.Println("After for loop with continue")

	// Вложенные циклы и лейблы
	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println("Треугольник")

	//Иногда бывает плохо. С лейблами получше
outer:
	for i := 0; i <= 2; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i:%d and j:%d and sum i+j=%d\n", i, j, i+j)
			if i == j {
				break outer
			}
		}
	}

}
