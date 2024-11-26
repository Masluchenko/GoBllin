package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	fmt.Println("__Калькулятор индекса массы тела__")

	for {

		userKg, userHeight := getUserInput()
		IMT, err := calculateIMT(userKg, userHeight)
		if err != nil {
			// fmt.Println(err)
			// continue
			panic("Незаданы параметры для расчета")
		}
		outputResult(IMT)
		isRepeatCalculation := checkRepeatCalculating()
		if !isRepeatCalculation {
			fmt.Print("BB")
			break
		}
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f\n", imt)
	fmt.Print(result)
	switch {
	case imt < 16:
		fmt.Println("ультра недоношенный")
	case imt < 18.5:
		fmt.Println("Человек вата")
	case imt < 25:
		fmt.Println("Типикл нормис")
	case imt < 30:
		fmt.Println("Кабанчик мясистый")
	}
}

func calculateIMT(userKg float64, userHeight float64) (float64, error) {
	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New("NO_PARAMS_ERROR")
	}
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	return userKg, userHeight
}

func checkRepeatCalculating() bool {
	var userChoise string
	fmt.Println("Вы хотите сделать ещерасчет? (y/n)")
	fmt.Scan(&userChoise)
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}
