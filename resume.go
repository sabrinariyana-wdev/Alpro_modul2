package main
import "fmt"

func main() {
	var (
	nama, nim, kelas string
	)

	fmt.Print("Masukan nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukan NIM: ")
	fmt.Scanln(&nim)
	fmt.Print("Masukan kelas: ")
	fmt.Scanln(&kelas)

	fmt.Println("Perkenalan nama saya " + nama + " dengan NIM " + nim + " Saya dari kelas " + kelas)
}