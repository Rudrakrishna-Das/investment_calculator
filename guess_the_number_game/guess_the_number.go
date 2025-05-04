package main

import (
	"errors"
	"fmt"
	"strings"
)
const ERROR_MESAGE = "please provide only Integer"
func main(){
	for{
		number, times, err := userChoice()
		if err != nil{
			fmt.Println(err)
		} else{
			t := 1
			for t <= times {
				fmt.Println(number, "x", t, "=", number *t)
				t += 1
			}
		}
		var runAgain string
		fmt.Print("Do you want to try again?")
		_,er := fmt.Scan(&runAgain)
		if er != nil || strings.ToLower(runAgain) == "no"{
			break
		}
	}
	
}

func userChoice() (int,int,error){
	var userChoice int
	var userMultiplytimes int
	
	fmt.Println("Enter a number: ")
	_,err := fmt.Scan(&userChoice)
	if err != nil{
		return 0, 0,errors.New(ERROR_MESAGE)
	}
	fmt.Println("How many multiples?")
	_,err = fmt.Scan(&userMultiplytimes)
	if err != nil{
		return 0, 0,errors.New(ERROR_MESAGE)
	}
	
	return userChoice,userMultiplytimes, nil
}