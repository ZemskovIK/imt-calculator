package main

import (
	"fmt"
	"math"
)

func main() {
    fmt.Println("__ Калькулятор индекса массы тела __")
	userWeight, userHeight := getUserInput()
	IMT := calculateIMT(userWeight, userHeight)
	outputResult(IMT)
	switch {
	case IMT < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case IMT < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case IMT < 25:
		fmt.Println("У вас нормальный вес")
	case IMT < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас степень ожирения")
	}
}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", IMT)
	fmt.Println(result)
}

func calculateIMT(userWeight, userHeight float64) float64 {
	const IMTPower = 2
	IMT := userWeight / math.Pow(userHeight / 100, IMTPower)
	return IMT
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userWeight float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userWeight)
	return userWeight, userHeight
}