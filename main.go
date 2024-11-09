/*
A typical restaurant bill, also known as a check or invoice, includes several key components. When designing a program
to generate a restaurant bill, you should consider the following elements:

1. **Header Information**:
  - Restaurant Name and Logo (if applicable)
  - Address and Contact Information of the Restaurant
  - Date and Time of the Bill

2. **Table and Server Details** (optional, but often included):
  - Table Number or Name
  - Server or Waiter's Name

3. **Order Details**:
  - List of Ordered Items: Each item should include a description (e.g., "Grilled Chicken Sandwich", "Iced Tea").
  - Quantity of Each Item: The number of each item ordered.
  - Price Per Item: The cost of each individual item.
  - Total Item Cost: Quantity multiplied by the price per item.

4. **Subtotals and Taxes**:
  - Subtotal: The total cost of all items before taxes and additional fees.
  - Sales Tax: The tax amount applied based on the subtotal.
  - Any Other Taxes or Fees: Depending on the location, there might be additional taxes or fees.

5. **Discounts or Offers** (if applicable):
  - Description of the Discount or Offer
  - Amount Deducted from the Subtotal

6. **Total Amount**:
  - The final amount due after adding taxes and subtracting any discounts.

7. **Payment Method** (optional, but useful for records):
  - Indicate whether the bill was paid by cash, credit card, or another method.

8. **Footer Information**:
  - Thank You Note or Restaurant's Custom Message
  - Instructions for Feedback or Surveys (if applicable)

9. **Legal or Regulatory Information** (if required):
  - Tax Identification Number (TIN) or similar identifiers
  - Any legal disclaimers or regulatory information required in your region.

When programming this in Go, you'll likely structure these as fields in a struct or class,
and your program will fill these fields based on user input or database queries. Remember to
format the final output in a clear and readable manner, whether it's intended for print or digital display.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	serversList := []string{
		"Jason",
		"Brian",
		"Sophie",
		"Lily",
	}

	//tableNumbers := []int8{1, 2, 3, 4, 5, 6, 7, 8}

	reader := bufio.NewReader(os.Stdin)
	billName, _ := getUserInput("Create a new bill name - ", reader)

	myBill := rawBill(billName, serversList)

	promptOptions(myBill)

	sep := strings.Repeat("-", 30)

	fmt.Print(sep)
	myBill.format()
	fmt.Print(sep)

}
