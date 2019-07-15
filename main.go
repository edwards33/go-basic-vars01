package main

import "fmt"

func main() {
	// init
	var num0 int

	// on init
	var num1 int = 1

	// without type
	var num2 = 20
	fmt.Println(num0, num1, num2)

	// variable short declaration
	num := 30
	// only for new variables
	// no new variables on left side of :=
	// num := 31

	num += 1
	fmt.Println("+=", num)

	// ++num unexists
	num++
	fmt.Println("++", num)

	// camelCase
	userIndex := 10
	// under_score: not accepted
	user_index := 10
	fmt.Println(userIndex, user_index)

	// multiple variable declaration
	var weight, height int = 10, 20

	// assignment to existing variables
	weight, height = 11, 21

	// short assignment
	// at least one variable should be new variable!
	weight, age := 12, 22

	fmt.Println(weight, height, age)
}
