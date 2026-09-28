package main
import "fmt"

func main() {
	var fahrenheit int
	var celcius float64

	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scanln(&fahrenheit)
	celcius = (float64(fahrenheit) - 32) * 5 / 9
	fmt.Printf("Suhu dalam Celcius: %.1f\n", celcius)
}