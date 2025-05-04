package main

import (
	"fmt"
	"os"
	"strconv"
)
const balanceFileName = "balance.txt";
func writebalanceToFile(balance float64){
	balanceText := fmt.Sprint(balance)
	os.WriteFile(balanceFileName,[]byte(balanceText),0644)

}
func readBalanceFile () float64{
	data, _ := os.ReadFile(balanceFileName)
	// balanceText := string(data)
	balance, _ := strconv.ParseFloat(string(data), 64)
	return balance	
}

func main () {
	var accountBalance float64= readBalanceFile()
	for {
		communication()

		fmt.Print("Your Chocie: ")
		var choice int
		fmt.Scan(&choice)
		switch choice{
			case 1:
				fmt.Println("Your Account balance",accountBalance)
			case 2:
				fmt.Print("Your Deposit :")
					var depositAmount float64
					fmt.Scan(&depositAmount)
					accountBalance += depositAmount
					fmt.Println("Your new balance",accountBalance)
					writebalanceToFile(accountBalance)
			case 3:
				fmt.Print("Your Withdraw Amount:")
				var withdrawAmount float64
				fmt.Scan(&withdrawAmount)
				if accountBalance > 0 {
					accountBalance -= withdrawAmount
					fmt.Println("Your new account balance", accountBalance)
				} else {
					fmt.Println("Sorry you cannot withdraw your balance is", accountBalance)
					continue
					
					
				}
			default :
				fmt.Println("Goodbye!")
				return
		}


		// if choice == 1 {
		// 	fmt.Println("Your Account balance",accountBalance)
		// } else if choice  == 2 {
		// 	fmt.Print("Your Deposit :")
		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)
		// 	accountBalance += depositAmount
		// 	fmt.Println("Your new balance",accountBalance)
		// } else if choice == 3{
		// 	fmt.Print("Your Withdraw Amount:")
		// 	var withdrawAmount float64
		// 	fmt.Scan(&withdrawAmount)
		// 	if accountBalance > 0 {
		// 		accountBalance -= withdrawAmount
		// 		fmt.Println("Your new account balance", accountBalance)
		// 	} else {
		// 		fmt.Println("Sorry you cannot withdraw your balance is", accountBalance)
		// 		continue
				
				
		// 	}
		// } else {
		// 	fmt.Println("Goodbye!")
		// 	break
		// }

	}
	
	
}