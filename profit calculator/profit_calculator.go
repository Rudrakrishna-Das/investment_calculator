package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)
const profitFileName = "profit.txt";
func writeFile (e,r,p float64){	
	results := fmt.Sprintf("EBT: %.2f\nRATIO: %.3f\nProfit : %.2f\n",e,r,p)
	os.WriteFile(profitFileName, []byte(results),0644)
	
}
func main(){
	revenue,expences,taxRate := needUserInput()
	ebt,profit,ratio := calcExpenses(revenue,expences,taxRate)
	writeFile(ebt,ratio,profit)

	fmt.Printf("Ratio %.2f",ratio)
	fmt.Printf("Ratio %.2f",ebt)
	fmt.Printf("Ratio %.2f",profit)
}
func errorHandle (amount float64) error {
	if amount <= 0.0{
		return errors.New("amount should be > 0")
	}
	return nil
}
 
func needUserInput () (float64,float64,float64){
	var revenue float64
	var expences float64
	var taxRate float64

	fmt.Print("Enter your revenue: ")
	fmt.Scan(&revenue)
	var err error
	err = errorHandle(revenue)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Print("Enter your expences: ")
	fmt.Scan(&expences)
	err = errorHandle(expences)
	if err != nil{
		log.Fatal(err)
	}
	

	fmt.Print("Enter your taxRate: ")
	fmt.Scan(&taxRate)
	err = errorHandle(revenue)
	if err != nil{
		log.Fatal(err)
	}
	

	return revenue,expences,taxRate

}
func calcExpenses (revenue,expenses,taxRate float64) (float64,float64,float64){
	ebt := revenue - expenses
	profit := ebt - (ebt *taxRate /100)
	ratio := ebt / profit

	return ebt,profit,ratio
}