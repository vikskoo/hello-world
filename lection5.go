package main

import (
	"fmt"
	"strings"
)

func main() {
	var value int
	fmt.Scan(&value)

	if value%2 == 0 {
		fmt.Println("The number", value, "is even")
	} else {
		fmt.Println("The number", value, "is odd")
	}

	var color string
	fmt.Scan(&color)
	if strings.Compare(color, "green") == 0 {
		fmt.Println("Color is green")
	} else if strings.Compare(color, "red") == 0 {
		fmt.Println("Color is red")
	} else {
		fmt.Println("Unknown color")
	}

	//Good. Инициализация в блоке условного оператора
	//Блок присваивания - только :=
	//Инициализированная переменная видна только внутри области видимости условного оператора
	if num := 10; num%2 == 0 {
		fmt.Println("EVEN")
	} else {
		fmt.Println("ODD")
	}

	//Не идеоматично
	if width := 100; width > 100 {
		fmt.Println("Width > 100")
	} else {
		fmt.Println("width <= 100")
	}

	//Странное правило номер 1: в Go стараются избегать блоков Else
	// Идеоматично!
	if height := 100; height > 100 {
		fmt.Println("height > 100")
		return
	}
	fmt.Println("Height <= 100")
}
