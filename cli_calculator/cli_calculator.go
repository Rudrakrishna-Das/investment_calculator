package main

import (
	"errors"
	"fmt"
	"strings"
)


func main() {
	var userWantsToRunTheProgram bool = true
	var userChoiceToRunTheProgram string
	// Two variables for store numbers
	for userWantsToRunTheProgram{
		number1, number2, calculationSign,err := userCoice()
	if err == nil{
		sum, err := calculation(number1, number2, calculationSign)
		if err != nil{
			fmt.Println(err)	
		} else{
			fmt.Println("The value",sum)
		}
	} else {
		fmt.Println(err)
	}
	fmt.Println("Do you want to do another calculation? :- (yes / no)")
	fmt.Scan(&userChoiceToRunTheProgram)
		if strings.ToLower(userChoiceToRunTheProgram) == "yes" || strings.ToLower(userChoiceToRunTheProgram) == "y"{
			userWantsToRunTheProgram = true
		} else{
			userWantsToRunTheProgram = false
		}
	}
	
	
	
}

func calculation(num1 float64, num2 float64, symbol string) (float64, error) {
	
	switch symbol {
	case "+":
		
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		return num1 / num2, nil
	default:
		return 0, errors.New("please Select a sighn from (+, -, *, /)")

	}
}
func userCoice () (float64, float64, string, error){
	const allowedSymbol = "+-*/"
	var num1 float64
	var num2 float64
	var symbol string
	fmt.Print("Please Select your 1st number :- ")
	fmt.Scan(&num1)
	if num1 < 0 {
		return 0,0,"", errors.New("sorry num1 cannot be lower than 0")
	}
	fmt.Print("Please Select your 2nd number :- ")
	fmt.Scan(&num2)
	if num1 < 0 {
		return 0,0,"", errors.New("sorry num2 cannot be lower than 0")
	}
	fmt.Print("Please Select your symbol from (+, -, *, /) ")
	fmt.Scan(&symbol)
	if !strings.Contains(allowedSymbol,symbol){
		return 0,0,"", errors.New("calculation sign should select from (+, -, *, /)")
	}
	return num1,num2,symbol, nil
}