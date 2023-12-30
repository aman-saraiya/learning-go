package main

import "fmt"

func main() {
	// websites := map[string]string{} // map is declared as map[key_type]value_type
	// we could also add some initial values as
	website := map[string]string{
		"Google": "https://google.com",
		"AWS":    "https://aws.com",
	}
	fmt.Println(website)
	fmt.Println(website["Google"])

	// mututaing maps
	website["AWS"] = "https://aws.amazon.com"
	fmt.Println(website)

	// we can always add new key,value pairs to maps.
	// unlike having arrays and slices, maps are always dynamic
	website["Facebook"] = "https://facebook.com"
	fmt.Println(website)

	//deleting key value pair
	delete(website, "Google")
	fmt.Println(website)
}
