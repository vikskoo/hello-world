package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	//boolean default false
	var firstBoolean bool
	fmt.Println(firstBoolean)

	//Boolean operands
	aBoolean, bBoolean := true, false
	fmt.Println("AND: ", aBoolean && bBoolean)
	fmt.Println("OR: ", aBoolean || bBoolean)
	fmt.Println("NOT: ", !aBoolean)

	//Numerics. Integers
	//int8, int16, int32, int64, int
	//uint8, uint16, uint32, uint64, uint

	var a int = 32
	b := 92
	fmt.Println("Value of a: ", a, "Value of b: ", b, "Value a+b: ", a+b)

	//Вывод типа через %T форматирование
	fmt.Printf("Type is %T\n", a)

	//Узнаем сколько байт занимает int
	fmt.Printf("Type %T size of %d bytes\n", a, unsafe.Sizeof(a))

	//Эксперимент (int 8 байт, 64 бита) int - платформозависимый
	fmt.Printf("Type %T size of %d bytes\n", b, unsafe.Sizeof(a))

	//Эскперимент 2 (Сложить разные типы нельзя, используем явное приведение типов что бы сложить если уверены что не будет коллизии)
	var first32 int32 = 12
	var second64 int64 = 13
	fmt.Println(int64(first32) + second64)

	//Эксперемент 3. Если проводится арифметич операции над int и intX то обязательно использовать приведение типов.
	//Несмотря на платформу int не значит int64
	var third64 int64 = 13636543
	var fourthInt int = 156267
	fmt.Println(third64 + int64(fourthInt))

	//Анологичным образом устроены uint8, uint16, uint32, uint64, uint

	//Numeric.Float
	//float32, float64
	floatFirst, floatSecond := 5.67, 12.54
	fmt.Printf("type of a %T and type of %T\n", floatFirst, floatSecond)
	//точность знаков после запятой %.3f
	fmt.Printf("Sum: %.3f\n", floatFirst+floatSecond)

	//Numeric. Complex
	c1 := complex(5, 7)
	c2 := 5 + 7i
	println(c1 + c2)

	//Strings
	name := "Иван"
	lastName := "Putin"
	concat := name + " " + lastName
	fmt.Println("Full name: ", concat)
	fmt.Println("Lenght of string", name, len(name))
	fmt.Println("Amount of chars: ", name, utf8.RuneCountInString(name))
	//Rune - это 1 utf-ный символ
	//Поиск подстроки в строке strings.Contains
	totalString, subString := "ABCDEFG", "ABC"
	fmt.Println(strings.Contains(totalString, subString))
	//rune -> alias int32
	//для инициализации символьным значением используем одинарные ковычки, только один символ
	var sampleRune rune
	var anotherRune rune = 'Q'
	var thirdRune rune = 23462
	fmt.Println(sampleRune)
	fmt.Printf("Rune as char %c\n", sampleRune)
	fmt.Printf("Rune as char %c\n", anotherRune)
	fmt.Printf("Rune as char %c\n", thirdRune)
	//Сравнение строк strings.Compare() -1 левая строка меньше правой, 0 если равна, 1 если больше
	fmt.Println(strings.Compare("abc", "abcd"))

}
