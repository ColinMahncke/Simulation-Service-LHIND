package functions

import (
	"fmt"

	"strconv"
)

//generate random people with a minimum and a maximum limit

func GenerateAreaCount(i float64) int {
	numbermin, error := strconv.Atoi(Getenv("area_min_rand", "150"))
	numbermax, error := strconv.Atoi(Getenv("area_max_rand", "200"))
	if error != nil {
		fmt.Println("error:", error)
		return 0
	}

	return randIntn(Bell(i)*float64(numbermin), Bell(i)*float64(numbermax))

}

func GenerateLineCount(i float64) int {
	numbermin, error := strconv.Atoi(Getenv("line_min_rand", "-5"))
	numbermax, error := strconv.Atoi(Getenv("line_max_rand", "5"))
	if error != nil {
		fmt.Println("error:", error)
		return 0
	}

	return randIntn(Bell(i)*float64(numbermin), Bell(i)*float64(numbermax))
}
