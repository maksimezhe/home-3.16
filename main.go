package main

import (
	"fmt"
)

// Task 1
// func printData(data string) {
// 	fmt.Println("Обработка данных:", data)
// }

// Task 3
// func checkAge(age int) {
// 	if age < 0 {
// 		panic("Возраст не может быть отрицательным!")
// 	}
// }

// Task 4
// func cleanup() {
// 	fmt.Println("Очистка ресурсов выполнена.")
// }

// Task 5
// func handlePanic() {
// 	r := recover()
// 	if r != nil {
// 		fmt.Println("Паника успешно перехвачена:", r)
// 	}
// }

// Task 6
func safeDivide(a, b float64) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Произошла паника:", r)
		}
	}()
	if b == 0 {
		panic("Деление на ноль!")
	}
	result := a / b
	fmt.Println("Результат деления:", result)
}

// Task 7
func getElement(slice []int, ind int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Произошла паника:", r)
		}
	}()
	fmt.Println("Элемент массива:", slice[ind])
}

func main() {
	// 1
	// fmt.Println("Старт программы")
	// defer printData("Данные для обработки")
	// fmt.Println("Конец программы")

	//Task 2
	// defer fmt.Println("Первый")
	// defer fmt.Println("Второй")
	// defer fmt.Println("Третий")

	// 3
	// age := -5
	// checkAge(age)

	// 4
	// defer cleanup()
	// panic("Критическая ошибка!")

	// // 5
	// defer handlePanic()
	// panic("Произошла паника!")

	// 6
	safeDivide(10, 0)

	// 7
	getElement([]int{1, 2, 3}, 10)
}
