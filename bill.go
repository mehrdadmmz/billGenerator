package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// bill struct with all its attributes
type bill struct {
	billName   string
	billDate   string
	items      map[string]float64
	taxRate    float64
	tipValue   float64
	total      float64
	serverName string
}

// bill creator
func rawBill(n string, serverNames []string) bill {

	reader := bufio.NewReader(os.Stdin)

	currentDate := time.Now()
	formattedDate := currentDate.Format("2006-01-02")

	// def values
	tax := 0.03
	tip := 0.0

	serverPrompt := "What was the name of the server? "
	for x := 0; x < len(serverNames); x++ {
		serverPrompt += serverNames[x] + " "
	}

	name, err := getUserInput(serverPrompt, reader)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server name added...")

	b := bill{
		billName:   n,
		billDate:   formattedDate,
		items:      map[string]float64{},
		taxRate:    tax,
		tipValue:   tip,
		serverName: name,
	}

	return b
}

// option handler
func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	prompt := "Choose an option (a - add item, t - add tip, s - save your bill, q - quit): "

	option, _ := getUserInput(prompt, reader)
	lowerOption := strings.ToLower(option)

	switch lowerOption {
	case "a":
		itemName, _ := getUserInput("Item name: ", reader)
		price, _ := getUserInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number.")
			promptOptions(b)
		}

		b.addItem(itemName, p)

		fmt.Println("Item added - ", itemName, "$", price)
		promptOptions(b)
	case "t":
		tipPrompt := "How much tip would you like to add? $"
		t, _ := getUserInput(tipPrompt, reader)

		tip, err := strconv.ParseFloat(t, 64)

		if err != nil {
			fmt.Println("The tip amount must be a number.")
			promptOptions(b)
		}

		b.updateTip(tip)

		fmt.Println("Tip added - $", tip)
		promptOptions(b)

	case "s":
		b.save()
		fmt.Println("Bill saved!")

	case "q":
		break

	default:
		fmt.Println("Please enter a valid option (a - t - s)")
		promptOptions(b)
	}

}

// user input getter
func getUserInput(prompt string, r *bufio.Reader) (string, error) {

	fmt.Println(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

// item adder
func (b *bill) addItem(food string, price float64) {

	b.items[food] = price
}

// update tip
func (b *bill) updateTip(t float64) {

	b.tipValue = t
}

// printing the bill object
func (b *bill) format() string {
	//reader := bufio.NewReader(os.Stdin)
	restaurantName := "Perperook"
	total := b.totalPrice()

	fs := "<<<" + restaurantName + ">>>\n"

	// date
	fs += fmt.Sprintf("%-25v ...%v \n", "date: ", b.billDate)

	// server name
	fs += fmt.Sprintf("%-25v ...%v \n", "server name: ", b.serverName)

	// list items
	for key, value := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", key+":", value)
		total += value
	}

	// tip
	fs += fmt.Sprintf("%-25v ...%0.2f \n", "tip: ", b.tipValue)

	// tax
	fs += fmt.Sprintf("%-25v ...%0.2f \n", "tax rate: ", b.taxRate)

	// total
	fs += fmt.Sprintf("%-25v ...%0.2f", "total: ", total)

	return fs

}

// total counter
func (b *bill) totalPrice() float64 {

	t := 0.0
	for key, _ := range b.items {
		t += b.items[key]
	}

	// add tax
	t *= (1 * b.taxRate)

	// add tip
	t += b.tipValue

	return t
}

// generating and saving a bill
func (b *bill) save() {

	data := []byte(b.format())
	err := os.WriteFile("generated_bills/"+b.billName+".txt", data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("New bill named %v is saved!", b.billName)
}
