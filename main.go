package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	fmt.Println("__ Калькулятор индекса массы тела __")
	for {
		userWeight, userHeight := getUserInput()
		IMT, err := calculateIMT(userWeight, userHeight)
		if err != nil {
			fmt.Println(err)
			continue
		}
		outputResult(IMT)
		isRepeateCalculation := checkRepeatCalculation()
		if !isRepeateCalculation {
			break
		}
	}
}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", IMT)
	fmt.Println(result)
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

func calculateIMT(userWeight, userHeight float64) (float64, error) {
	if userWeight <= 0 || userHeight <= 0 {
		return 0, errors.New("Не указан вес или высота")
	}
	IMT := userWeight / math.Pow(userHeight / 100, IMTPower)
	return IMT, nil
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

func checkRepeatCalculation() bool {
	var userChoice string
	fmt.Print("Вы хотите сделать еще расчёт (y/n): ")
	fmt.Scan(&userChoice)
	if  userChoice == "y" || userChoice == "Y" {
		return true
	}
	return false
}