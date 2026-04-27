package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bernaddwiki/koda-b7-go/internals/minitask1"
	"github.com/bernaddwiki/koda-b7-go/internals/minitask2"
	"github.com/bernaddwiki/koda-b7-go/internals/minitask3"
	"github.com/bernaddwiki/koda-b7-go/internals/minitask4"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printMenu()
		fmt.Print("Pilih fitur (0 untuk keluar): ")

		if !scanner.Scan() {
			fmt.Println("Input tidak tersedia.")
			return
		}

		choice, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			fmt.Println("Masukkan angka yang valid.")
			continue
		}

		switch choice {
		case 0:
			fmt.Println("Keluar. Terima kasih.")
			return
		case 1:
			handleCircleArea(scanner)
		case 2:
			handleCircleCircumference(scanner)
		case 3:
			handleCircleBoth(scanner)
		case 4:
			handleTriangle(scanner)
		case 5:
			handleSlice(scanner)
		case 6:
			minitask4.DefaultUser().Print()
		default:
			fmt.Println("Pilihan tidak tersedia, coba lagi.")
		}

		fmt.Println()
	}
}

func printMenu() {
	fmt.Println("====== MENU FITUR ======")
	fmt.Println("1. Hitung luas lingkaran")
	fmt.Println("2. Hitung keliling lingkaran")
	fmt.Println("3. Hitung luas dan keliling lingkaran")
	fmt.Println("4. Tampilkan segitiga siku")
	fmt.Println("5. Sisip angka ke dalam slice")
	fmt.Println("6. Tampilkan profil user")
	fmt.Println("0. Keluar")
	fmt.Println("========================")
}

func handleCircleArea(scanner *bufio.Scanner) {
	radius := readFloat(scanner, "Masukkan jari-jari lingkaran: ")
	fmt.Printf("Luas = %.2f\n", minitask1.HitungLuas(radius))
}

func handleCircleCircumference(scanner *bufio.Scanner) {
	radius := readFloat(scanner, "Masukkan jari-jari lingkaran: ")
	fmt.Printf("Keliling = %.2f\n", minitask1.HitungKeliling(radius))
}

func handleCircleBoth(scanner *bufio.Scanner) {
	radius := readFloat(scanner, "Masukkan jari-jari lingkaran: ")
	luas, keliling := minitask1.LuasKeliling(radius)
	fmt.Printf("Luas = %.2f\nKeliling = %.2f\n", luas, keliling)
}

func handleTriangle(scanner *bufio.Scanner) {
	n := readInt(scanner, "Masukkan tinggi segitiga siku: ")
	minitask2.SegitigaSiku(n)
}

func handleSlice(scanner *bufio.Scanner) {
	angka := readInt(scanner, "Masukkan angka yang akan disisipkan: ")
	minitask3.SisipAngkaSlice(angka)
}

func readFloat(scanner *bufio.Scanner, prompt string) float32 {
	for {
		fmt.Print(prompt)
		if !scanner.Scan() {
			fmt.Println("Input tidak tersedia.")
			return 0
		}
		value := strings.TrimSpace(scanner.Text())
		parsed, err := strconv.ParseFloat(value, 32)
		if err == nil {
			return float32(parsed)
		}
		fmt.Println("Masukkan angka desimal yang valid.")
	}
}

func readInt(scanner *bufio.Scanner, prompt string) int {
	for {
		fmt.Print(prompt)
		if !scanner.Scan() {
			fmt.Println("Input tidak tersedia.")
			return 0
		}
		value := strings.TrimSpace(scanner.Text())
		parsed, err := strconv.Atoi(value)
		if err == nil {
			return parsed
		}
		fmt.Println("Masukkan angka bulat yang valid.")
	}
}
