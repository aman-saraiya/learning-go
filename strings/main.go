package main

import "fmt"

func main() {
	// %v is to output the value of a variable or constant
	// review official documentation for all formatting options
	const distanceEarthSun float64 = 93
	fmt.Printf("The distance between earth and sun is %v million miles.\n", distanceEarthSun)

	var populationOfIndia float64 = 1.4
	fmt.Printf("The population of India currently is %v billion.\n", populationOfIndia)

	// formatting the float value that is printed, default with %f is 6 digits after decimal point
	const PI float64 = 3.14159265359
	fmt.Printf("The value of PI is: %f\n", PI) // review official documentation for all formatting options
	fmt.Printf("The value of PI rounded to 3 decimal places: %.3f\n", PI)

	// Creating formatted strings using Sprintf
	roundedPiValue := fmt.Sprintf("The value of PI rounded off to 2 decimal places: %.2f", PI)
	fmt.Println(roundedPiValue)

	fmt.Print("x", "y", "\n") // Note: This does not add a space in between like fmt.Println
	fmt.Println("x", "y")     // Adds space between x and y and prints new line at the end

	// building multiline string
	// to do this we need to use backticks (`) instead of double quotes (")
	// this works wierdly

	fmt.Printf(
		`The Value of PI: %v
		Rounded off value of PI %.2f`,
		PI, PI)
}
