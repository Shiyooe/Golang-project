package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())


	var secretNumber int = rand.Intn(100) + 1 
	var guess int                          
	var attempts int = 0                 
	var maxAttempts int = 7              
	var gameOver bool = false              

	fmt.Println("🎮 Welcome to Guess The Number!")
	fmt.Println("Aku sudah memilih angka 1 - 100")
	fmt.Println("Kamu punya", maxAttempts, "kesempatan")
	fmt.Println("--------------------------------")

	// ===== GAME LOOP =====
	for !gameOver {
		fmt.Print("Masukkan tebakanmu: ")
		fmt.Scan(&guess)

		attempts++

		if guess == secretNumber {
			fmt.Println("🎉 BENAR!")
			fmt.Println("Kamu menebak dalam", attempts, "percobaan")
			gameOver = true

		} else if guess < secretNumber {
			fmt.Println("📉 Terlalu kecil")

		} else {
			fmt.Println("📈 Terlalu besar")
		}

		
		if attempts >= maxAttempts && !gameOver {
			fmt.Println("💀 Kesempatan habis!")
			fmt.Println("Angka rahasianya adalah:", secretNumber)
			gameOver = true
		}

		fmt.Println("--------------------------------")
	}

	fmt.Println("Game selesai. Thanks sudah main ❤️")
}
