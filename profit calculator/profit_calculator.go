package main

import "fmt"

func main(){
	revenue,expences,taxRate := needUserInput()
	ebt,profit,ratio := calcExpenses(revenue,expences,taxRate)

	fmt.Printf("Ratio %.2f",ratio)
	fmt.Printf("Ratio %.2f",ebt)
	fmt.Printf("Ratio %.2f",profit)
}
 
func needUserInput () (float64,float64,float64){
	var revenue float64
	var expences float64
	var taxRate float64

	fmt.Print("Enter your revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter your expences: ")
	fmt.Scan(&expences)

	fmt.Print("Enter your taxRate: ")
	fmt.Scan(&taxRate)

	return revenue,expences,taxRate

}
func calcExpenses (revenue,expenses,taxRate float64) (float64,float64,float64){
	ebt := revenue - expenses
	profit := ebt - (ebt *taxRate /100)
	ratio := ebt / profit

	return ebt,profit,ratio
}