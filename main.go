package main

import "fmt"

func main() {
	revenue := getInput("Enter your revenue: ");
	expenses := getInput("Enter your total expenditure: ")
	taxRate := getInput("Tax Rate: ") 
	
	eBT, profit, ratio := calculateEbtProfitRate(revenue, expenses, taxRate)


	fmt.Printf(	"ebt = %.2f \nprofit = %.2f \nratio = %.3f", eBT, profit, ratio)

}

func getInput(question string) float64{
	var variable float64

	fmt.Print(question)
	fmt.Scan(&variable)

	return variable
}

func calculateEbtProfitRate(revenue, expenses, taxRate float64) (eBT, profit, rate float64){
	eBT = revenue - expenses
	profit = eBT - (1-taxRate/100)

	rate = eBT/ profit

	return eBT, profit, rate
}