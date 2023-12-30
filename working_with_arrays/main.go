package main

import "fmt"

type country struct {
	name    string
	capital string
}

func main() {
	// declaring and initializing an array
	countries := [3]string{"India", "China", "United States"}
	fmt.Println(countries)

	// accessing array elements
	fmt.Println(countries[0])
	fmt.Println(countries[1:])

	// creating slices with first and second element
	slice1 := countries[0:2]
	slice2 := countries[:2]
	fmt.Println(slice1)
	fmt.Println(slice2)

	// expanding the slice to the right
	slice1 = slice1[:3]
	slice2 = slice2[:3]
	fmt.Println(slice1)
	fmt.Println(slice2)
	// slice3 = slice1[1:]
	// this will not expand the slice, in order to expand we need to explicitly specify the end index

	// dynamic list
	capitals := []string{"Delhi", "Washington"}
	fmt.Println(capitals)

	// modifying value of a slice element
	capitals[1] = "Beijing"

	// creating a new slice with additional element
	capitals = append(capitals, "Washington")

	// working with list of structs
	usa := country{
		"United States",
		"Washington",
	}
	india := country{
		"India",
		"Delhi",
	}
	china := country{
		"China",
		"Beijing",
	}
	countryList := []country{usa, india}

	countryList = append(countryList, china)
	fmt.Println(countryList)

	// a slightly shorter alternative syntax while creating arrays/slices of structs

	countryList = []country{
		{
			"United States",
			"Washington",
		},
		{
			"India",
			"Delhi",
		},
		{
			"China",
			"Beijing",
		},
	}
	// here we've skipped the 'country'{} at the start as all the elements are known to be
	// of country only in this slice
	fmt.Println(countryList)
}
