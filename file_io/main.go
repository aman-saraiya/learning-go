package main

func main() {
	balance := 10000
	writeIntToFile(balance)
	readIntFromFile()

	balance += 1000
	writeIntToFile(balance)
	readIntFromFile()
}
