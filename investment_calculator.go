package main

import (
	"fmt"
	"math" 

)
const inflationRate = 6.5

func main (){
	
	var investmentAmount float64 = 1000
	var expectedReturnRate = 5.5
	var years float64 = 10

	// variables can be changes or re-assingned
	
	// Scan scans text read from standard input, storing successive space-separated values into successive arguments. Newlines count as space. It returns the number of items successfully scanned. If that 
	// is less than the number of arguments, err will report why.
	// investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5

	fmt.Print("Investment Amout:  ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)  
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, inflationRate, years) 
	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	formatFV := fmt.Sprintf("Future value: %.2f\n", futureValue)
	formatFRV := fmt.Sprintf("Future real vlaue (adjusted for inflation): %.2f", futureRealValue)

	fmt.Print(formatFV, formatFRV)

// output information
	//
	// fmt.Printf(`Future value is: %.2f
	
	// Future real value is: %.2f`, futureValue, futureRealValue)
	// fmt.Println("Future value is: ",futureValue)
	// fmt.Println("Future value (adjusted for inflation)", futureRealValue)
}
func calculateFutureValues(investmentAmount, expectedReturnRate, years float64)(fv float64, rfv float64){
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	// here we use the return keyword
	// return fv, rfv
	return
	
}