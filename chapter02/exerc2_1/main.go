package main

import (
	"fmt"
	"killerasus/exerc2_1/tempconv"
)

func main() {
	var temp1 tempconv.Celsius = 50
	fmt.Printf("Initial temperature: %s\n", temp1)
	fmt.Printf("Converted to Fahrenheit: %s\n", tempconv.CToF(temp1))
	fmt.Printf("Converted to Kelvin: %s\n", tempconv.CToK(temp1))
	fmt.Printf("Water boiling temperature difference to freezing temperature: %s\n", tempconv.BoilingC-tempconv.FreezingC)
	fmt.Printf("Zero Kelvin in Celsius = %s\n", tempconv.KToC(tempconv.AbsoluteZeroK))
}
