package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Item represents a shopping item
type Item struct {
	Name     string
	Price    float64
	Quantity int
}

// Total calculates total price for the item
func (i Item) Total() float64 {
	return i.Price * float64(i.Quantity)
}

// ShoppingCart manages shopping items
type ShoppingCart struct {
	Items []Item
}

// AddItem adds an item to cart
func (sc *ShoppingCart) AddItem(item Item) {
	sc.Items = append(sc.Items, item)
	fmt.Printf("✓ %s ditambahkan ke keranjang\n", item.Name)
}

// GetTotal calculates total cart value
func (sc *ShoppingCart) GetTotal() float64 {
	total := 0.0
	for _, item := range sc.Items {
		total += item.Total()
	}
	return total
}

// ApplyDiscount applies percentage discount
func (sc *ShoppingCart) ApplyDiscount(percent float64) float64 {
	total := sc.GetTotal()
	discount := total * (percent / 100)
	return total - discount
}

// DisplayCart shows all items in cart
func (sc *ShoppingCart) DisplayCart() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("                    KERANJANG BELANJA")
	fmt.Println(strings.Repeat("=", 60))
	
	if len(sc.Items) == 0 {
		fmt.Println("Keranjang masih kosong")
		return
	}

	fmt.Printf("%-20s %-10s %-10s %-15s\n", "Nama Item", "Harga", "Jumlah", "Subtotal")
	fmt.Println(strings.Repeat("-", 60))
	
	for _, item := range sc.Items {
		fmt.Printf("%-20s $%-10.0f %-10d $%-15.0f\n", 
			item.Name, item.Price, item.Quantity, item.Total())
	}
	
	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("%-41s $%-15.0f\n", "TOTAL:", sc.GetTotal())
	fmt.Println(strings.Repeat("=", 60))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cart := ShoppingCart{}

	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║     APLIKASI SHOPPING CALCULATOR       ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Println()

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah item")
		fmt.Println("2. Lihat keranjang")
		fmt.Println("3. Hitung total")
		fmt.Println("4. Terapkan diskon")
		fmt.Println("5. Keluar")
		fmt.Print("\nPilih menu (1-5): ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			// Input nama item
			fmt.Print("Nama item: ")
			scanner.Scan()
			name := scanner.Text()

			// Input harga
			fmt.Print("Harga satuan ($): ")
			scanner.Scan()
			priceStr := scanner.Text()
			price, err := strconv.ParseFloat(priceStr, 64)
			if err != nil {
				fmt.Println("❌ Harga tidak valid!")
				continue
			}

			// Input quantity
			fmt.Print("Jumlah: ")
			scanner.Scan()
			qtyStr := scanner.Text()
			quantity, err := strconv.Atoi(qtyStr)
			if err != nil {
				fmt.Println("❌ Jumlah tidak valid!")
				continue
			}

			item := Item{
				Name:     name,
				Price:    price,
				Quantity: quantity,
			}
			cart.AddItem(item)

		case "2":
			cart.DisplayCart()

		case "3":
			total := cart.GetTotal()
			fmt.Printf("\n💰 Total Belanja: $%.0f\n", total)

		case "4":
			fmt.Print("Masukkan persentase diskon (%): ")
			scanner.Scan()
			discountStr := scanner.Text()
			discount, err := strconv.ParseFloat(discountStr, 64)
			if err != nil {
				fmt.Println("❌ Persentase tidak valid!")
				continue
			}

			originalTotal := cart.GetTotal()
			finalTotal := cart.ApplyDiscount(discount)
			savings := originalTotal - finalTotal

			fmt.Printf("\n💳 Total sebelum diskon: $%.0f\n", originalTotal)
			fmt.Printf("🎉 Diskon %.0f%%: -$%.0f\n", discount, savings)
			fmt.Printf("✨ Total setelah diskon: $%.0f\n", finalTotal)

		case "5":
			fmt.Println("\n👋 Terima kasih telah berbelanja!")
			return

		default:
			fmt.Println("❌ Pilihan tidak valid!")
		}
	}
}