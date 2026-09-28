package main
import "fmt"

func main(){
	var (
		phi float64 = 3.14
		r, luas float64
	)
	
	fmt.Print("Masukan jari-jari lingkaran: ")
	fmt.Scanln(&r)
	luas = phi * r * r
	fmt.Printf("Luas lingkaran dengan jari-jari %.0f adalah %.1f\n", r, luas)
}