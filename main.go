package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Menu struct {
	Name  string
	Price float64
	Hours int
}

type Order struct {
	Name     string
	Qty      int
	Subtotal float64
	Total    float64
}

func main() {

	menu := map[string]Menu{
		"historical": {
			Name:  "Historical",
			Price: 300000,
		},
		"cultural": {
			Name:  "Cultural",
			Price: 250000,
		},
		"adventure": {
			Name:  "Adventure",
			Price: 350000,
		},
		"nature": {
			Name:  "Nature",
			Price: 280000,
		},
		"food tour": {
			Name:  "Food Tour",
			Price: 320000,
		},
	}

	reader := bufio.NewReader(os.Stdin)

	var orders []Order
	var total float64

	for {
		fmt.Println("\nWelcome to the Tour-Guide Booking System!")
		fmt.Println("Here are our services:")

		// Menampilkan menu
		for _, key := range []string{
			"historical",
			"cultural",
			"adventure",
			"nature",
			"food tour",
		} {
			item := menu[key]
			fmt.Printf("%-12s : Rp%.0f/hour\n", item.Name, item.Price)
		}

		// Pertanyaan 1
		fmt.Print("\nWhat type of tour guide would you like to book? ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		// Mengecek apakah service tersedia
		item, exists := menu[input]

		if !exists {
			fmt.Println("Service not available. Please select a valid tour guide type from the menu.")
			continue
		}

		// Pertanyaan 2
		fmt.Print("How many hours? ")

		qtyInput, _ := reader.ReadString('\n')
		qtyInput = strings.TrimSpace(qtyInput)

		qty, err := strconv.Atoi(qtyInput)

		if err != nil || qty <= 0 {
			fmt.Println("Invalid Number of Hours. Please enter a positive number of hours in numerical format.")
			continue
		}

		if qty > 24 {
			fmt.Println("Number of hours too large. Please enter a reasonable number of hours (maximum 24).")
			continue
		}

		// Menyimpan jumlah jam ke Menu
		item.Hours = qty

		// Menghitung subtotal
		itemSubtotal := item.Price * float64(item.Hours)

		// Menambahkan subtotal ke total
		total += itemSubtotal

		// Menyimpan order
		orders = append(orders, Order{
			Name:     item.Name,
			Qty:      item.Hours,
			Subtotal: itemSubtotal,
			Total:    total,
		})

		// Pertanyaan 3
		fmt.Print("Would you like to book another service? (yes/no) ")

		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(strings.ToLower(answer))

		if answer == "no" {
			break
		}

		if answer != "yes" {
			fmt.Println("Please enter yes or no.")
			continue
		}
	}

	// Menampilkan rincian pesanan
	fmt.Println("\n===== Order Summary =====")

	for _, order := range orders {
		price := menu[strings.ToLower(order.Name)].Price
		fmt.Printf(
			"%s: %d hour(s) x Rp%.0f/hour = Rp%.0f\n",
			order.Name,
			order.Qty,
			price,
			order.Subtotal,
		)
	}

	// Menampilkan total keseluruhan
	fmt.Printf("\nYour final total is Rp %.0f\n", total)
}
