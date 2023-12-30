package main

import "fmt"

func main() {
	userNames := []string{}
	userNames = append(userNames, "John", "Alex")
	fmt.Println(userNames)
	// when we create an empty slice, Go behind the scene creates
	// array for it and then once we start adding items (using append) it will
	// start creating new arrays.
	// This works but if we already have a rough estimate beforehand of how
	// many elements we might add, then we could preallocate some space
	// This will lead Go to create a bigger array right from the start and
	// thus reducing some overhead when we append new elements

	// we can specify the size to allocate space using the make function
	userNames = make([]string, 2, 4) //(type of slice to create, initial slice length, capacity)
	// Go allocates space to accomodate upto the capacity value
	userNames[0] = "John"
	userNames[1] = "Alex"
	userNames = append(userNames, "Ray")
	fmt.Println(userNames)

	// creating maps directly
	examScores := map[string]float64{}
	examScores["English"] = 93.4
	examScores["Mathematics"] = 95.9
	fmt.Println(examScores)

	// creating maps using make
	examScores = make(map[string]float64, 3)
	// here we only need 1 more argument as for maps we can't
	// really have initial empty slots, so the second argument
	// is used just to specify the intended length of the map
	examScores["English"] = 93.4
	examScores["Mathematics"] = 95.9
	examScores["Science"] = 97.5
	// so here upto 3 key value pairs, Go doesn't have to
	// recreate new arrays

	// adding the 4th key now requires to reallocate space
	examScores["Computer"] = 100

	// Working with type aliases
	// types (map[string]float64) like this are rather long so,
	// we could create a type alias for such types and assign it
	// a shorter name

	courseRatings := make(floatMap, 3)
	courseRatings["English"] = 4.6
	courseRatings["Mathematics"] = 4.8
	courseRatings["Science"] = 4.9
	fmt.Println(courseRatings)

	// using for loops with arrays, slices and maps
	// this works irrespective of the way the collection was
	// created, with or without using make
	for idx, val := range userNames {
		fmt.Println(idx, val)
	}

	// if we don't care about items or values we can just do this also
	count := 0
	for range userNames {
		count++
	}
	fmt.Println("Count:", count)

	// using with maps
	for key, val := range courseRatings {
		if key == "Mathematics" {
			val = 4.5 // this doesn't change the actual map
		}
		fmt.Println(key, val)
	}
	fmt.Println(courseRatings)
}

type floatMap map[string]float64
