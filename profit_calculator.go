package main

import "fmt"

func Calculate_profit() {
	var revenue, expenses, taxRate float64


	getInput("Enter your revenue: ")
	getInput("Enter your total expenditure: ")
	getInput("Enter your tax rate: ")
	fmt.Print("Enter your revenue: ")
	
	eBT, profit, ratio := calculateEbtProfitRate(revenue, expenses, taxRate)


	formatedEbt := fmt.Sprintf("EBT is %.2f \n", eBT)
	fmt.Printf(formatedEbt)
	fmt.Printf("Profit is %.2f \n", profit)
	fmt.Printf("Ratio %.2f \n", ratio)

}

