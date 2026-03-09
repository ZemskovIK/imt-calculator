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
}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", IMT)
	fmt.Print(result)
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